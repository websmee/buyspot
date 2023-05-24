package background

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

type AssetUpdater struct {
	assetRepository     usecases.AssetRepository
	exchangeInfoService usecases.ExchangeInfoService
	logger              *log.Logger
}

func NewAssetUpdater(
	assetRepository usecases.AssetRepository,
	exchangeInfoService usecases.ExchangeInfoService,
	logger *log.Logger,
) *AssetUpdater {
	return &AssetUpdater{
		assetRepository,
		exchangeInfoService,
		logger,
	}
}

func (u *AssetUpdater) Run(ctx context.Context) error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(time.Hour * 24 * 7).Do(func() {
		u.logger.Println("updating assets")

		assets, err := u.assetRepository.GetAvailableAssets(ctx)
		if err != nil {
			u.logger.Println(fmt.Errorf("could not get available assets, err: %w", err))
			return
		}

		for i := 0; i < len(assets); i = i + 10 {
			j := i + 10
			if i+10 > len(assets) {
				j = len(assets)
			}

			if err := u.updateAssets(ctx, assets[i:j]); err != nil {
				u.logger.Println(err)
				return
			}
		}

		u.logger.Println("assets updated")
	})

	s.StartAsync()

	return err
}

func (u *AssetUpdater) updateAssets(ctx context.Context, assets []domain.Asset) error {
	symbols := make([]string, len(assets))
	for i, asset := range assets {
		symbols[i] = asset.Symbol
	}

	exchangeInfo, err := u.exchangeInfoService.GetExchangeInfo(ctx, symbols)
	if err != nil {
		return fmt.Errorf("could not get exchange info, err: %w", err)
	}

	assetsMap := make(map[string]domain.Asset)
	for _, asset := range assets {
		assetsMap[asset.Symbol] = asset
	}

	for _, info := range exchangeInfo {
		if asset, ok := assetsMap[info.Symbol]; ok {
			asset.IsAvailable = info.IsAvailableForTrading
			asset.ExchangeInfo = info
			if err := u.assetRepository.CreateOrUpdate(ctx, &asset); err != nil {
				return fmt.Errorf("could not save asset, err: %w", err)
			}
		}
	}

	return nil
}
