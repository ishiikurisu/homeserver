package main

import (
  brain "github.com/ishiikurisu/homeserver/bot"
  telegram "github.com/go-telegram-bot-api/telegram-bot-api"
  "github.com/ishiikurisu/homeserver/server"
  "fmt"
)

func main() {
  go runBot()
  runSite()
}

func runBot() {
    token, oops := brain.GetToken()
    if oops != nil {
        fmt.Printf("%s\n", oops)
        return
    }
    b, oops := brain.New(token, "")
    if oops != nil {
        fmt.Printf("%s\n", oops)
        return
    }
    bot, oops := telegram.NewBotAPI(b.Token)
    if oops != nil {
        fmt.Printf("%s\n", oops)
        return
    }

    u := telegram.NewUpdate(0)
    u.Timeout = 60
    updates, oops := bot.GetUpdatesChan(u)
    for update := range updates {
        if update.Message == nil {
            continue
        }
        id := update.Message.Chat.ID
        fmt.Printf("%d tried to talk!\n", id)
        msg := telegram.NewMessage(id, b.Answer(id))
        msg.ReplyToMessageID = update.Message.MessageID
        bot.Send(msg)
    }
}

func runSite() {
  s := server.NewServer("8000")
  s.Run()
}
