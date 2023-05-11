package http

import (
	"time"
)

type Balance struct {
	Symbol string  `json:"symbol"`
	Amount float64 `json:"amount"`
}

type Asset struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

type ChartsData struct {
	Times    []string  `json:"times"`
	Prices   []float64 `json:"prices"`
	Forecast []float64 `json:"forecast"`
	Actual   []float64 `json:"actual"`
	Volumes  []int64   `json:"volumes"`
}

type NewsArticle struct {
	Sentiment NewsArticleSentiment `json:"sentiment"`
	Title     string               `json:"title"`
	Content   string               `json:"content"`
	Summary   string               `json:"summary"`
	URL       string               `json:"url"`
	ImgURL    string               `json:"imgURL"`
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
	Index              int                   `json:"index"`
	Asset              Asset                 `json:"asset"`
	ActiveOrders       int                   `json:"activeOrders"`
	PriceForecast      float64               `json:"priceForecast"`
	ChartsDataByQuotes map[string]ChartsData `json:"chartsDataByQuotes"`
	CurrentSpotsIndex  int                   `json:"currentSpotsIndex"`
	CurrentSpotsTotal  int                   `json:"currentSpotsTotal"`
	News               []NewsArticle         `json:"news"`
	BuyOrderSettings   BuyOrderSettings      `json:"buyOrderSettings"`
	IsProfitable       bool                  `json:"isProfitable"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BuySpotRequest struct {
	Amount     float64 `json:"amount"`
	Symbol     string  `json:"symbol"`
	TakeProfit float64 `json:"takeProfit"`
	StopLoss   float64 `json:"stopLoss"`
}

type BuySpotResponse struct {
	UpdatedBalance Balance `json:"updatedBalance"`
}

type SellOrderResponse struct {
	OrderID        string  `json:"orderID"`
	UpdatedBalance Balance `json:"updatedBalance"`
}

type SpotsData struct {
	CurrentSpotsTotal int `json:"currentSpotsTotal"`
}

type Order struct {
	ID          string    `json:"id"`
	FromAmount  float64   `json:"fromAmount"`
	FromSymbol  string    `json:"fromSymbol"`
	ToAmount    float64   `json:"toAmount"`
	ToSymbol    string    `json:"toSymbol"`
	ToAssetName string    `json:"toAssetName"`
	TakeProfit  float64   `json:"takeProfit"`
	StopLoss    float64   `json:"stopLoss"`
	Created     time.Time `json:"created"`
}

type Prices struct {
	Quote           string             `json:"quote"`
	PricesBySymbols map[string]float64 `json:"pricesBySymbols"`
}
