package domain

type Prices struct {
	Base            string             `json:"base"`
	PricesBySymbols map[string]float64 `json:"prices_by_symbols"`
}
