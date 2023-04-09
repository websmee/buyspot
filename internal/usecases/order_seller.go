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
	userID := domain.GetCtxUserID(ctx)
	if userID == "" {
		return nil, domain.ErrUnauthorized
	}

	balance, err := s.balanceService.GetUserActiveBalance(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s', err: %w", userID, err)
	}

	order, err := s.orderRepository.GetUserOrderByID(ctx, userID, orderID)
	if err != nil {
		return nil, fmt.Errorf("could get order by ID = '%s' for user ID = '%s', err: %w", orderID, userID, err)
	}

	sellAmount, err := s.converterService.Convert(ctx, userID, order.ToAmount, order.ToSymbol, balance.Symbol)
	if err != nil {
		return nil, fmt.Errorf(
			"could not convert %s to %s for user ID = '%s', err: %w",
			order.ToSymbol,
			balance.Symbol,
			userID,
			err,
		)
	}

	order.CloseAmount = sellAmount
	order.CloseSymbol = balance.Symbol
	order.Updated = time.Now()
	order.Status = domain.OrderStatusClosed
	if err := s.orderRepository.SaveOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("could not save order after conversion, err: %w", err)
	}

	balance, err = s.balanceService.GetUserActiveBalance(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s' after conversion, err: %w", userID, err)
	}

	return balance, nil
}
