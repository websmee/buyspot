package api

import (
	"time"
)

type Balance struct {
	Ticker string  `json:"ticker"`
	Amount float64 `json:"amount"`
}

type Asset struct {
	Name        string `json:"name"`
	Ticker      string `json:"ticker"`
	Description string `json:"description"`
}

type ChartsData struct {
	Times    []string  `json:"times"`
	Prices   []float64 `json:"prices"`
	Forecast []float64 `json:"forecast"`
	Volumes  []int64   `json:"volumes"`
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

type Option struct {
	Value float64 `json:"value"`
	Text  string  `json:"text"`
}

type BuyOrderSettings struct {
	Amount            float64  `json:"amount"`
	TakeProfit        float64  `json:"takeProfit"`
	TakeProfitOptions []Option `json:"takeProfitOptions"`
	StopLoss          float64  `json:"stopLoss"`
	StopLossOptions   []Option `json:"stopLossOptions"`
}

type Spot struct {
	Asset             Asset            `json:"asset"`
	ActiveOrders      int              `json:"activeOrders"`
	PriceForecast     float64          `json:"priceForecast"`
	ChartsData        ChartsData       `json:"chartsData"`
	CurrentSpotsIndex int              `json:"currentSpotsIndex"`
	CurrentSpotsTotal int              `json:"currentSpotsTotal"`
	News              []NewsArticle    `json:"news"`
	BuyOrderSettings  BuyOrderSettings `json:"buyOrderSettings"`
}

type BuySpotResponse struct {
	UpdatedBalance Balance `json:"updatedBalance"`
}

type SpotsData struct {
	CurrentSpotsTotal int `json:"currentSpotsTotal"`
}

type Order struct {
	ID           string    `json:"id"`
	Amount       float64   `json:"amount"`
	AmountTicker string    `json:"amountTicker"`
	Asset        *Asset    `json:"asset"`
	PnL          float64   `json:"pnl"`
	TakeProfit   float64   `json:"takeProfit"`
	StopLoss     float64   `json:"stopLoss"`
	Created      time.Time `json:"created"`
}
