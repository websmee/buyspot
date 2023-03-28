package background

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"

	"websmee/buyspot/internal/usecases"
)

type CurrentPricesUpdater struct {
	currentPricesRepository usecases.CurrentPricesRepository
	balanceService          usecases.BalanceService
	pricesService           usecases.PricesService
	logger                  *log.Logger
}

func NewCurrentPricesUpdater(
	currentPricesRepository usecases.CurrentPricesRepository,
	balanceService usecases.BalanceService,
	pricesService usecases.PricesService,
	logger *log.Logger,
) *CurrentPricesUpdater {
	return &CurrentPricesUpdater{
		currentPricesRepository,
		balanceService,
		pricesService,
		logger,
	}
}

func (u *CurrentPricesUpdater) Run(ctx context.Context) error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(time.Minute).Do(func() {
		u.logger.Println("updating current prices")

		symbols, err := u.balanceService.GetAvailableSymbols(ctx)
		if err != nil {
			u.logger.Println(fmt.Errorf("could not get available symbols, err: %w", err))
			return
		}

		for i := range symbols {
			prices, err := u.pricesService.GetCurrentPrices(ctx, symbols[i])
			if err != nil {
				u.logger.Println(
					fmt.Errorf("could not get current prices for %s, err: %w",
						symbols[i],
						err,
					),
				)
				continue
			}

			if err := u.currentPricesRepository.SaveCurrentPrices(ctx, prices, symbols[i]); err != nil {
				u.logger.Println(
					fmt.Errorf("could not save current prices for %s, err: %w",
						symbols[i],
						err,
					),
				)
			}
		}

		u.logger.Println("current prices updated")
	})

	s.StartAsync()

	return err
}
