package domain

type Spot struct {
	Asset                      *Asset             `json:"asset"`
	ActiveOrders               int                `json:"active_orders"`
	Advice                     *Advice            `json:"advice"`
	HistoryMarketDataByQuotes  map[string][]Kline `json:"history_market_data"`
	ForecastMarketDataByQuotes map[string][]Kline `json:"forecast_market_data"`
	News                       []NewsArticle      `json:"news"`
}
