package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"websmee/buyspot/internal/domain"
)

type CurrentSpotsRepository struct {
	client *redis.Client
}

func NewCurrentSpotsRepository(client *redis.Client) *CurrentSpotsRepository {
	return &CurrentSpotsRepository{client}
}

func (r *CurrentSpotsRepository) GetSpotByIndex(ctx context.Context, index int) (*domain.Spot, error) {
	if cmd := r.client.Get(ctx, getRedisKeySpotByIndex(index)); cmd != nil {
		data, err := cmd.Bytes()
		if err != nil {
			if err == redis.Nil {
				return nil, domain.ErrSpotNotFound
			}

			return nil, fmt.Errorf("could not find spot in redis, err: %w", err)
		}

		var spot domain.Spot
		if err := json.Unmarshal(data, &spot); err != nil {
			return nil, fmt.Errorf("could not decode spot data, err: %w", err)
		}

		return &spot, nil
	}

	return nil, domain.ErrSpotNotFound
}

func (r *CurrentSpotsRepository) GetSpotsCount(ctx context.Context) (int, error) {
	if cmd := r.client.Keys(ctx, "current-spots:*"); cmd != nil {
		keys, err := cmd.Result()
		if err != nil {
			if err == redis.Nil {
				return 0, nil
			}

			return 0, fmt.Errorf("could not get current spots count from redis, err: %w", err)
		}

		return len(keys), nil
	}

	return 0, nil
}

func (r *CurrentSpotsRepository) SaveSpots(ctx context.Context, spots []domain.Spot, expiration time.Duration) error {
	for i := range spots {
		if err := r.saveSpot(ctx, &spots[i], i+1, expiration); err != nil {
			return err
		}
	}

	return nil
}

func (r *CurrentSpotsRepository) saveSpot(
	ctx context.Context,
	spot *domain.Spot,
	index int,
	expiration time.Duration,
) error {
	data, err := json.Marshal(spot)
	if err != nil {
		return fmt.Errorf("could not encode spot struct, err: %w", err)
	}

	if cmd := r.client.Set(ctx, getRedisKeySpotByIndex(index), data, expiration); cmd != nil {
		if _, err := cmd.Result(); err != nil {
			return fmt.Errorf("could not save spot to redis, err: %w", err)
		}
	}

	return nil
}

func getRedisKeySpotByIndex(index int) string {
	return fmt.Sprintf("current-spots:%d", index)
}
