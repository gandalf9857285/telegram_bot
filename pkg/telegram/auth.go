package telegram

import (
	"context"
	"fmt"
)

func (b *Bot) generateAuthorizationLink(chatID int64) (string, error) {
	redirectURL := b.generateRedirectURL(chatID)
	requestToken, err := b.client.GetRequestToken(context.Background(), b.redirectURL)
	if err != nil {
		return "", err
	}

	return b.client.GetAuthorizationURL(requestToken, redirectURL)
}

func (b *Bot) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", b.redirectURL, chatID)
}
