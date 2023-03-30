package usecases

import (
	"context"
	"fmt"

	"websmee/buyspot/internal/domain"
)

type PricesReader struct {
	assetRepository         AssetRepository
	currentPricesRepository CurrentPricesRepository
	balanceService          BalanceService
}

func NewPricesReader(
	assetRepository AssetRepository,
	currentPricesRepository CurrentPricesRepository,
	balanceService BalanceService,
) *PricesReader {
	return &PricesReader{
		assetRepository,
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

	assets, err := r.assetRepository.GetAvailableAssets(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get available assets, err: %w", err)
	}

	var symbols []string
	for i := range assets {
		symbols = append(symbols, assets[i].Symbol)
	}
	symbols = append(symbols, balance.Symbol)

	return r.currentPricesRepository.GetPrices(ctx, symbols, balance.Symbol)
}
