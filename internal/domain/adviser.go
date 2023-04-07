package domain

import (
	"context"
)

type Adviser struct {
	priceRiseInARow    int
	volumeRiseInARow   int
	volumeSpikePercent float64
	minForecast        float64
	maxForecast        float64
}

func NewAdviser(
	priceRiseInARow int,
	volumeRiseInARow int,
	volumeSpikePercent float64,
	minForecast float64,
	maxForecast float64,
) *Adviser {
	return &Adviser{
		priceRiseInARow,
		volumeRiseInARow,
		volumeSpikePercent,
		minForecast,
		maxForecast,
	}
}

func (r *Adviser) GetAdvice(_ context.Context, marketData []Kline) (*Advice, error) {
	forecast := r.getPriceForecast(marketData)
	hours := r.getForecastHours(marketData)
	if r.checkIfPriceRises(marketData) &&
		r.checkIfVolumeRises(marketData) &&
		r.checkVolumeSpike(marketData) &&
		forecast > r.minForecast &&
		forecast < r.maxForecast {
		return &Advice{
			PriceForecast: forecast,
			ForecastHours: hours,
			BuyOrderSettings: BuyOrderSettings{
				Amount:            100,
				TakeProfit:        forecast,
				TakeProfitOptions: []float64{forecast - 1, forecast, forecast + 1, forecast + 2, forecast + 3},
				StopLoss:          -forecast,
				StopLossOptions:   []float64{-forecast + 1, -forecast, -forecast - 1, -forecast - 2, -forecast - 3},
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

func (r *Adviser) getPriceForecast(marketData []Kline) float64 {
	diff := marketData[len(marketData)-1].Close - marketData[len(marketData)-r.priceRiseInARow].Close
	return diff / marketData[len(marketData)-r.priceRiseInARow].Close * 100
}

func (r *Adviser) getForecastHours(_ []Kline) int {
	return r.priceRiseInARow * 2
}

func (r *Adviser) checkIfPriceRises(marketData []Kline) bool {
	for i := len(marketData) - 1; i > len(marketData)-r.priceRiseInARow; i-- {
		if marketData[i].Close < marketData[i-1].Close {
			return false
		}
		if marketData[i].High < marketData[i-1].High {
			return false
		}
	}

	return true
}

func (r *Adviser) checkIfVolumeRises(marketData []Kline) bool {
	for i := len(marketData) - 1; i > len(marketData)-r.volumeRiseInARow; i-- {
		if marketData[i].Volume < marketData[i-1].Volume {
			return false
		}
	}

	return true
}

func (r *Adviser) checkVolumeSpike(marketData []Kline) bool {
	diff := marketData[len(marketData)-1].Volume - marketData[len(marketData)-2].Volume
	return diff/marketData[len(marketData)-2].Volume*100 >= r.volumeSpikePercent
}
