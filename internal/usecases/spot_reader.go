package usecases

import (
	"context"
	"fmt"

	"websmee/buyspot/internal/domain"
)

type CurrentSpotsReader interface {
	GetSpotByIndex(ctx context.Context, index int) (*domain.Spot, error)
	GetSpotsCount(ctx context.Context) (int, error)
}

type SpotReader struct {
	currentSpotsReader CurrentSpotsReader
	orderRepository    OrderRepository
}

func NewSpotReader(
	currentSpotsReader CurrentSpotsReader,
	orderRepository OrderRepository,
) *SpotReader {
	return &SpotReader{
		currentSpotsReader,
		orderRepository,
	}
}

func (r *SpotReader) GetSpotByIndex(ctx context.Context, index int) (*domain.Spot, error) {
	user := domain.GetCtxUser(ctx)
	if user == nil {
		return nil, domain.ErrUnauthorized
	}

	spot, err := r.currentSpotsReader.GetSpotByIndex(ctx, index)
	if err != nil {
		return nil, fmt.Errorf("could not get spot by index %d, err: %w", index, err)
	}

	activeOrdersCount, err := r.orderRepository.GetUserActiveOrdersCountBySymbol(ctx, user.ID, spot.Asset.Symbol)
	if err != nil {
		return nil, fmt.Errorf(
			"could not get user (ID = '%s') active orders count by symbol %s, err: %w",
			user.ID,
			spot.Asset.Symbol,
			err,
		)
	}

	spot.ActiveOrders = activeOrdersCount

	return spot, nil
}

func (r *SpotReader) GetSpotsCount(ctx context.Context) (int, error) {
	return r.currentSpotsReader.GetSpotsCount(ctx)
}
