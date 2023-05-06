package domain

import (
	"context"
	"math"
)

type Indicator interface {
	Check(klines []Kline) bool
}

type Adviser struct {
	checkHours         int
	forecastHours      int
	forecastMultiplier float64
	indicators         []Indicator
}

func NewAdviser(
	checkHours int,
	forecastHours int,
	forecastMultiplier float64,
	indicators ...Indicator,
) *Adviser {
	return &Adviser{
		checkHours,
		forecastHours,
		forecastMultiplier,
		indicators,
	}
}

func (r *Adviser) GetAdvice(_ context.Context, marketData []Kline) (*Advice, error) {
	if len(r.indicators) == 0 {
		return nil, nil
	}

	if len(marketData) < r.checkHours {
		return nil, nil
	}

	for _, indicator := range r.indicators {
		if !indicator.Check(marketData[len(marketData)-r.checkHours:]) {
			return nil, nil
		}
	}

	forecast := r.GetForecast(marketData[len(marketData)-r.checkHours:])
	return &Advice{
		PriceForecast: forecast,
		ForecastHours: r.forecastHours,
		BuyOrderSettings: BuyOrderSettings{
			Amount:            100,
			TakeProfit:        forecast,
			TakeProfitOptions: []float64{forecast, forecast * 2, forecast * 4},
			StopLoss:          -forecast,
			StopLossOptions:   []float64{-forecast, -forecast * 2, -forecast * 4},
		},
	}, nil
}

func (r *Adviser) GetForecast(klines []Kline) float64 {
	sum := 0.0
	diffSum := 0.0
	max := 10
	for i := len(klines) - 1; i >= 0; i-- {
		if i > 0 && max > 0 {
			max--
			sum += klines[i].Close
			diffSum += math.Abs(klines[i-1].Close - klines[i].Close)
		}
	}

	avg := sum / float64(len(klines))
	forecast := diffSum / float64(len(klines)-1) / avg * 100
	forecast *= r.forecastMultiplier

	return math.Ceil(forecast)
}

func (r *Adviser) CheckAdvice(advice *Advice, followingMarketData []Kline) int {
	startPrice := followingMarketData[0].Open
	for i := range followingMarketData {
		highDiff := followingMarketData[i].High - startPrice
		lowDiff := followingMarketData[i].Low - startPrice
		highPercent := highDiff / startPrice * 100
		lowPercent := lowDiff / startPrice * 100
		isTakeProfit := highPercent >= advice.BuyOrderSettings.TakeProfit
		isStopLoss := lowPercent <= advice.BuyOrderSettings.StopLoss
		if isTakeProfit && !isStopLoss {
			return 1
		}
		if !isTakeProfit && isStopLoss {
			return -1
		}
		if isTakeProfit && isStopLoss {
			return 0
		}
	}

	return 0
}
