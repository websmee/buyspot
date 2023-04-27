package indicator

import (
	"github.com/cinar/indicator"

	"websmee/buyspot/internal/domain"
)

type RSI struct {
	period  int
	trigger float64
}

func NewRSI(period int, trigger float64) *RSI {
	return &RSI{period, trigger}
}

func (r *RSI) Check(klines []domain.Kline) bool {
	closing := make([]float64, len(klines))
	for i := range klines {
		closing[i] = klines[i].Close
	}

	_, rsi := indicator.RsiPeriod(r.period, closing[len(closing)-r.period:])
	lastRSI := rsi[len(rsi)-1]

	return lastRSI >= r.trigger
}
