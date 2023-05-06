package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"websmee/buyspot/internal/domain"
)

type SpotRepository struct {
	client *mongo.Client
}

func NewSpotRepository(client *mongo.Client) *SpotRepository {
	return &SpotRepository{client}
}

func (r *SpotRepository) getCollection() *mongo.Collection {
	return r.client.Database("buyspot_main").Collection("spots")
}

func (r *SpotRepository) SaveSpots(ctx context.Context, spots []domain.Spot) error {
	var docs []interface{}
	for _, spot := range spots {
		docs = append(docs, spot)
	}

	if _, err := r.getCollection().InsertMany(ctx, docs); err != nil {
		return err
	}

	return nil
}
