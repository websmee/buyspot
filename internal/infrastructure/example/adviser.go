package example

import (
	"context"

	"websmee/buyspot/internal/domain"
)

type Adviser struct {
}

func NewAdviser() *Adviser {
	return &Adviser{}
}

func (r *Adviser) GetAdviceBySymbol(ctx context.Context, symbol, base string) (*domain.Advice, error) {
	return &domain.Advice{
		PriceForecast: 3,
		ForecastHours: 8,
		BuyOrderSettings: domain.BuyOrderSettings{
			Amount:            100,
			TakeProfit:        3,
			TakeProfitOptions: []float64{1, 2, 3},
			StopLoss:          -2,
			StopLossOptions:   []float64{-1, -2, -3},
		},
	}, nil
}
