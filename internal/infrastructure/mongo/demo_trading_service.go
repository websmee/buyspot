package mongo

import (
	"context"
	"fmt"
	"strconv"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/infrastructure/redis"
)

type DemoTradingService struct {
	currentPricesRepository *redis.CurrentPricesRepository
	balanceService          *DemoBalanceService
}

func NewDemoTradingService(
	currentPricesRepository *redis.CurrentPricesRepository,
	balanceService *DemoBalanceService,
) *DemoTradingService {
	return &DemoTradingService{
		currentPricesRepository,
		balanceService,
	}
}

func (s *DemoTradingService) Buy(
	ctx context.Context,
	user *domain.User,
	balanceSymbol string,
	balanceAmount float64,
	tradeSymbol string,
) (string, error) {
	price, err := s.currentPricesRepository.GetPrice(ctx, tradeSymbol, balanceSymbol)
	if err != nil {
		return "0", fmt.Errorf("could not get price for %s%s, err: %w", tradeSymbol, balanceSymbol, err)
	}

	if err := s.balanceService.ChangeBalance(ctx, user.ID.Hex(), balanceSymbol, -balanceAmount); err != nil {
		return "0", fmt.Errorf(
			"could not change balance %s for user id = '%s' in mongo, err: %w",
			balanceSymbol,
			user.ID.Hex(),
			err,
		)
	}

	return fmt.Sprintf("%f", balanceAmount/price), nil
}

func (s *DemoTradingService) Sell(
	ctx context.Context,
	user *domain.User,
	tradeSymbol string,
	tradeAmount string,
	balanceSymbol string,
) (float64, error) {
	price, err := s.currentPricesRepository.GetPrice(ctx, tradeSymbol, balanceSymbol)
	if err != nil {
		return 0, fmt.Errorf("could not get price for %s%s, err: %w", tradeSymbol, balanceSymbol, err)
	}

	tradeAmountFloat, _ := strconv.ParseFloat(tradeAmount, 64)
	totalPrice := tradeAmountFloat * price
	if err := s.balanceService.ChangeBalance(ctx, user.ID.Hex(), balanceSymbol, totalPrice); err != nil {
		return 0, fmt.Errorf(
			"could not change balance %s for user id = '%s' in mongo, err: %w",
			balanceSymbol,
			user.ID.Hex(),
			err,
		)
	}

	return totalPrice, nil
}
