package usecases

import (
	"context"
	"fmt"

	"websmee/buyspot/internal/domain"
)

type BalanceReader struct {
	balanceService BalanceService
}

func NewBalanceReader(
	balanceService BalanceService,
) *BalanceReader {
	return &BalanceReader{
		balanceService,
	}
}

func (r *BalanceReader) GetActiveBalance(ctx context.Context) (*domain.Balance, error) {
	userID := domain.GetCtxUserID(ctx)
	if userID == "" {
		return nil, domain.ErrUnauthorized
	}

	balance, err := r.balanceService.GetUserActiveBalance(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not get user active balance, err: %w", err)
	}

	return balance, nil
}
