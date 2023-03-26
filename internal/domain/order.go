package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            primitive.ObjectID `bson:"_id"`
	UserID        string             `bson:"user_id"`
	FromAmount    float64            `bson:"from_amount"`
	FromTicker    string             `bson:"from_ticker"`
	ToAmount      float64            `bson:"to_amount"`
	ToTicker      string             `bson:"to_ticker"`
	ToTickerPrice float64            `bson:"to_ticker_price"`
	ToAssetName   string             `bson:"to_asset_name"`
	TakeProfit    float64            `bson:"take_profit"`
	StopLoss      float64            `bson:"stop_loss"`
	Created       time.Time          `bson:"created"`
	Updated       time.Time          `bson:"updated"`
	CloseAmount   float64            `bson:"close_amount"`
	CloseTicker   string             `bson:"close_ticker"`
	Status        OrderStatus        `bson:"status"`
}

type OrderStatus string

const (
	OrderStatusNew    OrderStatus = "new"
	OrderStatusActive OrderStatus = "active"
	OrderStatusClosed OrderStatus = "closed"
)
