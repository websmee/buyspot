package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"websmee/buyspot/internal/domain"
)

type AdviserRepository struct {
	client *mongo.Client
}

func NewAdviserRepository(client *mongo.Client) *AdviserRepository {
	return &AdviserRepository{
		client: client,
	}
}

func (r *AdviserRepository) getCollection() *mongo.Collection {
	return r.client.Database("buyspot_main").Collection("advisers")
}

func (r *AdviserRepository) SaveAdvisers(ctx context.Context, advisers []domain.Adviser) error {
	newAdvisers := make([]interface{}, len(advisers))
	for i, adviser := range advisers {
		newAdvisers[i] = adviser
	}

	_, err := r.getCollection().InsertMany(ctx, newAdvisers)
	if err != nil {
		return fmt.Errorf("could not save advisers to mongo, err: %w", err)
	}

	return nil
}

func (r *AdviserRepository) GetLatestAdvisers(ctx context.Context) ([]domain.Adviser, error) {
	cur, err := r.getCollection().Find(ctx, bson.M{"is_latest": true})
	if err != nil {
		return nil, fmt.Errorf("could not get latest advisers from mongo, err: %w", err)
	}

	var advisers []domain.Adviser
	for cur.Next(ctx) {
		var adviser domain.Adviser
		if err := cur.Decode(&adviser); err != nil {
			return nil, fmt.Errorf("could not decode adviser data, err: %w", err)
		}
		advisers = append(advisers, adviser)
	}

	return advisers, nil
}

func (r *AdviserRepository) MarkAllAdvisersAsNotLatest(ctx context.Context) error {
	_, err := r.getCollection().UpdateMany(
		ctx,
		bson.M{},
		bson.M{"$set": bson.M{"is_latest": false}},
	)
	if err != nil {
		return fmt.Errorf("could not mark all advisers as not latest in mongo, err: %w", err)
	}

	return nil
}
