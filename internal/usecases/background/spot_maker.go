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
	orderRepository        usecases.OrderRepository
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
		orderRepository,
		logger,
	}
}

func (m *SpotMaker) Run(ctx context.Context) error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(time.Minute).Do(func() {
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
		advice, err := m.adviser.GetAdviceByTicker(ctx, assets[i].Ticker)
		if err != nil {
			m.logger.Println(fmt.Errorf("could not find spot by ticker %s, err: %w", assets[i].Ticker, err))
			continue
		}

		historyMarketData, err := m.marketDataRepository.GetAllHours(ctx, assets[i].Ticker)
		if err != nil {
			m.logger.Println(fmt.Errorf("could not get market data by ticker %s, err: %w", assets[i].Ticker, err))
			continue
		}

		forecastMarketData := buildForecastHours(
			ctx,
			assets[i].Ticker,
			historyMarketData[len(historyMarketData)-1].High,
			advice.PriceForecast,
			advice.ForecastHours,
		)

		news, err := m.newsRepository.GetFreshNewsByTicker(ctx, assets[i].Ticker, time.Now().Add(-24*30*time.Hour))
		if err != nil {
			m.logger.Println(fmt.Errorf("could not get news by ticker %s, err: %w", assets[i].Ticker, err))
			continue
		}

		activeOrdersCount, err := m.orderRepository.GetUserActiveOrdersCountByTicker(ctx, assets[i].Ticker)
		if err != nil {
			m.logger.Println(
				fmt.Errorf(
					"could not get active orders count by ticker %s, err: %w",
					assets[i].Ticker,
					err,
				))
			continue
		}

		spots = append(spots, domain.Spot{
			Asset:              &assets[i],
			ActiveOrders:       activeOrdersCount,
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
	assetTicker string,
	currentPrice float64,
	priceForecast float64,
	hours int,
) []domain.Candlestick {
	price := currentPrice
	endPrice := currentPrice + (currentPrice * priceForecast / 100)
	diff := (endPrice - currentPrice) / float64(hours)
	curvature := diff * 0.9
	var candlesticks []domain.Candlestick
	for i := 1; i <= hours; i++ {
		candlesticks = append(
			candlesticks,
			getForecastCandlestick(assetTicker, price, 0, time.Now().Add(time.Duration(i)*time.Hour)),
		)
		p := float64(i) / float64(hours)
		price += diff - curvature + (2 * curvature * p)
	}

	return candlesticks
}

func getForecastCandlestick(assetTicker string, price float64, volume int64, t time.Time) domain.Candlestick {
	return domain.Candlestick{
		Open:        price,
		Low:         price,
		High:        price,
		Close:       price,
		AdjClose:    price,
		Volume:      volume,
		Timestamp:   t,
		Interval:    "1h",
		AssetTicker: assetTicker,
	}
}
