package usecases

import (
	"context"
	"fmt"

	"websmee/buyspot/internal/domain"
)

type PricesReader struct {
	currentPricesRepository CurrentPricesRepository
	balanceService          BalanceService
}

func NewPricesReader(
	currentPricesRepository CurrentPricesRepository,
	balanceService BalanceService,
) *PricesReader {
	return &PricesReader{
		currentPricesRepository,
		balanceService,
	}
}

func (r *PricesReader) GetCurrentPrices(ctx context.Context) (*domain.Prices, error) {
	user := domain.GetCtxUser(ctx)
	if user == nil {
		return nil, domain.ErrUnauthorized
	}

	balance, err := r.balanceService.GetUserActiveBalance(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not get balance for user ID = '%s', err: %w", user.ID, err)
	}

	return r.currentPricesRepository.GetCurrentPrices(ctx, balance.Symbol)
}
