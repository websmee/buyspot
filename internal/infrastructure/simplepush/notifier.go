package simplepush

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"websmee/buyspot/internal/domain"
)

const apiURL = "https://api.simplepush.io/send"

type Notifier struct {
}

func NewNotifier() *Notifier {
	return &Notifier{}
}

func (c *Notifier) Notify(_ context.Context, user *domain.User, title, message string) error {
	data := url.Values{}
	data.Set("key", user.NotificationsKey)
	data.Set("title", title)
	data.Set("msg", message)
	data.Set("event", title)

	u, _ := url.ParseRequestURI(apiURL)
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))

	res, err := client.Do(r)
	if err != nil {
		return fmt.Errorf("could not send notification, err: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("could not send notification, status code: %d", res.StatusCode)
	}

	return nil
}
