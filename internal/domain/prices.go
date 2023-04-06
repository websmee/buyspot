package domain

type Prices struct {
	Quote           string             `json:"quote"`
	PricesBySymbols map[string]float64 `json:"prices_by_symbols"`
}
