package telegram

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type NotificationsSubscriber struct {
	telegramBotAPIToken string
}

func NewNotificationsSubscriber(telegramBotAPIToken string) *NotificationsSubscriber {
	return &NotificationsSubscriber{telegramBotAPIToken: telegramBotAPIToken}
}

func (b *NotificationsSubscriber) Run(
	updateKey func(username, key string) error,
	handleErr func(err error),
) error {
	botAPI, err := tgbotapi.NewBotAPI(b.telegramBotAPIToken)
	if err != nil {
		return fmt.Errorf("could not create telegram botAPI: %w", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := botAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.Text == "/start" {
			if update.Message.From != nil {
				text := "success"
				if err := updateKey(update.Message.From.UserName, strconv.Itoa(int(update.Message.Chat.ID))); err != nil {
					text = "error"
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				msg.ReplyToMessageID = update.Message.MessageID

				if _, err := botAPI.Send(msg); err != nil {
					handleErr(err)
				}
			}
		}
	}

	return nil
}
