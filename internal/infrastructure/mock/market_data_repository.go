package mock

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"websmee/buyspot/internal/domain"
)

type MarketDataRepository struct {
}

func NewMarketDataRepository() *MarketDataRepository {
	return &MarketDataRepository{}
}

func (r *MarketDataRepository) GetAllHours(_ context.Context, assetTicker string) ([]domain.Candlestick, error) {
	prices := []float64{
		21234.12, 21224.23, 21214.56, 21264.78, 21214.90, 21134.12, 21154.34, 21164.56,
		21174.56, 21184.56, 21214.56, 21224.56, 21234.56, 21244.56, 21264.56, 21284.56,
		21319.56,
	}

	volumes := []int64{
		5000, 4000, 6000, 5000, 6000, 4000, 3000, 2000,
		5000, 4000, 6000, 5000, 6000, 7000, 8000, 9000,
		10000,
	}

	var candlesticks []domain.Candlestick
	for i := range prices {
		candlesticks = append(
			candlesticks,
			getHourCandlestick(assetTicker, prices[i], volumes[i], time.Now().Add(time.Duration(i-24)*time.Hour)),
		)
	}

	return candlesticks, nil
}

func getHourCandlestick(assetTicker string, price float64, volume int64, t time.Time) domain.Candlestick {
	return domain.Candlestick{
		ID:          primitive.NewObjectID(),
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
