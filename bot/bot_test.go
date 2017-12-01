package bot

import (
    "testing"
    "fmt"
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
    } else if oops != nil {
        t.Error(fmt.Sprintf("Some other error happenned while trying to load the allowed ids file: %s\n", oops))
    }
    // IDEA Check for corruption in data
}
