package usecases

import (
	"context"
	"fmt"

	"websmee/buyspot/internal/domain"
)

type OrderReader struct {
	orderRepository OrderRepository
}

func NewOrderReader(orderRepository OrderRepository) *OrderReader {
	return &OrderReader{orderRepository}
}

func (r *OrderReader) GetUserOrders(ctx context.Context) ([]domain.Order, error) {
	userID := domain.GetCtxUserID(ctx)
	if userID == "" {
		return nil, domain.ErrUnauthorized
	}

	orders, err := r.orderRepository.GetUserActiveOrders(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not get orders for user ID = '%s', err: %w", userID, err)
	}

	return orders, nil
}
