package bot

import (
    "testing"
)

func TestCanCreateAnEmptyBot(t *testing.T) {
    bot := Empty()
    if len(bot.Api) > 0 {
        t.Error("There is an API here? Why?")
        return
    }
}
