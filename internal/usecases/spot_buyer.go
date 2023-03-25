package usecases

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

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

func (b *SpotBuyer) BuySpot(ctx context.Context, amount float64, ticker string, takeProfit, stopLoss float64) (*domain.Balance, error) {
	user := domain.GetCtxUser(ctx)
	if user == nil {
		return nil, domain.ErrUnauthorized
	}

	balance, err := b.balanceService.GetUserBalance(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s', err: %w", user.ID, err)
	}

	asset, err := b.assetRepository.GetAssetByTicket(ctx, ticker)
	if err != nil {
		return nil, fmt.Errorf("could not available assets, err: %w", err)
	}
	if asset == nil {
		return nil, fmt.Errorf("could not find asset by ticker %s", ticker)
	}

	order := &domain.Order{
		ID:          primitive.NewObjectID(),
		FromAmount:  amount,
		FromTicker:  balance.Ticker,
		ToAmount:    0,
		ToTicker:    ticker,
		ToAssetName: asset.Name,
		TakeProfit:  takeProfit,
		StopLoss:    stopLoss,
		Created:     time.Now(),
		Status:      domain.OrderStatusNew,
	}

	if err := b.orderRepository.SaveOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("could not save new order, err: %w", err)
	}

	boughtAmount, err := b.converterService.Convert(ctx, amount, balance.Ticker, ticker)
	if err != nil {
		return nil, fmt.Errorf(
			"could not convert %s to %s for user ID = '%s', err: %w",
			balance.Ticker,
			ticker,
			user.ID,
			err,
		)
	}

	order.ToAmount = boughtAmount
	order.Status = domain.OrderStatusActive
	if err := b.orderRepository.SaveOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("could not save order after conversion, err: %w", err)
	}

	balance, err = b.balanceService.GetUserBalance(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s' after conversion, err: %w", user.ID, err)
	}

	return balance, nil
}
