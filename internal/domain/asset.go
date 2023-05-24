package domain

type Asset struct {
	Symbol       string            `bson:"symbol"`
	Order        int               `bson:"order"`
	IsAvailable  bool              `bson:"is_available"`
	Name         string            `bson:"name"`
	Description  string            `bson:"description"`
	ExchangeInfo AssetExchangeInfo `bson:"exchange_info"`
}

type AssetExchangeInfo struct {
	Symbol                string `bson:"symbol"`
	IsAvailableForTrading bool   `bson:"is_available_for_trading"`
	Data                  string `bson:"data"`
}
