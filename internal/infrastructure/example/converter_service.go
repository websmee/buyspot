package example

import (
	"context"

	"websmee/buyspot/internal/domain"
)

type ConverterService struct {
}

func NewConverterService() *ConverterService {
	return &ConverterService{}
}

func (ConverterService) Convert(ctx context.Context, user *domain.User, amount float64, fromTicker, toTicker string) (float64, error) {
	prices, _ := NewPricesService().GetCurrentPrices(ctx, fromTicker)
	price := prices.PricesByTickers[toTicker]

	return amount / price, nil
}
