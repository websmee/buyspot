package domain

import (
	"time"
)

type Kline struct {
	StartTime            time.Time `bson:"start_time"`
	EndTime              time.Time `bson:"end_time"`
	Open                 float64   `bson:"open"`
	Close                float64   `bson:"close"`
	High                 float64   `bson:"high"`
	Low                  float64   `bson:"low"`
	Volume               float64   `bson:"volume"`
	TradeNum             int64     `bson:"trade_num"`
	QuoteVolume          float64   `bson:"quote_volume"`
	ActiveBuyVolume      float64   `bson:"active_buy_volume"`
	ActiveBuyQuoteVolume float64   `bson:"active_buy_quote_volume"`
}

type Interval string

const (
	IntervalHour Interval = "1h"
)
