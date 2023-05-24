package domain

import (
	"github.com/cinar/indicator"
)

type RSIIndicator struct {
	Period  int     `bson:"period"`
	Trigger float64 `bson:"trigger"`
}

func (r *RSIIndicator) Check(klines []Kline) bool {
	closing := make([]float64, len(klines))
	for i := range klines {
		closing[i] = klines[i].Close
	}

	_, rsi := indicator.RsiPeriod(r.Period, closing[len(closing)-r.Period:])
	lastRSI := rsi[len(rsi)-1]

	return lastRSI >= r.Trigger
}
