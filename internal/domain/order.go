package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID          primitive.ObjectID
	Amount      float64
	AssetTicker string
	Created     time.Time
}
