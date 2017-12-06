package bot

import (
    "testing"
    "fmt"
    "github.com/ishiikurisu/house"
)

func TestCanCreateAnEmptyBot(t *testing.T) {
    bot := Empty()
    if len(bot.Token) > 0 || len(bot.Allowed) > 0 {
        t.Error("There is an API here? Why?")
        return
    }
}

func TestCanCreateANewBot(t *testing.T) {
    bot, oops := New("0", "")
    if oops != nil {
        t.Error(fmt.Sprintf("%s\n", oops))
        return
    }
    fmt.Printf("%s\n", bot.Token)
}

func TestLoadingConfigurationFile(t *testing.T) {
    allowed, oops := LoadAllowed("./data/homeserver/allowed.yml")
    if len(allowed) != 1 {
        t.Error("Couldn't load the allowed ids' file")
        return
    } else if oops != nil {
        t.Error(fmt.Sprintf("Some other error happenned while trying to load the allowed ids file: %s\n", oops))
        return
    }

    if allowed[0] != 190141641 {
        t.Error("Data corruption!")
        return
    }
}

func TestBotOnlyAnswersIfIdIsInAllowedList(t *testing.T) {
    bot, oops := New("0", "")
    if oops != nil {
        t.Error("Shouldn't procede!")
        return
    }

    answer := bot.Answer(0)
    if answer != "What are you doing here?" {
        t.Error("Why is it allowing answers to unknown folk?")
    }
    answer = bot.Answer(190141641)
    if answer == "What are you doing here?" {
        t.Error("Why isn't it treating the costumer correctly?")
    }
}

func TestDiscoverIpAddress(t *testing.T) {
  if house.GetOS() == "win32" {
    return
  }

  ip := DiscoverIpAddress()
  if ip != "192.168.0.111" {
    t.Error(fmt.Sprintf("Incorrect IP address detected: %s\n", ip))
  }
}
