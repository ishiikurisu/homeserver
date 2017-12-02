package bot

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "strconv"
    "fmt"
    "github.com/ishiikurisu/house"
    "os/exec"
    "strings"
)

type Bot struct {
    Token string
    Allowed []int64

}

// Creates an empty bot.
func Empty() *Bot {
    bot := Bot {
        Token: "",
        Allowed: make([]int64, 0),
    }
    return &bot
}

// Creates a new bot. It requires the bot's Telegram token and the path to the
// allowed ids list. If the path is empty, it will try to load from
// `./data/homeserver/allowed.yml`.
// IDEA: Maybe it should receive the allowed ids instead of a filepath.
func New(token, allowedPath string) (*Bot, error) {
    bot := Empty()
    allowed := make([]int64, 0)
    path := allowedPath

    if len(path) == 0 {
        path = "./data/homeserver/allowed.yml"
    }
    allowed, oops := LoadAllowed(path)

    bot.Token = token
    bot.Allowed = allowed
    return bot, oops
}

// This file must be a YAML file with a `allowed` list where each item must be
// an allowed id to use the bot.
func LoadAllowed(where string) ([]int64, error) {
    outlet := make([]int64, 0)
    raw, oops := ioutil.ReadFile(where)
    if oops != nil {
        return outlet, oops
    }

    var f interface{}
    oops = yaml.Unmarshal(raw, &f)
    if oops != nil {
        return outlet, oops
    }

    everything := f.(map[interface{}]interface{})
    midlet := everything["allowed"].([]interface{})
    for _, it := range midlet {
        v, oops := strconv.ParseInt(fmt.Sprintf("%v", it), 10, 64)
        if oops != nil {
            return make([]int64, 0), oops
        }
        outlet = append(outlet, v)
    }

    return outlet, nil
}

// Generates the answer for the user based on its id. Should send back the IP
// address of the current machine.
func (bot *Bot) Answer(targetId int64) string {
    allowed := false
    outlet := "What are you doing here?"

    for _, id := range bot.Allowed {
        if id == targetId {
            allowed = true
        }
    }

    if allowed {
        ip := DiscoverIpAddress()
        outlet = "I am running on Windows! $$ No IP for you unless you pay. $$"
        if len(ip) > 0 {
            outlet = fmt.Sprintf("The IP address is %s\nHave fun!\n", ip)
        }
    }

    return outlet
}

// Discovers the IP address of the current machine.
func DiscoverIpAddress() string {
  ip := ""

  if house.GetOS() == "win32" {
    return ip
  }

  cmd := exec.Command("ip", "addr", "show")
  rawOutput, oops := cmd.Output()
  if oops == nil {
    output := string(rawOutput)
    lines := strings.Split(output, "\n")
    for _, line := range lines {
      fields := strings.Split(strings.Trim(line, " "), " ")
      if (fields[0] == "inet") && (strings.Contains(fields[1], "192")) {
        ip = strings.Split(fields[1], "/")[0]
      }
    }
  }


  return ip
}

// Gets the token of the Telegram bot.
func GetToken() (string, error) {
    outlet := ""
    raw, oops := ioutil.ReadFile("./data/homeserver/token.yml")
    if oops != nil {
        return outlet, oops
    }

    var f interface{}
    oops = yaml.Unmarshal(raw, &f)
    if oops != nil {
        return outlet, oops
    }

    everything := f.(map[interface{}]interface{})
    midlet := everything["token"].(interface{})
    outlet = fmt.Sprintf("%v", midlet)

    return outlet, nil
}
