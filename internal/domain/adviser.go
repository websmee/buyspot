package domain

import (
	"context"
)

type Adviser struct {
}

func NewAdviser() *Adviser {
	return &Adviser{}
}

func (r *Adviser) GetAdviceBySymbol(ctx context.Context, symbol string) (*Advice, error) {
	return &Advice{
		PriceForecast: 3,
		ForecastHours: 8,
		BuyOrderSettings: BuyOrderSettings{
			Amount:            100,
			TakeProfit:        3,
			TakeProfitOptions: []float64{1, 2, 3},
			StopLoss:          -2,
			StopLossOptions:   []float64{-1, -2, -3},
		},
	}, nil
}
