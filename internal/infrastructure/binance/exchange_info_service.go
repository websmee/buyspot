package binance

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/adshao/go-binance/v2"

	"websmee/buyspot/internal/domain"
)

type ExchangeInfoService struct {
	client *binance.Client
}

func NewExchangeInfoService(client *binance.Client) *ExchangeInfoService {
	return &ExchangeInfoService{client}
}

func (s *ExchangeInfoService) GetExchangeInfo(ctx context.Context, symbols []string) ([]domain.AssetExchangeInfo, error) {
	symbolsWithQuotes := make([]string, len(symbols))
	for i := range symbols {
		symbolsWithQuotes[i] = symbols[i] + "USDT"
	}

	exchangeInfo, err := s.client.NewExchangeInfoService().Symbols(symbolsWithQuotes...).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get exchange info from binance, err: %w", err)
	}

	var result []domain.AssetExchangeInfo
	for i := range exchangeInfo.Symbols {
		data, _ := json.Marshal(exchangeInfo.Symbols[i])
		result = append(result, domain.AssetExchangeInfo{
			Symbol:                exchangeInfo.Symbols[i].BaseAsset,
			IsAvailableForTrading: isAvailable(exchangeInfo, exchangeInfo.Symbols[i].BaseAsset),
			Data:                  string(data),
		})
	}

	return result, nil
}
