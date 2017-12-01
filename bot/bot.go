package bot

type Bot struct {
    Api string

}

func Empty() *Bot {
    bot := Bot {
        Api: "",
    }
    return &bot
}
