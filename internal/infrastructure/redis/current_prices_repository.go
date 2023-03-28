package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"

	"websmee/buyspot/internal/domain"
)

type CurrentPricesRepository struct {
	client *redis.Client
}

func NewCurrentPricesRepository(client *redis.Client) *CurrentPricesRepository {
	return &CurrentPricesRepository{client}
}

func (r *CurrentPricesRepository) GetCurrentPrices(ctx context.Context, inSymbol string) (*domain.Prices, error) {
	if cmd := r.client.Get(ctx, getRedisKeyCurrentPricesInSymbol(inSymbol)); cmd != nil {
		data, err := cmd.Bytes()
		if err != nil {
			if err == redis.Nil {
				return nil, domain.ErrCurrentPricesNotFound
			}

			return nil, fmt.Errorf("could not find current prices in redis, err: %w", err)
		}

		var prices domain.Prices
		if err := json.Unmarshal(data, &prices); err != nil {
			return nil, fmt.Errorf("could not decode prices data, err: %w", err)
		}

		return &prices, nil
	}

	return nil, domain.ErrCurrentPricesNotFound
}

func (r *CurrentPricesRepository) SaveCurrentPrices(ctx context.Context, prices *domain.Prices, inSymbol string) error {
	data, err := json.Marshal(prices)
	if err != nil {
		return fmt.Errorf("could not encode prices struct, err: %w", err)
	}

	if cmd := r.client.Set(ctx, getRedisKeyCurrentPricesInSymbol(inSymbol), data, 0); cmd != nil {
		if _, err := cmd.Result(); err != nil {
			return fmt.Errorf("could not save current prices to redis, err: %w", err)
		}
	}

	return nil
}

func getRedisKeyCurrentPricesInSymbol(symbol string) string {
	return fmt.Sprintf("current-prices:%s", symbol)
}
