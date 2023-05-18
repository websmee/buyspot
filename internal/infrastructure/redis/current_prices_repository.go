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

func (r *CurrentPricesRepository) GetPrices(ctx context.Context, symbols []string, quote string) (*domain.Prices, error) {
	if len(symbols) == 0 {
		return nil, nil
	}

	var prices domain.Prices
	prices.Quote = quote
	prices.PricesBySymbols = make(map[string]float64)

	for _, symbol := range symbols {
		price, err := r.GetPrice(ctx, symbol, quote)
		if err != nil {
			return nil, err
		}

		prices.PricesBySymbols[symbol] = price
	}

	return &prices, nil
}

func (r *CurrentPricesRepository) GetPrice(ctx context.Context, symbol, quote string) (float64, error) {
	if symbol == quote {
		return 1, nil
	}

	if cmd := r.client.Get(ctx, getRedisKeyPrice(symbol, quote)); cmd != nil {
		price, err := cmd.Float64()
		if err != nil {
			if err == redis.Nil {
				return 0, domain.ErrPriceNotFound
			}

			return 0, fmt.Errorf("could not find %s%s price in redis, err: %w", symbol, quote, err)
		}

		return price, nil
	}

	return 0, nil
}

func (r *CurrentPricesRepository) UpdatePrice(
	ctx context.Context,
	price float64,
	symbol string,
	quote string,
	expiration time.Duration,
) error {
	if err := r.client.Set(ctx, getRedisKeyPrice(symbol, quote), price, expiration).Err(); err != nil {
		return fmt.Errorf("could not update %s%s price in redis, err: %w", symbol, quote, err)
	}

	return nil
}

func getRedisKeyPrice(symbol, quote string) string {
	return fmt.Sprintf("current-prices:%s:%s", quote, symbol)
}
