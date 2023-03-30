package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"websmee/buyspot/internal/domain"
)

type CurrentPricesRepository struct {
	client *redis.Client
}

func NewCurrentPricesRepository(client *redis.Client) *CurrentPricesRepository {
	return &CurrentPricesRepository{client}
}

func (r *CurrentPricesRepository) GetPrices(ctx context.Context, symbols []string, base string) (*domain.Prices, error) {
	if len(symbols) == 0 {
		return nil, nil
	}

	var prices domain.Prices
	prices.Base = base
	prices.PricesBySymbols = make(map[string]float64)

	for _, symbol := range symbols {
		price, err := r.GetPrice(ctx, symbol, base)
		if err != nil {
			return nil, err
		}

		prices.PricesBySymbols[symbol] = price
	}

	return &prices, nil
}

func (r *CurrentPricesRepository) GetPrice(ctx context.Context, symbol, base string) (float64, error) {
	if symbol == base {
		return 1, nil
	}

	if cmd := r.client.Get(ctx, getRedisKeyPrice(symbol, base)); cmd != nil {
		price, err := cmd.Float64()
		if err != nil {
			if err == redis.Nil {
				return 0, nil
			}

			return 0, fmt.Errorf("could not find %s%s price in redis, err: %w", symbol, base, err)
		}

		return price, nil
	}

	return 0, nil
}

func (r *CurrentPricesRepository) UpdatePrice(
	ctx context.Context,
	price float64,
	symbol string,
	base string,
	expiration time.Duration,
) error {
	if err := r.client.Set(ctx, getRedisKeyPrice(symbol, base), price, expiration).Err(); err != nil {
		return fmt.Errorf("could not update %s%s price in redis, err: %w", symbol, base, err)
	}

	return nil
}

func getRedisKeyPrice(symbol, base string) string {
	return fmt.Sprintf("current-prices:%s:%s", base, symbol)
}
