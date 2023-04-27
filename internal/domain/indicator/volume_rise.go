package indicator

import (
	"websmee/buyspot/internal/domain"
)

type VolumeRise struct {
	inARow int
}

func NewVolumeRise(inARow int) *VolumeRise {
	return &VolumeRise{inARow}
}

func (v *VolumeRise) Check(klines []domain.Kline) bool {
	current := v.inARow
	for i := len(klines) - 1; i >= 0; i-- {
		if i > 0 && klines[i].Volume < klines[i-1].Volume {
			return false
		}
		current--
		if current <= 0 {
			break
		}
	}

	return true
}
