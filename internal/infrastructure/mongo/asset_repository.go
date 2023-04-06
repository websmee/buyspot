package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"websmee/buyspot/internal/domain"
)

type AssetRepository struct {
	client *mongo.Client
}

func NewAssetRepository(client *mongo.Client) *AssetRepository {
	return &AssetRepository{client}
}

func (r *AssetRepository) getCollection() *mongo.Collection {
	return r.client.Database("buyspot_main").Collection("assets")
}

func (r *AssetRepository) GetAvailableAssets(ctx context.Context) ([]domain.Asset, error) {
	cur, err := r.getCollection().Find(ctx, bson.M{
		"is_available": true,
	})
	if err != nil {
		return nil, fmt.Errorf("could not get available assets from mongo, err: %w", err)
	}

	var assets []domain.Asset
	for cur.Next(ctx) {
		var asset domain.Asset
		if err := cur.Decode(&asset); err != nil {
			return nil, fmt.Errorf("could not decode asset data, err: %w", err)
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

func (r *AssetRepository) GetAssetBySymbol(ctx context.Context, symbol string) (*domain.Asset, error) {
	var asset domain.Asset
	if err := r.getCollection().FindOne(ctx, bson.M{
		"symbol": symbol,
	}).Decode(&asset); err != nil {
		return nil, fmt.Errorf("could not get %s asset from mongo, err: %w", symbol, err)
	}

	return &asset, nil
}

func (r *AssetRepository) CreateOrUpdate(
	ctx context.Context,
	asset *domain.Asset,
) error {
	res, err := r.getCollection().UpdateOne(
		ctx,
		bson.M{"symbol": asset.Symbol},
		bson.M{"$set": asset},
	)
	if err != nil {
		return fmt.Errorf("could not update %s asset in mongo, err: %w", asset.Symbol, err)
	}

	if res.MatchedCount == 0 {
		_, err := r.getCollection().InsertOne(ctx, asset)
		if err != nil {
			return fmt.Errorf("could not insert %s asset in mongo, err: %w", asset.Symbol, err)
		}
	}

	return nil
}
