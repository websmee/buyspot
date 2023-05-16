package usecases

import (
	"context"
	"fmt"
	"time"

	"websmee/buyspot/internal/domain"
)

type OrderSeller struct {
	userRepository     UserRepository
	orderRepository    OrderRepository
	tradingService     TradingService
	demoTradingService TradingService
	balanceService     BalanceService
}

func NewOrderSeller(
	userRepository UserRepository,
	orderRepository OrderRepository,
	tradingService TradingService,
	demoTradingService TradingService,
	balanceService BalanceService,
) *OrderSeller {
	return &OrderSeller{
		userRepository,
		orderRepository,
		tradingService,
		demoTradingService,
		balanceService,
	}
}

func (s *OrderSeller) SellOrder(ctx context.Context, orderID string) (*domain.Balance, error) {
	userID := domain.GetCtxUserID(ctx)
	if userID == "" {
		return nil, domain.ErrUnauthorized
	}

	user, err := s.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not get user by ID = '%s', err: %w", userID, err)
	}

	balance, err := s.balanceService.GetUserActiveBalance(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s', err: %w", userID, err)
	}

	order, err := s.orderRepository.GetUserOrderByID(ctx, userID, orderID)
	if err != nil {
		return nil, fmt.Errorf("could get order by ID = '%s' for user ID = '%s', err: %w", orderID, userID, err)
	}

	var sellAmount float64
	if user.IsDemo {
		sellAmount, err = s.demoTradingService.Sell(ctx, user, order.ToSymbol, order.ToAmount, balance.Symbol)
		if err != nil {
			return nil, fmt.Errorf(
				"could not sell %s for %s as demo user with ID = '%s', err: %w",
				order.ToSymbol,
				balance.Symbol,
				userID,
				err,
			)
		}
	} else {
		sellAmount, err = s.tradingService.Sell(ctx, user, order.ToSymbol, order.ToAmount, balance.Symbol)
		if err != nil {
			return nil, fmt.Errorf(
				"could not sell %s for %s as user with ID = '%s', err: %w",
				order.ToSymbol,
				balance.Symbol,
				userID,
				err,
			)
		}
	}

	order.CloseAmount = sellAmount
	order.CloseSymbol = balance.Symbol
	order.Updated = time.Now()
	order.Status = domain.OrderStatusClosed
	if err := s.orderRepository.SaveOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("could not save order after conversion, err: %w", err)
	}

	balance, err = s.balanceService.GetUserActiveBalance(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s' after conversion, err: %w", userID, err)
	}

	return balance, nil
}
