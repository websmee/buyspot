package usecases

import (
	"context"
	"fmt"
	"time"

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

	balance, err := s.balanceService.GetUserActiveBalance(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s', err: %w", user.ID, err)
	}

	order, err := s.orderRepository.GetUserOrderByID(ctx, user.ID, orderID)
	if err != nil {
		return nil, fmt.Errorf("could get order by ID = '%s' for user ID = '%s', err: %w", orderID, user.ID, err)
	}

	sellAmount, err := s.converterService.Convert(ctx, user, order.ToAmount, order.ToTicker, balance.Ticker)
	if err != nil {
		return nil, fmt.Errorf(
			"could not convert %s to %s for user ID = '%s', err: %w",
			order.ToTicker,
			balance.Ticker,
			user.ID,
			err,
		)
	}

	order.CloseAmount = sellAmount
	order.CloseTicker = balance.Ticker
	order.Updated = time.Now()
	order.Status = domain.OrderStatusClosed
	if err := s.orderRepository.SaveOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("could not save order after conversion, err: %w", err)
	}

	balance, err = s.balanceService.GetUserActiveBalance(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s' after conversion, err: %w", user.ID, err)
	}

	return balance, nil
}
