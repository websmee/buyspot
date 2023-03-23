package domain

type Spot struct {
	Asset              *Asset        `json:"asset"`
	ActiveOrders       int           `json:"active_orders"`
	Advice             *Advice       `json:"advice"`
	HistoryMarketData  []Candlestick `json:"history_market_data"`
	ForecastMarketData []Candlestick `json:"forecast_market_data"`
	News               []NewsArticle `json:"news"`
}
