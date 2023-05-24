package domain

import (
	"math"
	"math/rand"
	"time"
)

type Adviser struct {
	CheckHours            int                  `bson:"check_hours"`
	ForecastHours         int                  `bson:"forecast_hours"`
	ForecastMultiplier    float64              `bson:"forecast_multiplier"`
	MinForecast           float64              `bson:"min_forecast"`
	RSI                   RSIIndicator         `bson:"rsi"`
	VolumeRise            VolumeRiseIndicator  `bson:"volume_rise"`
	VolumeSpike           VolumeSpikeIndicator `bson:"volume_spike"`
	AdviceFrequencyPerDay float64              `bson:"frequency_per_day"`
	SuccessRatePercent    int                  `bson:"success_rate_percent"`
	IsLatest              bool                 `bson:"is_latest"`
}

func (r *Adviser) GetAdvice(marketData []Kline) *Advice {
	if len(marketData) < r.CheckHours {
		return nil
	}

	if !r.RSI.Check(marketData[len(marketData)-r.CheckHours:]) {
		return nil
	}

	if !r.VolumeRise.Check(marketData[len(marketData)-r.CheckHours:]) {
		return nil
	}

	if !r.VolumeSpike.Check(marketData[len(marketData)-r.CheckHours:]) {
		return nil
	}

	forecast := r.GetForecast(marketData[len(marketData)-r.CheckHours:])
	if forecast < r.MinForecast {
		return nil
	}

	return &Advice{
		PriceForecast: forecast,
		ForecastHours: r.ForecastHours,
		BuyOrderSettings: BuyOrderSettings{
			Amount:            100,
			TakeProfit:        forecast,
			TakeProfitOptions: []float64{forecast, forecast * 2, forecast * 4},
			StopLoss:          -forecast,
			StopLossOptions:   []float64{-forecast, -forecast * 2, -forecast * 4},
		},
	}
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
	forecast *= r.ForecastMultiplier

	return math.Ceil(forecast)
}

func BuildForecastHours(
	currentPrice float64,
	priceForecast float64,
	after time.Time,
	hours int,
) []Kline {
	price := currentPrice
	endPrice := currentPrice + (currentPrice * priceForecast / 100)
	diff := (endPrice - currentPrice) / float64(hours)
	curvature := diff * 0.9
	noiseSum := 0.0
	var klines []Kline
	for i := 0; i < hours; i++ {
		klines = append(
			klines,
			getForecastKline(price, 0, after.Add(time.Duration(i)*time.Hour)),
		)
		p := float64(i+1) / float64(hours)
		noise := diff * rand.Float64() * 2
		if i%2 == 0 {
			noise *= -1
		}
		if i == hours-1 {
			noise = -noiseSum
		} else {
			noiseSum += noise
		}

		price += diff - curvature + (2 * curvature * p) + noise
	}

	return klines
}

func getForecastKline(price float64, volume float64, startTime time.Time) Kline {
	return Kline{
		Open:      price,
		Low:       price,
		High:      price,
		Close:     price,
		Volume:    volume,
		StartTime: startTime,
		EndTime:   startTime.Add(time.Hour),
	}
}
