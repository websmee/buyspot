package binance

import (
	"context"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"

	"websmee/buyspot/internal/domain"
)

type MarketDataStream struct {
}

func NewMarketDataStream(apiKey string) *MarketDataStream {
	return &MarketDataStream{}
}

func (s MarketDataStream) Subscribe(
	ctx context.Context,
	symbol string,
	interval domain.Interval,
	handler func(kline *domain.Kline),
	errorHandler func(err error),
) error {
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		open, _ := strconv.ParseFloat(event.Kline.Open, 64)
		cls, _ := strconv.ParseFloat(event.Kline.Close, 64)
		low, _ := strconv.ParseFloat(event.Kline.Low, 64)
		high, _ := strconv.ParseFloat(event.Kline.High, 64)
		volume, _ := strconv.ParseFloat(event.Kline.Volume, 64)
		quoteVolume, _ := strconv.ParseFloat(event.Kline.QuoteVolume, 64)
		activeBuyVolume, _ := strconv.ParseFloat(event.Kline.ActiveBuyVolume, 64)
		activeBuyQuoteVolume, _ := strconv.ParseFloat(event.Kline.ActiveBuyQuoteVolume, 64)

		handler(&domain.Kline{
			StartTime:            time.UnixMilli(event.Kline.StartTime),
			EndTime:              time.UnixMilli(event.Kline.EndTime),
			Open:                 open,
			Close:                cls,
			High:                 high,
			Low:                  low,
			Volume:               volume,
			TradeNum:             event.Kline.TradeNum,
			QuoteVolume:          quoteVolume,
			ActiveBuyVolume:      activeBuyVolume,
			ActiveBuyQuoteVolume: activeBuyQuoteVolume,
		})
	}

	doneC, stopC, err := binance.WsKlineServe(symbol, string(interval), wsKlineHandler, errorHandler)
	if err != nil {
		return err
	}

	stopC <- <-ctx.Done()
	<-doneC

	return nil
}
