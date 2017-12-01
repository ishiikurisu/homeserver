package bot

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "strconv"
    "fmt"
)

type Bot struct {
    Token string
    Allowed []int

}

// Creates an empty bot.
func Empty() *Bot {
    bot := Bot {
        Token: "",
        Allowed: make([]int, 0),
    }
    return &bot
}

// Creates a new bot. It requires the bot's Telegram token and the path to the
// allowed ids list. If the path is empty, it will try to load from
// `./data/homeserver/allowed.yml`.
func New(token, allowedPath string) (*Bot, error) {
    bot := Empty()
    allowed := make([]int, 0)
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
func LoadAllowed(where string) ([]int, error) {
    outlet := make([]int, 0)
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
        v, oops := strconv.Atoi(fmt.Sprintf("%v", it))
        if oops != nil {
            return make([]int, 0), oops
        }
        outlet = append(outlet, v)
    }

    return outlet, nil
}
