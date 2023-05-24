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

type AdviserCreator struct {
	trainer              *domain.AdviserTrainer
	assetRepository      usecases.AssetRepository
	marketDataRepository usecases.MarketDataRepository
	adviserRepository    usecases.AdviserRepository
	logger               *log.Logger
}

func NewAdviserCreator(
	assetRepository usecases.AssetRepository,
	marketDataRepository usecases.MarketDataRepository,
	adviserRepository usecases.AdviserRepository,
	logger *log.Logger,
) *AdviserCreator {
	trainer := domain.NewAdviserTrainer(
		[]any{ // initial parameters:
			3.0,  // forecastMultiplier
			3.0,  // minForecast
			5,    // rsiPeriod
			60.0, // rsiTrigger
			1.5,  // volumeSpikeRate
			1,    // volumeRiseInARow
		},
		func(parameters []any) *domain.Adviser { // create adviser from parameters:
			return &domain.Adviser{
				CheckHours:         24,
				ForecastHours:      8,
				ForecastMultiplier: parameters[0].(float64),
				MinForecast:        parameters[1].(float64),
				RSI:                domain.RSIIndicator{Period: parameters[2].(int), Trigger: parameters[3].(float64)},
				VolumeSpike:        domain.VolumeSpikeIndicator{Rate: parameters[4].(float64)},
				VolumeRise:         domain.VolumeRiseIndicator{InARow: parameters[5].(int)},
			}
		},
		func(parameters []any, index int) bool { // increment parameters:
			switch index {
			case 0: // forecastMultiplier
				if parameters[index].(float64) < 7 {
					parameters[index] = parameters[index].(float64) + 1
					return true
				}
			case 1: // minForecast
				if parameters[index].(float64) < 10 {
					parameters[index] = parameters[index].(float64) + 1
					return true
				}
			case 2: // rsiPeriod
				if parameters[index].(int) < 20 {
					parameters[index] = parameters[index].(int) + 5
					return true
				}
			case 3: // rsiTrigger
				if parameters[index].(float64) < 90 {
					parameters[index] = parameters[index].(float64) + 5
					return true
				}
			case 4: // volumeSpikeRate
				if parameters[index].(float64) < 4 {
					parameters[index] = parameters[index].(float64) + 0.5
					return true
				}
			case 5: // volumeRiseInARow
				if parameters[index].(int) < 4 {
					parameters[index] = parameters[index].(int) + 1
					return true
				}
			}

			return false
		},
		func(result *domain.AdviserTrainerResult) bool { // check result:
			return (result.SuccessRatePercent >= 45 && result.AdviceFrequencyPerDay > 5) ||
				(result.SuccessRatePercent >= 50 && result.AdviceFrequencyPerDay > 1) ||
				(result.SuccessRatePercent >= 55 && result.AdviceFrequencyPerDay > 0.5) ||
				(result.SuccessRatePercent >= 60 && result.AdviceFrequencyPerDay > 0.1)
		},
	)

	return &AdviserCreator{
		trainer,
		assetRepository,
		marketDataRepository,
		adviserRepository,
		logger,
	}
}

func (c *AdviserCreator) Run(ctx context.Context) error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(time.Hour * 24 * 7).Do(func() {
		c.logger.Println("creating new advisers")

		assets, err := c.assetRepository.GetAvailableAssets(ctx)
		if err != nil {
			c.logger.Println(fmt.Errorf("could not get available assets, err: %w", err))
			return
		}

		historyBySymbols := make(map[string][]domain.Kline)
		for _, asset := range assets {
			history, err := c.marketDataRepository.GetKlines(
				ctx,
				asset.Symbol,
				"USDT",
				time.Now().AddDate(0, 0, -30),
				time.Now(),
				domain.IntervalHour,
			)
			if err != nil {
				c.logger.Println(fmt.Errorf("could not get %s history, err: %w", asset.Symbol, err))
				return
			}

			historyBySymbols[asset.Symbol] = history
		}

		advisers := c.trainer.Train(historyBySymbols)
		if len(advisers) == 0 {
			c.logger.Println("no new advisers")
			return
		}

		if err := c.adviserRepository.MarkAllAdvisersAsNotLatest(ctx); err != nil {
			c.logger.Println(fmt.Errorf("could not mark advisers as not latest, err: %w", err))
			return
		}

		for i := range advisers {
			advisers[i].IsLatest = true
		}

		if err := c.adviserRepository.SaveAdvisers(ctx, advisers); err != nil {
			c.logger.Println(fmt.Errorf("could not save advisers, err: %w", err))
			return
		}

		c.logger.Println("new advisers created")
	})

	s.StartAsync()

	return err
}
