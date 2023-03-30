package binance

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"

	"websmee/buyspot/internal/domain"
)

type MarketDataService struct {
	client *binance.Client
}

func NewMarketDataService(client *binance.Client) *MarketDataService {
	return &MarketDataService{client}
}

func (s *MarketDataService) GetMonth(
	ctx context.Context,
	symbol string,
	base string,
	interval domain.Interval,
) ([]domain.Kline, error) {
	klines, err := s.client.NewKlinesService().
		Symbol(symbol + base).
		Interval(string(interval)).
		StartTime(time.Now().AddDate(0, -1, 0).UnixMilli()).
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get %s%s klines from binance, err: %w", symbol, base, err)
	}

	var result []domain.Kline
	for i := range klines {
		open, _ := strconv.ParseFloat(klines[i].Open, 64)
		cls, _ := strconv.ParseFloat(klines[i].Close, 64)
		low, _ := strconv.ParseFloat(klines[i].Low, 64)
		high, _ := strconv.ParseFloat(klines[i].High, 64)
		volume, _ := strconv.ParseFloat(klines[i].Volume, 64)

		result = append(result, domain.Kline{
			StartTime: time.UnixMilli(klines[i].OpenTime),
			EndTime:   time.UnixMilli(klines[i].CloseTime),
			Open:      open,
			Close:     cls,
			High:      high,
			Low:       low,
			Volume:    volume,
		})
	}

	return result, nil
}
