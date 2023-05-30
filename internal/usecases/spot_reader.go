package usecases

import (
	"context"
	"fmt"
	"time"

	"websmee/buyspot/internal/domain"
)

type SpotReader struct {
	currentSpotsRepository CurrentSpotsRepository
	orderRepository        OrderRepository
	marketDataRepository   MarketDataRepository
}

func NewSpotReader(
	currentSpotsRepository CurrentSpotsRepository,
	orderRepository OrderRepository,
	marketDataRepository MarketDataRepository,
) *SpotReader {
	return &SpotReader{
		currentSpotsRepository,
		orderRepository,
		marketDataRepository,
	}
}

func (r *SpotReader) GetSpotByIndex(ctx context.Context, index int) (*domain.Spot, error) {
	userID := domain.GetCtxUserID(ctx)
	if userID == "" {
		return nil, domain.ErrUnauthorized
	}

	spot, err := r.currentSpotsRepository.GetSpotByIndex(ctx, index)
	if err != nil {
		return nil, fmt.Errorf("could not get spot by index %d, err: %w", index, err)
	}

	spot.ActualMarketDataByQuotes = make(map[string][]domain.Kline)
	for i := range spot.HistoryMarketDataByQuotes {
		last := spot.HistoryMarketDataByQuotes[i][len(spot.HistoryMarketDataByQuotes[i])-1]
		hours := len(spot.ForecastMarketDataByQuotes[i])
		if last.EndTime.Before(time.Now()) {
			spot.ActualMarketDataByQuotes[i], err = r.marketDataRepository.GetKlines(
				ctx,
				spot.Asset.Symbol,
				i,
				last.EndTime.Add(-time.Hour),
				last.EndTime.Add(time.Duration(hours-1)*time.Hour),
				domain.IntervalHour,
			)
			if err != nil {
				return nil, fmt.Errorf(
					"could not get %s%s actual market data, err: %w",
					spot.Asset.Symbol,
					i,
					err,
				)
			}
		}
	}

	activeOrdersCount, err := r.orderRepository.GetUserActiveOrdersCountBySymbol(ctx, userID, spot.Asset.Symbol)
	if err != nil {
		return nil, fmt.Errorf(
			"could not get user (ID = '%s') active orders count by symbol %s, err: %w",
			userID,
			spot.Asset.Symbol,
			err,
		)
	}

	spot.ActiveOrders = activeOrdersCount

	return spot, nil
}

func (r *SpotReader) GetSpotsCount(ctx context.Context) (int, error) {
	return r.currentSpotsRepository.GetSpotsCount(ctx)
}
