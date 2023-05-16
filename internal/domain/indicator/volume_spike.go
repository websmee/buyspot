package indicator

import (
	"websmee/buyspot/internal/domain"
)

type VolumeSpike struct {
	rate float64
}

func NewVolumeSpike(rate float64) *VolumeSpike {
	return &VolumeSpike{rate}
}

func (v *VolumeSpike) Check(klines []domain.Kline) bool {
	if len(klines) < 2 {
		return false
	}

	previous := klines[len(klines)-2]
	current := klines[len(klines)-1]

	return current.Volume/previous.Volume > v.rate
}
