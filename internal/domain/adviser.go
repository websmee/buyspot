package domain

import (
	"context"
)

type Adviser struct {
	riseInARow  int
	minForecast float64
}

func NewAdviser(riseInARow int, minForecast float64) *Adviser {
	return &Adviser{riseInARow, minForecast}
}

func (r *Adviser) GetAdvice(ctx context.Context, marketData []Kline) (*Advice, error) {
	rise, forecast, hours := checkIfPricesRise(marketData, r.riseInARow)
	if rise && forecast > r.minForecast {
		return &Advice{
			PriceForecast: forecast,
			ForecastHours: hours,
			BuyOrderSettings: BuyOrderSettings{
				Amount:            100,
				TakeProfit:        forecast,
				TakeProfitOptions: []float64{forecast, forecast + 1, forecast + 2, forecast + 3, forecast + 4},
				StopLoss:          -forecast,
				StopLossOptions:   []float64{-forecast, -forecast - 1, -forecast - 2, -forecast - 3, -forecast - 4},
			},
		}, nil
	}

	return nil, nil
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

func checkIfPricesRise(marketData []Kline, inARow int) (rise bool, forecast float64, hours int) {
	diff := marketData[len(marketData)-1].Close - marketData[len(marketData)-inARow].Close
	forecast = diff / marketData[len(marketData)-inARow].Close * 100
	hours = inARow * 2
	for i := len(marketData) - 1; i > len(marketData)-inARow; i-- {
		if marketData[i].Close < marketData[i-1].Close {
			return false, 0, 0
		}
		if marketData[i].High < marketData[i-1].High {
			return false, 0, 0
		}
	}

	return true, forecast, hours
}
