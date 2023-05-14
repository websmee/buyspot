package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *SpotRepository) SaveSpot(ctx context.Context, spot *domain.Spot) (string, error) {
	res, err := r.getCollection().InsertOne(ctx, spot)
	if err != nil {
		return "", fmt.Errorf("could not save spot to mongo, err: %w", err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *SpotRepository) GetSpotByID(ctx context.Context, id string) (*domain.Spot, error) {
	var spot domain.Spot

	err := r.getCollection().FindOne(ctx, map[string]string{"_id": id}).Decode(&spot)
	if err != nil {
		return nil, fmt.Errorf("could not get spot by ID = '%s', err: %w", id, err)
	}

	return &spot, nil
}
