package domain

type Spot struct {
	Asset              *Asset        `json:"asset"`
	ActiveOrders       int           `json:"active_orders"`
	Advice             *Advice       `json:"advice"`
	HistoryMarketData  []Kline       `json:"history_market_data"`
	ForecastMarketData []Kline       `json:"forecast_market_data"`
	News               []NewsArticle `json:"news"`
}
