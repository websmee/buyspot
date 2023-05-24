package background

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/go-co-op/gocron"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

type SpotMaker struct {
	balanceService         usecases.BalanceService
	currentSpotsRepository usecases.CurrentSpotsRepository
	SpotRepository         usecases.SpotRepository
	marketDataRepository   usecases.MarketDataRepository
	newsRepository         usecases.NewsRepository
	assetRepository        usecases.AssetRepository
	adviserRepository      usecases.AdviserRepository
	userRepository         usecases.UserRepository
	notifier               usecases.Notifier
	logger                 *log.Logger
}

func NewSpotMaker(
	balanceService usecases.BalanceService,
	currentSpotsRepository usecases.CurrentSpotsRepository,
	SpotRepository usecases.SpotRepository,
	marketDataRepository usecases.MarketDataRepository,
	newsRepository usecases.NewsRepository,
	assetRepository usecases.AssetRepository,
	adviserRepository usecases.AdviserRepository,
	userRepository usecases.UserRepository,
	notifier usecases.Notifier,
	logger *log.Logger,
) *SpotMaker {
	return &SpotMaker{
		balanceService,
		currentSpotsRepository,
		SpotRepository,
		marketDataRepository,
		newsRepository,
		assetRepository,
		adviserRepository,
		userRepository,
		notifier,
		logger,
	}
}

func (m *SpotMaker) Run(ctx context.Context) error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Cron("1 * * * *").Do(func() {
		m.logger.Println("making new spots")

		spots, err := m.makeSpots(ctx)
		if err != nil {
			m.logger.Println(fmt.Errorf("could not get new spots, err: %w", err))
			return
		}

		if len(spots) == 0 {
			m.logger.Println("no new spots")
			return
		}

		for i := range spots {
			err := m.SpotRepository.SaveSpot(ctx, &spots[i])
			if err != nil {
				m.logger.Println(fmt.Errorf("could not save spot to history, err: %w", err))
				return
			}
		}

		err = m.currentSpotsRepository.SaveSpots(ctx, spots, time.Hour-time.Minute)
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
	advisers, err := m.adviserRepository.GetLatestAdvisers(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get current advisers, err: %w", err)
	}
	sort.Slice(advisers, func(i, j int) bool {
		return advisers[i].SuccessRatePercent > advisers[j].SuccessRatePercent
	})

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

		var advice *domain.Advice
		for _, adviser := range advisers {
			advice = adviser.GetAdvice(adviceMarketData)
			if advice != nil {
				advice.Confidence = adviser.SuccessRatePercent
				break
			}
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
			err = m.notify(ctx, &users[i], spots)
			if err != nil {
				m.logger.Println(fmt.Errorf("could not notify user %s, err: %w", users[i].Email, err))
			}
		}(i)
	}
	wg.Wait()

	return
}

func (m *SpotMaker) notify(ctx context.Context, user *domain.User, spots []domain.Spot) error {
	if user.NotificationsKey == "" {
		return nil
	}

	if len(spots) == 0 {
		return nil
	}

	symbols := make([]string, 0, len(spots))
	for _, spot := range spots {
		symbols = append(symbols, fmt.Sprintf("%s %d%%", spot.Asset.Symbol, spot.Advice.Confidence))
	}

	return m.notifier.Notify(
		ctx,
		user,
		"NEW SPOTS",
		strings.Join(symbols, ", "),
	)
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
			forecastMarketData[balanceSymbols[j]] = domain.BuildForecastHours(
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
