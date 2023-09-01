package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"github.com/zhashkevych/telegram-pocket-bot/pkg/config"
	"github.com/zhashkevych/telegram-pocket-bot/pkg/storage"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	pocket.Client *pocket.Client
	tokenRepository repository.TokenRepository
	redirectURL string
}

func NewBot(bot *tgbotapi.BotAPI, pocket.Client *pocket.Client, tr repository.TokenRepository, redirectURL string) *Bot {
	return &Bot{bot: bot, pocket.Client: pocket.Client, redirectURL: redirectURL, tokenRepository: tr}
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