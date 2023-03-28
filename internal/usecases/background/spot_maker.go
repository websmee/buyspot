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

type SpotMaker struct {
	currentSpotsRepository usecases.CurrentSpotsRepository
	marketDataRepository   usecases.MarketDataRepository
	newsRepository         usecases.NewsRepository
	assetRepository        usecases.AssetRepository
	adviser                usecases.Adviser
	logger                 *log.Logger
}

func NewSpotMaker(
	currentSpotsRepository usecases.CurrentSpotsRepository,
	marketDataRepository usecases.MarketDataRepository,
	newsRepository usecases.NewsRepository,
	assetRepository usecases.AssetRepository,
	adviser usecases.Adviser,
	orderRepository usecases.OrderRepository,
	logger *log.Logger,
) *SpotMaker {
	return &SpotMaker{
		currentSpotsRepository,
		marketDataRepository,
		newsRepository,
		assetRepository,
		adviser,
		logger,
	}
}

func (m *SpotMaker) Run(ctx context.Context) error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(time.Minute).Do(func() {
		m.logger.Println("making new spots")

		spots, err := m.makeSpots(ctx)
		if err != nil {
			m.logger.Println(fmt.Errorf("could not get new spots, err: %w", err))
			return
		}

		err = m.currentSpotsRepository.SaveSpots(ctx, spots, time.Minute)
		if err != nil {
			m.logger.Println(fmt.Errorf("could not save new spots, err: %w", err))
			return
		}

		m.logger.Println("new spots saved")
	})

	s.StartAsync()

	return err
}

func (m *SpotMaker) makeSpots(ctx context.Context) ([]domain.Spot, error) {
	assets, err := m.assetRepository.GetAvailableAssets(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get available assets, err: %w", err)
	}

	var spots []domain.Spot
	for i := range assets {
		advice, err := m.adviser.GetAdviceBySymbol(ctx, assets[i].Symbol)
		if err != nil {
			m.logger.Println(fmt.Errorf("could not find spot by symbol %s, err: %w", assets[i].Symbol, err))
			continue
		}

		historyMarketData, err := m.marketDataRepository.GetMonth(ctx, assets[i].Symbol, domain.IntervalHour)
		if err != nil {
			m.logger.Println(fmt.Errorf("could not get market data by symbol %s, err: %w", assets[i].Symbol, err))
			continue
		}

		forecastMarketData := buildForecastHours(
			ctx,
			assets[i].Symbol,
			historyMarketData[len(historyMarketData)-1].High,
			advice.PriceForecast,
			advice.ForecastHours,
		)

		news, err := m.newsRepository.GetFreshNewsBySymbol(ctx, assets[i].Symbol, time.Now().Add(-24*30*time.Hour))
		if err != nil {
			m.logger.Println(fmt.Errorf("could not get news by symbol %s, err: %w", assets[i].Symbol, err))
			continue
		}

		spots = append(spots, domain.Spot{
			Asset:              &assets[i],
			Advice:             advice,
			HistoryMarketData:  historyMarketData,
			ForecastMarketData: forecastMarketData,
			News:               news,
		})
	}

	return spots, nil
}

func buildForecastHours(
	_ context.Context,
	assetSymbol string,
	currentPrice float64,
	priceForecast float64,
	hours int,
) []domain.Kline {
	price := currentPrice
	endPrice := currentPrice + (currentPrice * priceForecast / 100)
	diff := (endPrice - currentPrice) / float64(hours)
	curvature := diff * 0.9
	var klines []domain.Kline
	for i := 1; i <= hours; i++ {
		klines = append(
			klines,
			getForecastKline(price, 0, time.Now().Add(time.Duration(i)*time.Hour)),
		)
		p := float64(i) / float64(hours)
		price += diff - curvature + (2 * curvature * p)
	}

	return klines
}

func getForecastKline(price float64, volume float64, t time.Time) domain.Kline {
	return domain.Kline{
		Open:    price,
		Low:     price,
		High:    price,
		Close:   price,
		Volume:  volume,
		EndTime: t,
	}
}
