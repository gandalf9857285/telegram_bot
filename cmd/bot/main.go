package main

import (
	"log"

	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"github.com/zhashkevych/telegram-bot-youtube/pkg/telegram"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6586940301:AAGiotAJpM1BAEUJv4BmofG3kiH9i8nmMnE")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("")
	if err != nil {
		log.Fatal(err)
	}

	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, "http://localhost")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
