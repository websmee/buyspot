package example

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"websmee/buyspot/internal/domain"
)

type OrderRepository struct {
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) GetUserOrders(_ context.Context, _ string) ([]domain.Order, error) {
	return []domain.Order{
		{
			ID:          primitive.NewObjectID(),
			FromAmount:  123.45,
			FromTicker:  "USDT",
			ToAmount:    0.0043,
			ToTicker:    "BTC",
			ToAssetName: "Bitcoin",
			TakeProfit:  3,
			StopLoss:    -1,
			Created:     time.Now().Add(-1 * time.Hour),
		},
		{
			ID:          primitive.NewObjectID(),
			FromAmount:  111.11,
			FromTicker:  "USDT",
			ToAmount:    0.061,
			ToTicker:    "ETH",
			ToAssetName: "Ethereum",
			TakeProfit:  4,
			StopLoss:    -2,
			Created:     time.Now().Add(-24 * time.Hour),
		},
	}, nil
}

func (r *OrderRepository) GetActiveOrdersCountByTicker(ctx context.Context, ticker string) (int, error) {
	return 3, nil
}

func (r *OrderRepository) SaveOrder(ctx context.Context, order *domain.Order) error {
	return nil
}
