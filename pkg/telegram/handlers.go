package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
	replyStartTemplate = "Ссылка:\n%s"
) 

func (b *Bot) handleCommand(message *tgbotapi.Messages) {
	
	switch message.Command() {
		case commandStart:
			return b.handleStartCommand(message)
		default:
			return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Messages) {
	
	log.Printf("[%s] %s", message.Message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Message.Chat.ID, message.Text)
	b.bot.Send(msg)
		
}

func (b *Bot)  handleStartCommand(message *tgbotapi.Message) error {
	authLink, err  := b.generateAuthorizationLink(message.Chat.ID)
	if err = nil {
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(replyStartTemplate, authLink))
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Нельзя так просто взять и…")
	_, err := b.bot.Send(msg)
	return err
}
