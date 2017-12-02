package main

import (
  brain "github.com/ishiikurisu/homeserver/bot"
  telegram "gopkg.in/telegram-bot-api.v4"
  "github.com/ishiikurisu/homeserver/server"
  "fmt"
)

func main() {
  go runBot()
  go runSite()
  for true {

  }
}

func runBot() {
  token, oops := brain.GetToken()
  if oops != nil {
    return
  }
  b := brain.New(token, "")
  bot, oops := tgbotapi.NewBotAPI(b.Token)
	if oops != nil {
		fmt.Printf("%s\n", oops)
    return
	}

	u := telegram.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
    id := update.Message.Chat.ID
		msg := telegram.NewMessage(id, b.Answer(id))
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}

func runSite() {
  s := server.NewServer()
  s.Run()
}
