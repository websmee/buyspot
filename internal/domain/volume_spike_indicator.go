package domain

type VolumeSpikeIndicator struct {
	Rate float64 `bson:"rate"`
}

func (v *VolumeSpikeIndicator) Check(klines []Kline) bool {
	if len(klines) < 2 {
		return false
	}

	previous := klines[len(klines)-2]
	current := klines[len(klines)-1]

	return current.Volume/previous.Volume > v.Rate
}
