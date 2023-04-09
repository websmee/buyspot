package example

import (
	"context"

	"websmee/buyspot/internal/usecases"
)

type ConverterService struct {
	currentPricesRepository usecases.CurrentPricesRepository
}

func NewConverterService(currentPricesRepository usecases.CurrentPricesRepository) *ConverterService {
	return &ConverterService{currentPricesRepository}
}

func (s *ConverterService) Convert(ctx context.Context, userID string, amount float64, fromSymbol, toSymbol string) (float64, error) {
	price, _ := s.currentPricesRepository.GetPrice(ctx, toSymbol, fromSymbol)
	if price == 0 {
		price, _ = s.currentPricesRepository.GetPrice(ctx, fromSymbol, toSymbol)
		return amount * price, nil
	}

	return amount / price, nil
}
