package binance

import (
	"context"
	"fmt"
	"strconv"

	"github.com/adshao/go-binance/v2"

	"websmee/buyspot/internal/domain"
)

type BalanceService struct {
}

func NewBalanceService() *BalanceService {
	return &BalanceService{}
}

func (s *BalanceService) GetUserActiveBalance(ctx context.Context, user *domain.User) (*domain.Balance, error) {
	if user.BinanceAPIKey == "" || user.BinanceSecretKey == "" {
		return nil, fmt.Errorf("user '%s' does not have binance api access", user.ID.Hex())
	}

	acc, err := binance.NewClient(user.BinanceAPIKey, user.BinanceSecretKey).NewGetAccountService().Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get account info from binance, err: %w", err)
	}

	for _, binanceBalance := range acc.Balances {
		if binanceBalance.Asset == "USDT" {
			amount, _ := strconv.ParseFloat(binanceBalance.Free, 64)
			return &domain.Balance{
				UserID:   user.ID.Hex(),
				Symbol:   binanceBalance.Asset,
				Amount:   amount,
				IsActive: true,
			}, nil
		}
	}

	return nil, fmt.Errorf("could not find USDT balance in binance, err: %w", err)
}

func (s *BalanceService) GetUserBalances(ctx context.Context, user *domain.User) ([]domain.Balance, error) {
	balance, err := s.GetUserActiveBalance(ctx, user)
	if err != nil {
		return nil, err
	}

	return []domain.Balance{*balance}, nil
}

func (s *BalanceService) GetAvailableSymbols(_ context.Context) ([]string, error) {
	return []string{"USDT"}, nil
}
