package example

import (
	"context"
	"time"

	"websmee/buyspot/internal/domain"
)

type MarketDataRepository struct {
}

func NewMarketDataRepository() *MarketDataRepository {
	return &MarketDataRepository{}
}

func (r *MarketDataRepository) GetMonth(_ context.Context, symbol string, interval domain.Interval) ([]domain.Kline, error) {
	prices := []float64{
		21234.12, 21224.23, 21214.56, 21264.78, 21214.90, 21134.12, 21154.34, 21164.56,
		21174.56, 21184.56, 21214.56, 21224.56, 21234.56, 21244.56, 21264.56, 21284.56,
		21319.56,
	}

	volumes := []float64{
		5000, 4000, 6000, 5000, 6000, 4000, 3000, 2000,
		5000, 4000, 6000, 5000, 6000, 7000, 8000, 9000,
		10000,
	}

	var klines []domain.Kline
	for i := range prices {
		klines = append(
			klines,
			getHourKline(prices[i], volumes[i], time.Now().Add(time.Duration(i-24)*time.Hour)),
		)
	}

	return klines, nil
}

func (r *MarketDataRepository) CreateOrUpdate(
	_ context.Context,
	symbol string,
	interval domain.Interval,
	kline *domain.Kline,
) error {
	return nil
}

func getHourKline(price float64, volume float64, t time.Time) domain.Kline {
	return domain.Kline{
		Open:    price,
		Low:     price,
		High:    price,
		Close:   price,
		Volume:  volume,
		EndTime: t,
	}
}
