package main

import (
	"fmt"
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(token: "6586940301:AAGiotAJpM1BAEUJv4BmofG3kiH9i8nmMnE")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	pocket.Client, err := pocket.NewClient("")
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot, pocket.Client, "http://localhost"
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}