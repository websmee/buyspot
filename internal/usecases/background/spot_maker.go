package background

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-co-op/gocron"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

type SpotMaker struct {
	balanceService         usecases.BalanceService
	currentSpotsRepository usecases.CurrentSpotsRepository
	marketDataRepository   usecases.MarketDataRepository
	newsRepository         usecases.NewsRepository
	assetRepository        usecases.AssetRepository
	adviser                usecases.Adviser
	userRepository         usecases.UserRepository
	notifier               usecases.NewSpotsNotifier
	logger                 *log.Logger
}

func NewSpotMaker(
	balanceService usecases.BalanceService,
	currentSpotsRepository usecases.CurrentSpotsRepository,
	marketDataRepository usecases.MarketDataRepository,
	newsRepository usecases.NewsRepository,
	assetRepository usecases.AssetRepository,
	adviser usecases.Adviser,
	userRepository usecases.UserRepository,
	notifier usecases.NewSpotsNotifier,
	logger *log.Logger,
) *SpotMaker {
	return &SpotMaker{
		balanceService,
		currentSpotsRepository,
		marketDataRepository,
		newsRepository,
		assetRepository,
		adviser,
		userRepository,
		notifier,
		logger,
	}
}

func (m *SpotMaker) Run(ctx context.Context) error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(time.Hour).Do(func() {
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

		m.notifyUsers(ctx, spots)

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

	balanceSymbols, err := m.balanceService.GetAvailableSymbols(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get available balance symbols, err: %w", err)
	}

	var spots []domain.Spot
	for i := range assets {
		adviceMarketData, err := m.marketDataRepository.GetKlines(
			ctx,
			assets[i].Symbol,
			"USDT",
			time.Now().AddDate(0, -1, 0),
			time.Now(),
			domain.IntervalHour,
		)
		if err != nil {
			m.logger.Println(fmt.Errorf(
				"could not get %sUSDT market data, err: %w",
				assets[i].Symbol,
				err,
			))
			continue
		}

		advice, err := m.adviser.GetAdvice(ctx, adviceMarketData)
		if err != nil {
			m.logger.Println(fmt.Errorf(
				"could not find spot for %s, err: %w",
				assets[i].Symbol,
				err,
			))
			continue
		}

		if advice == nil {
			continue
		}

		spot := m.GetSpot(ctx, advice, time.Now(), &assets[i], balanceSymbols)
		if spot == nil {
			continue
		}

		spots = append(spots, *spot)
	}

	return spots, nil
}

func (m *SpotMaker) notifyUsers(ctx context.Context, spots []domain.Spot) {
	if len(spots) == 0 {
		return
	}

	users, err := m.userRepository.GetUsers(ctx)
	if err != nil {
		m.logger.Println(fmt.Errorf("could not get users to notify, err: %w", err))
	}

	var wg sync.WaitGroup
	for i := range users {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err = m.notifier.Notify(ctx, &users[i], spots)
			if err != nil {
				m.logger.Println(fmt.Errorf("could not notify user %s, err: %w", users[i].Email, err))
			}
		}(i)
	}
	wg.Wait()

	return
}

func (m *SpotMaker) GetSpot(
	ctx context.Context,
	advice *domain.Advice,
	before time.Time,
	asset *domain.Asset,
	balanceSymbols []string,
) *domain.Spot {
	var err error
	historyMarketData := make(map[string][]domain.Kline, 0)
	forecastMarketData := make(map[string][]domain.Kline, 0)
	actualMarketData := make(map[string][]domain.Kline, 0)
	for j := range balanceSymbols {
		historyMarketData[balanceSymbols[j]], err = m.marketDataRepository.GetKlines(
			ctx,
			asset.Symbol,
			balanceSymbols[j],
			before.AddDate(0, -1, 0),
			before,
			domain.IntervalHour,
		)
		if err != nil {
			m.logger.Println(fmt.Errorf(
				"could not get %s%s market data, err: %w",
				asset.Symbol,
				balanceSymbols[j],
				err,
			))
			return nil
		}

		if len(historyMarketData[balanceSymbols[j]]) > 0 {
			forecastMarketData[balanceSymbols[j]] = buildForecastHours(
				ctx,
				historyMarketData[balanceSymbols[j]][len(historyMarketData[balanceSymbols[j]])-1].Close,
				advice.PriceForecast,
				before,
				advice.ForecastHours,
			)
		}
	}

	news, err := m.newsRepository.GetNewsBySymbol(ctx, asset.Symbol, time.Now().Add(-24*30*time.Hour), before)
	if err != nil {
		m.logger.Println(fmt.Errorf(
			"could not get %s news, err: %w",
			asset.Symbol,
			err,
		))
		return nil
	}

	return &domain.Spot{
		Asset:                      asset,
		Advice:                     advice,
		HistoryMarketDataByQuotes:  historyMarketData,
		ForecastMarketDataByQuotes: forecastMarketData,
		ActualMarketDataByQuotes:   actualMarketData,
		News:                       news,
	}
}

func buildForecastHours(
	_ context.Context,
	currentPrice float64,
	priceForecast float64,
	after time.Time,
	hours int,
) []domain.Kline {
	price := currentPrice
	endPrice := currentPrice + (currentPrice * priceForecast / 100)
	diff := (endPrice - currentPrice) / float64(hours)
	curvature := diff * 0.9
	var klines []domain.Kline
	for i := 0; i < hours; i++ {
		klines = append(
			klines,
			getForecastKline(price, 0, after.Add(time.Duration(i)*time.Hour)),
		)
		p := float64(i+1) / float64(hours)
		price += diff - curvature + (2 * curvature * p)
	}

	return klines
}

func getForecastKline(price float64, volume float64, startTime time.Time) domain.Kline {
	return domain.Kline{
		Open:      price,
		Low:       price,
		High:      price,
		Close:     price,
		Volume:    volume,
		StartTime: startTime,
		EndTime:   startTime.Add(time.Hour),
	}
}
