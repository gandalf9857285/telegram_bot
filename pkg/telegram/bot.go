package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	pocket.Client *pocket.Client
	redirectURL string
}

func NewBot(bot *tgbotapi.BotAPI, pocket.Client *pocket.Client, redirectURL string) *Bot {
	return &Bot{bot: bot, pocket.Client: pocket.Client, redirectURL: redirectURL}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	
	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(update)
	return nil
}

func (b *Bot) handleUpdates(update tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil { // If we got a message
			continue
		}
		update.Message.InCommand() {
			b.handleCommand(update.Message)
			continue
		}
		b.handleMessage(update.Message)
	}
}


func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return bot.GetUpdatesChan(u)
	
}