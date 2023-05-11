package local

import (
	"context"
	"fmt"

	"websmee/buyspot/internal/infrastructure/mongo"
	"websmee/buyspot/internal/infrastructure/redis"
)

type TradingService struct {
	currentPricesRepository *redis.CurrentPricesRepository
	balanceService          *mongo.BalanceService
}

func NewTradingService(
	currentPricesRepository *redis.CurrentPricesRepository,
	balanceService *mongo.BalanceService,
) *TradingService {
	return &TradingService{
		currentPricesRepository,
		balanceService,
	}
}

func (s *TradingService) Buy(
	ctx context.Context,
	userID string,
	balanceSymbol string,
	balanceAmount float64,
	tradeSymbol string,
) (float64, error) {
	price, err := s.currentPricesRepository.GetPrice(ctx, tradeSymbol, balanceSymbol)
	if err != nil {
		return 0, fmt.Errorf("could not get price for %s%s, err: %w", tradeSymbol, balanceSymbol, err)
	}

	if err := s.balanceService.ChangeBalance(ctx, userID, balanceSymbol, -balanceAmount); err != nil {
		return 0, fmt.Errorf(
			"could not change balance %s for user id = '%s' in mongo, err: %w",
			balanceSymbol,
			userID,
			err,
		)
	}

	return balanceAmount / price, nil
}

func (s *TradingService) Sell(
	ctx context.Context,
	userID string,
	tradeSymbol string,
	tradeAmount float64,
	balanceSymbol string,
) (float64, error) {
	price, err := s.currentPricesRepository.GetPrice(ctx, tradeSymbol, balanceSymbol)
	if err != nil {
		return 0, fmt.Errorf("could not get price for %s%s, err: %w", tradeSymbol, balanceSymbol, err)
	}

	totalPrice := tradeAmount * price
	if err := s.balanceService.ChangeBalance(ctx, userID, balanceSymbol, totalPrice); err != nil {
		return 0, fmt.Errorf(
			"could not change balance %s for user id = '%s' in mongo, err: %w",
			balanceSymbol,
			userID,
			err,
		)
	}

	return totalPrice, nil
}
