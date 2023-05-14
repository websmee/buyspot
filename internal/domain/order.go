package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            primitive.ObjectID `bson:"_id"`
	UserID        string             `bson:"user_id"`
	SpotID        string             `bson:"spot_id"`
	FromAmount    float64            `bson:"from_amount"`
	FromSymbol    string             `bson:"from_symbol"`
	ToAmount      float64            `bson:"to_amount"`
	ToSymbol      string             `bson:"to_symbol"`
	ToSymbolPrice float64            `bson:"to_symbol_price"`
	ToAssetName   string             `bson:"to_asset_name"`
	TakeProfit    float64            `bson:"take_profit"`
	StopLoss      float64            `bson:"stop_loss"`
	Created       time.Time          `bson:"created"`
	Updated       time.Time          `bson:"updated"`
	CloseAmount   float64            `bson:"close_amount"`
	CloseSymbol   string             `bson:"close_symbol"`
	Status        OrderStatus        `bson:"status"`
}

type OrderStatus string

const (
	OrderStatusNew    OrderStatus = "new"
	OrderStatusActive OrderStatus = "active"
	OrderStatusClosed OrderStatus = "closed"
)
