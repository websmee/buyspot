package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Spot struct {
	ID                 primitive.ObjectID `json:"id"`
	Asset              Asset              `json:"asset"`
	ActiveOrders       int64              `json:"active_orders"`
	PriceForecast      float64            `json:"price_forecast"`
	HistoryMarketData  []Candlestick      `json:"history_market_data"`
	ForecastMarketData []Candlestick      `json:"forecast_market_data"`
	News               []NewsArticle      `json:"news"`
	BuyOrderSettings   BuyOrderSettings   `json:"buy_order_settings"`
}

type Asset struct {
	Ticker      string `json:"ticker"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewsArticle struct {
	Sentiment NewsArticleSentiment `json:"sentiment"`
	Title     string               `json:"title"`
	Content   string               `json:"content"`
	Created   time.Time            `json:"created"`
	Views     int64                `json:"views"`
}

type NewsArticleSentiment string

const (
	NewsArticleSentimentNeutral  NewsArticleSentiment = "NEUTRAL"
	NewsArticleSentimentPositive NewsArticleSentiment = "POSITIVE"
	NewsArticleSentimentNegative NewsArticleSentiment = "NEGATIVE"
)

type BuyOrderSettings struct {
	Amount            float64   `json:"amount"`
	TakeProfit        float64   `json:"take_profit"`
	TakeProfitOptions []float64 `json:"take_profit_options"`
	StopLoss          float64   `json:"stop_loss"`
	StopLossOptions   []float64 `json:"stop_loss_options"`
}
