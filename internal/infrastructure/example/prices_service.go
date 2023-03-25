package example

import (
	"context"

	"websmee/buyspot/internal/domain"
)

type PricesService struct {
}

func NewPricesService() *PricesService {
	return &PricesService{}
}

func (r *PricesService) GetCurrentPrices(_ context.Context, _ string) (*domain.Prices, error) {
	return &domain.Prices{
		InTicker: "USDT",
		PricesByTickers: map[string]float64{
			"USDT": 1,
			"BTC":  27506.60,
			"ETH":  1749.46,
			"SHIB": 0.0000106,
		},
	}, nil
}
