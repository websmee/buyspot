package usecases

import (
	"context"
	"fmt"

	"websmee/buyspot/internal/domain"
)

type OrderSeller struct {
	orderRepository  OrderRepository
	converterService ConverterService
	balanceService   BalanceService
}

func NewOrderSeller(
	orderRepository OrderRepository,
	converterService ConverterService,
	balanceService BalanceService,
) *OrderSeller {
	return &OrderSeller{
		orderRepository,
		converterService,
		balanceService,
	}
}

func (s *OrderSeller) SellOrder(ctx context.Context, orderID string) (*domain.Balance, error) {
	user := domain.GetCtxUser(ctx)
	if user == nil {
		return nil, domain.ErrUnauthorized
	}

	balance, err := s.balanceService.GetUserBalance(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s', err: %w", user.ID, err)
	}

	order, err := s.orderRepository.GetUserOrderByID(ctx, user.ID, orderID)
	if err != nil {
		return nil, fmt.Errorf("could get order by ID = '%s' for user ID = '%s', err: %w", orderID, user.ID, err)
	}

	if _, err := s.converterService.Convert(ctx, order.ToAmount, order.ToTicker, balance.Ticker); err != nil {
		return nil, fmt.Errorf(
			"could not convert %s to %s for user ID = '%s', err: %w",
			order.ToTicker,
			balance.Ticker,
			user.ID,
			err,
		)
	}

	order.Status = domain.OrderStatusCompleted
	if err := s.orderRepository.SaveOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("could not save order after conversion, err: %w", err)
	}

	balance, err = s.balanceService.GetUserBalance(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s' after conversion, err: %w", user.ID, err)
	}

	return balance, nil
}
