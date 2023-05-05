package simplepush

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const apiURL = "https://api.simplepush.io/send"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SendNotification(key, title, message, event string) error {
	data := url.Values{}
	data.Set("key", key)
	data.Set("title", title)
	data.Set("msg", message)
	data.Set("event", event)

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
