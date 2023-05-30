package telegram

import (
	"context"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"websmee/buyspot/internal/domain"
)

type Notifier struct {
	telegramBotAPIToken string
}

func NewNotifier(telegramBotAPIToken string) *Notifier {
	return &Notifier{telegramBotAPIToken}
}

func (c *Notifier) NotifyAll(ctx context.Context, title, message string) error {
	return c.notifyChat(ctx, -1001963992438, title, message) // move chatID to config
}

func (c *Notifier) NotifyUser(ctx context.Context, user *domain.User, title, message string) error {
	if user.TelegramUsername == "" || user.NotificationsKey == "" {
		return nil
	}

	chatID, err := strconv.Atoi(user.NotificationsKey)
	if err != nil {
		return fmt.Errorf("ivalid notifications key %s, err: %w", user.NotificationsKey, err)
	}

	return c.notifyChat(ctx, int64(chatID), title, message)
}

func (c *Notifier) notifyChat(_ context.Context, chatID int64, title, message string) error {
	bot, err := tgbotapi.NewBotAPI(c.telegramBotAPIToken)
	if err != nil {
		return fmt.Errorf("could not create telegram botAPI: %w", err)
	}

	if _, err := bot.Send(tgbotapi.NewMessage(
		chatID,
		fmt.Sprintf("%s: %s", title, message),
	)); err != nil {
		return fmt.Errorf("could not send message to telegram: %w", err)
	}

	return nil
}
