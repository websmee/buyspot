package indicator

import "websmee/buyspot/internal/domain"

type True struct {
}

func NewTrue() *True {
	return &True{}
}

func (t *True) Check(klines []domain.Kline) bool {
	return true
}
