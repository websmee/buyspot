package usecases

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"websmee/buyspot/internal/domain"
)

type SpotBuyer struct {
	userRepository     UserRepository
	orderRepository    OrderRepository
	spotRepository     SpotRepository
	tradingService     TradingService
	demoTradingService TradingService
	balanceService     BalanceService
	demoBalanceService BalanceService
	assetRepository    AssetRepository
}

func NewSpotBuyer(
	userRepository UserRepository,
	orderRepository OrderRepository,
	spotRepository SpotRepository,
	tradingService TradingService,
	demoTradingService TradingService,
	balanceService BalanceService,
	demoBalanceService BalanceService,
	assetRepository AssetRepository,
) *SpotBuyer {
	return &SpotBuyer{
		userRepository,
		orderRepository,
		spotRepository,
		tradingService,
		demoTradingService,
		balanceService,
		demoBalanceService,
		assetRepository,
	}
}

func (b *SpotBuyer) BuySpot(ctx context.Context, spotID string, amount float64, symbol string, takeProfit, stopLoss float64) (*domain.Balance, error) {
	userID := domain.GetCtxUserID(ctx)
	if userID == "" {
		return nil, domain.ErrUnauthorized
	}

	user, err := b.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not get user by ID = '%s', err: %w", userID, err)
	}

	var balance *domain.Balance
	if user.IsDemo {
		balance, err = b.demoBalanceService.GetUserActiveBalance(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("could not get demo user active balance, err: %w", err)
		}
	} else {
		balance, err = b.balanceService.GetUserActiveBalance(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("could not get user active balance, err: %w", err)
		}
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
		SpotID:      spotID,
		FromAmount:  amount,
		FromSymbol:  balance.Symbol,
		ToAmount:    "0",
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

	var boughtAmount string
	if user.IsDemo {
		boughtAmount, err = b.demoTradingService.Buy(ctx, user, balance.Symbol, amount, symbol)
		if err != nil {
			return nil, fmt.Errorf(
				"could not buy %s for %s as demo user with ID = '%s', err: %w",
				order.ToSymbol,
				balance.Symbol,
				userID,
				err,
			)
		}
	} else {
		boughtAmount, err = b.tradingService.Buy(ctx, user, balance.Symbol, amount, symbol)
		if err != nil {
			return nil, fmt.Errorf(
				"could not buy %s for %s as user with ID = '%s', err: %w",
				order.ToSymbol,
				balance.Symbol,
				userID,
				err,
			)
		}
	}

	order.ToAmount = boughtAmount
	boughtAmountFloat, _ := strconv.ParseFloat(boughtAmount, 64)
	order.ToSymbolPrice = amount / boughtAmountFloat
	order.Updated = time.Now()
	order.Status = domain.OrderStatusActive
	if err := b.orderRepository.SaveOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("could not save order after conversion, err: %w", err)
	}

	balance, err = b.balanceService.GetUserActiveBalance(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s' after conversion, err: %w", userID, err)
	}

	return balance, nil
}
