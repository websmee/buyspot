package simplepush

import (
	"context"
	"strings"

	"websmee/buyspot/internal/domain"
)

type NewSpotsNotifier struct {
	client *Client
}

func NewNewSpotsNotifier(client *Client) *NewSpotsNotifier {
	return &NewSpotsNotifier{client}
}

func (n *NewSpotsNotifier) Notify(_ context.Context, user *domain.User, spots []domain.Spot) error {
	if user.NotificationsKey == "" {
		return nil
	}

	if len(spots) == 0 {
		return nil
	}

	symbols := make([]string, 0, len(spots))
	for _, spot := range spots {
		symbols = append(symbols, spot.Asset.Symbol)
	}

	return n.client.SendNotification(
		user.NotificationsKey,
		"NEW SPOTS",
		strings.Join(symbols, ", "),
		"new_spots",
	)
}
