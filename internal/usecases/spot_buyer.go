package usecases

import (
	"context"
	"fmt"
	"time"

	"websmee/buyspot/internal/domain"
)

type SpotBuyer struct {
	orderRepository  OrderRepository
	converterService ConverterService
	balanceService   BalanceService
	assetRepository  AssetRepository
}

func NewSpotBuyer(
	orderRepository OrderRepository,
	converterService ConverterService,
	balanceService BalanceService,
	assetRepository AssetRepository,
) *SpotBuyer {
	return &SpotBuyer{
		orderRepository,
		converterService,
		balanceService,
		assetRepository,
	}
}

func (b *SpotBuyer) BuySpot(ctx context.Context, amount float64, symbol string, takeProfit, stopLoss float64) (*domain.Balance, error) {
	userID := domain.GetCtxUserID(ctx)
	if userID == "" {
		return nil, domain.ErrUnauthorized
	}

	balance, err := b.balanceService.GetUserActiveBalance(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s', err: %w", userID, err)
	}

	asset, err := b.assetRepository.GetAssetBySymbol(ctx, symbol)
	if err != nil {
		return nil, fmt.Errorf("could not available assets, err: %w", err)
	}
	if asset == nil {
		return nil, fmt.Errorf("could not find asset by symbol %s", symbol)
	}

	order := &domain.Order{
		UserID:      userID,
		FromAmount:  amount,
		FromSymbol:  balance.Symbol,
		ToAmount:    0,
		ToSymbol:    symbol,
		ToAssetName: asset.Name,
		TakeProfit:  takeProfit,
		StopLoss:    stopLoss,
		Created:     time.Now(),
		Updated:     time.Now(),
		Status:      domain.OrderStatusNew,
	}

	if err := b.orderRepository.SaveOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("could not save new order, err: %w", err)
	}

	boughtAmount, err := b.converterService.Convert(ctx, userID, amount, balance.Symbol, symbol)
	if err != nil {
		return nil, fmt.Errorf(
			"could not convert %s to %s for user ID = '%s', err: %w",
			balance.Symbol,
			symbol,
			userID,
			err,
		)
	}

	order.ToAmount = boughtAmount
	order.ToSymbolPrice = amount / boughtAmount
	order.Updated = time.Now()
	order.Status = domain.OrderStatusActive
	if err := b.orderRepository.SaveOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("could not save order after conversion, err: %w", err)
	}

	balance, err = b.balanceService.GetUserActiveBalance(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s' after conversion, err: %w", userID, err)
	}

	return balance, nil
}
