package background

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

const priceExpiration = time.Hour

type MarketDataUpdater struct {
	balanceService          usecases.BalanceService
	assetRepository         usecases.AssetRepository
	marketDataStream        usecases.MarketDataStream
	marketDataRepository    usecases.MarketDataRepository
	currentPricesRepository usecases.CurrentPricesRepository
	logger                  *log.Logger
	doneChs                 []chan struct{}
}

func NewMarketDataUpdater(
	balanceService usecases.BalanceService,
	assetRepository usecases.AssetRepository,
	marketDataStream usecases.MarketDataStream,
	marketDataRepository usecases.MarketDataRepository,
	currentPricesRepository usecases.CurrentPricesRepository,
	logger *log.Logger,
) *MarketDataUpdater {
	return &MarketDataUpdater{
		balanceService,
		assetRepository,
		marketDataStream,
		marketDataRepository,
		currentPricesRepository,
		logger,
		[]chan struct{}{},
	}
}

func (m *MarketDataUpdater) Run(ctx context.Context) error {
	assets, err := m.assetRepository.GetAvailableAssets(ctx)
	if err != nil {
		return err
	}

	balanceSymbols, err := m.balanceService.GetAvailableSymbols(ctx)
	if err != nil {
		return err
	}

	g := new(errgroup.Group)
	for i := range balanceSymbols {
		i := i
		for j := range assets {
			j := j
			g.Go(func() error {
				return m.marketDataStream.Subscribe(ctx, assets[j].Symbol, balanceSymbols[i], domain.IntervalHour,
					func(kline *domain.Kline) {
						if err := m.currentPricesRepository.UpdatePrice(
							ctx,
							kline.Close,
							assets[j].Symbol,
							balanceSymbols[i],
							priceExpiration,
						); err != nil {
							m.logger.Println(fmt.Errorf(
								"could not update %s%s price, err: %w",
								assets[j].Symbol,
								balanceSymbols[i],
								err,
							))
						}
						if err := m.marketDataRepository.CreateOrUpdate(
							ctx,
							assets[j].Symbol,
							balanceSymbols[i],
							domain.IntervalHour,
							kline,
						); err != nil {
							m.logger.Println(fmt.Errorf(
								"could not save %s%s kline, err: %w",
								assets[j].Symbol,
								balanceSymbols[i],
								err,
							))
						}
					},
					func(err error) {
						m.logger.Println(fmt.Errorf(
							"%s%s stream error: %w",
							assets[j].Symbol,
							balanceSymbols[i],
							err,
						))
					},
				)
			})
		}
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not subscribe to market data stream, err: %w", err)
	}

	return nil
}
