package domain

type VolumeRiseIndicator struct {
	InARow int `bson:"in_a_row"`
}

func (v *VolumeRiseIndicator) Check(klines []Kline) bool {
	if v.InARow <= 1 {
		return true
	}

	if v.InARow > len(klines) {
		return false
	}

	current := v.InARow
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
