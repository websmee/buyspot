package domain

type Prices struct {
	InSymbol        string             `json:"in_symbol"`
	PricesBySymbols map[string]float64 `json:"prices_by_symbols"`
}
