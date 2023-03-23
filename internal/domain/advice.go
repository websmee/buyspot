package domain

type Advice struct {
	PriceForecast    float64          `json:"price_forecast"`
	ForecastHours    int              `json:"forecast_hours"`
	BuyOrderSettings BuyOrderSettings `json:"buy_order_settings"`
}

type BuyOrderSettings struct {
	Amount            float64   `json:"amount"`
	TakeProfit        float64   `json:"take_profit"`
	TakeProfitOptions []float64 `json:"take_profit_options"`
	StopLoss          float64   `json:"stop_loss"`
	StopLossOptions   []float64 `json:"stop_loss_options"`
}
