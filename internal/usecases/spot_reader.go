package usecases

import (
	"context"

	"websmee/buyspot/internal/domain"
)

type CurrentSpotsReader interface {
	GetSpotByIndex(ctx context.Context, index int) (*domain.Spot, error)
	GetSpotsCount(ctx context.Context) (int, error)
}

type SpotReader struct {
	currentSpotsReader CurrentSpotsReader
}

func NewSpotReader(currentSpotsReader CurrentSpotsReader) *SpotReader {
	return &SpotReader{currentSpotsReader}
}

func (r *SpotReader) GetSpotByIndex(ctx context.Context, index int) (*domain.Spot, error) {
	return r.currentSpotsReader.GetSpotByIndex(ctx, index)
}

func (r *SpotReader) GetSpotsCount(ctx context.Context) (int, error) {
	return r.currentSpotsReader.GetSpotsCount(ctx)
}
