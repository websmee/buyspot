package domain

type Prices struct {
	InTicker        string             `json:"in_ticker"`
	PricesByTickers map[string]float64 `json:"prices_by_tickers"`
}
