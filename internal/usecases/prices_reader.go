package usecases

import (
	"context"
	"fmt"

	"websmee/buyspot/internal/domain"
)

type PricesReader struct {
	userRepository          UserRepository
	assetRepository         AssetRepository
	currentPricesRepository CurrentPricesRepository
	balanceService          BalanceService
}

func NewPricesReader(
	userRepository UserRepository,
	assetRepository AssetRepository,
	currentPricesRepository CurrentPricesRepository,
	balanceService BalanceService,
) *PricesReader {
	return &PricesReader{
		userRepository,
		assetRepository,
		currentPricesRepository,
		balanceService,
	}
}

func (r *PricesReader) GetCurrentPrices(ctx context.Context) (*domain.Prices, error) {
	userID := domain.GetCtxUserID(ctx)
	if userID == "" {
		return nil, domain.ErrUnauthorized
	}

	// todo: use cashed balance
	//user, err := r.userRepository.GetByID(ctx, userID)
	//if err != nil {
	//	return nil, fmt.Errorf("could not get user by ID = '%s', err: %w", userID, err)
	//}

	//balance, err := r.balanceService.GetUserActiveBalance(ctx, user)
	//if err != nil {
	//	return nil, fmt.Errorf("could not get balance for user ID = '%s', err: %w", userID, err)
	//}
	balance := &domain.Balance{
		Symbol: "USDT",
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
