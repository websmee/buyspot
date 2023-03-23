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
	user := domain.GetCtxUser(ctx)
	if user == nil {
		return nil, domain.ErrUnauthorized
	}

	orders, err := r.orderRepository.GetUserOrders(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("could not get orders for user ID = '%s', err: %w", user.ID, err)
	}

	return orders, nil
}
