package domain

type Advice struct {
	PriceForecast    float64          `json:"price_forecast" bson:"price_forecast"`
	ForecastHours    int              `json:"forecast_hours" bson:"forecast_hours"`
	BuyOrderSettings BuyOrderSettings `json:"buy_order_settings" bson:"buy_order_settings"`
	IsProfitable     bool             `json:"is_profitable" bson:"is_profitable"`
	Confidence       int              `json:"confidence" bson:"confidence"`
}

type BuyOrderSettings struct {
	Amount            float64   `json:"amount" bson:"amount"`
	TakeProfit        float64   `json:"take_profit" bson:"take_profit"`
	TakeProfitOptions []float64 `json:"take_profit_options" bson:"take_profit_options"`
	StopLoss          float64   `json:"stop_loss" bson:"stop_loss"`
	StopLossOptions   []float64 `json:"stop_loss_options" bson:"stop_loss_options"`
}
