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

func (r *BalanceService) GetUserActiveBalance(_ context.Context, userID string) (*domain.Balance, error) {
	return &domain.Balance{
		Amount: 1234.56,
		Symbol: "USDT",
	}, nil
}

func (r *BalanceService) GetUserBalances(ctx context.Context, userID string) ([]domain.Balance, error) {
	return []domain.Balance{
		{
			Amount: 1234.56,
			Symbol: "USDT",
		},
	}, nil
}

func (r *BalanceService) GetAvailableSymbols(_ context.Context) ([]string, error) {
	return []string{"USDT"}, nil
}
