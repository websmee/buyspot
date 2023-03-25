package example

import "context"

type ConverterService struct {
}

func NewConverterService() *ConverterService {
	return &ConverterService{}
}

func (ConverterService) Convert(ctx context.Context, amount float64, fromTicker, toTicker string) (float64, error) {
	return 100, nil
}
