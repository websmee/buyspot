package example

import (
	"context"

	"websmee/buyspot/internal/domain"
)

type BalanceService struct {
}

func NewBalanceService() *BalanceService {
	return &BalanceService{}
}

func (r *BalanceService) GetUserBalance(_ context.Context, _ *domain.User) (*domain.Balance, error) {
	return &domain.Balance{
		Amount: 1234.56,
		Ticker: "USDT",
	}, nil
}

func (r *BalanceService) GetAvailableTickers(_ context.Context) ([]string, error) {
	return []string{"USDT"}, nil
}
