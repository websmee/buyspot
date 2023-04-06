package admin

import (
	"context"
	"fmt"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

type MarketDataUpdater struct {
	secretKey            string
	balanceService       usecases.BalanceService
	assetRepository      usecases.AssetRepository
	marketDataRepository usecases.MarketDataRepository
	marketDataService    usecases.MarketDataService
}

func NewMarketDataUpdater(
	secretKey string,
	balanceService usecases.BalanceService,
	assetRepository usecases.AssetRepository,
	marketDataRepository usecases.MarketDataRepository,
	marketDataService usecases.MarketDataService,
) *MarketDataUpdater {
	return &MarketDataUpdater{
		secretKey,
		balanceService,
		assetRepository,
		marketDataRepository,
		marketDataService,
	}
}

func (r *MarketDataUpdater) Update(ctx context.Context, secretKey string) error {
	if secretKey != r.secretKey {
		return domain.ErrForbidden
	}

	assets, err := r.assetRepository.GetAvailableAssets(ctx)
	if err != nil {
		return fmt.Errorf("could not get available assets, err: %w", err)
	}

	balanceSymbols, err := r.balanceService.GetAvailableSymbols(ctx)
	if err != nil {
		return fmt.Errorf("could not get balance symbols, err: %w", err)
	}

	for i := range balanceSymbols {
		for j := range assets {
			klines, err := r.marketDataService.GetMonth(ctx, assets[j].Symbol, balanceSymbols[i], domain.IntervalHour)
			if err != nil {
				return fmt.Errorf(
					"could not get %s%s market data, err: %w",
					assets[j].Symbol,
					balanceSymbols[i],
					err,
				)
			}

			for k := range klines {
				if err := r.marketDataRepository.CreateOrUpdate(
					ctx,
					assets[j].Symbol,
					balanceSymbols[i],
					domain.IntervalHour,
					&klines[k],
				); err != nil {
					return fmt.Errorf(
						"could not save %s%s kline, err: %w",
						assets[j].Symbol,
						balanceSymbols[i],
						err,
					)
				}
			}
		}
	}

	return nil
}
