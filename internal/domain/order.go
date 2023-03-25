package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID          primitive.ObjectID
	FromAmount  float64
	FromTicker  string
	ToAmount    float64
	ToTicker    string
	ToAssetName string
	TakeProfit  float64
	StopLoss    float64
	Created     time.Time
}
