package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Candlestick struct {
	ID          primitive.ObjectID
	Open        float64
	Low         float64
	High        float64
	Close       float64
	AdjClose    float64
	Volume      int64
	Timestamp   time.Time
	Interval    Interval
	AssetTicker string
}

type Interval string

const (
	IntervalHour Interval = "1h"
)
