package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"websmee/buyspot/internal/domain"
)

type MarketDataRepository struct {
	client *mongo.Client
}

func NewMarketDataRepository(client *mongo.Client) *MarketDataRepository {
	return &MarketDataRepository{client}
}

func (r *MarketDataRepository) getCollection(
	symbol string,
	quote string,
	interval domain.Interval,
) *mongo.Collection {
	return r.client.
		Database("buyspot_market_data").
		Collection(fmt.Sprintf("%s%s_%s", symbol, quote, interval))
}

func (r *MarketDataRepository) GetMonth(
	ctx context.Context,
	symbol string,
	quote string,
	before time.Time,
	interval domain.Interval,
) ([]domain.Kline, error) {
	cur, err := r.getCollection(symbol, quote, interval).Find(ctx, bson.M{
		"$and": []bson.M{
			{"start_time": bson.M{"$gte": before.AddDate(0, -1, 0)}},
			{"start_time": bson.M{"$lt": before}},
		},
	})
	if err != nil {
		return nil, fmt.Errorf(
			"could not get %s%s klines from mongo, err: %w",
			symbol,
			quote,
			err,
		)
	}

	var klines []domain.Kline
	for cur.Next(ctx) {
		var kline domain.Kline
		if err := cur.Decode(&kline); err != nil {
			return nil, fmt.Errorf("could not decode %s%s kline data, err: %w", symbol, quote, err)
		}
		klines = append(klines, kline)
	}

	return klines, nil
}

func (r *MarketDataRepository) GetPeriod(
	ctx context.Context,
	symbol string,
	quote string,
	interval domain.Interval,
) ([]domain.Kline, error) {
	cur, err := r.getCollection(symbol, quote, interval).Find(ctx, bson.M{
		"start_time": bson.M{
			"$gte": time.Now().AddDate(0, -1, 0),
		},
	})
	if err != nil {
		return nil, fmt.Errorf(
			"could not get %s%s klines from mongo, err: %w",
			symbol,
			quote,
			err,
		)
	}

	var klines []domain.Kline
	for cur.Next(ctx) {
		var kline domain.Kline
		if err := cur.Decode(&kline); err != nil {
			return nil, fmt.Errorf("could not decode %s%s kline data, err: %w", symbol, quote, err)
		}
		klines = append(klines, kline)
	}

	return klines, nil
}

func (r *MarketDataRepository) CreateOrUpdate(
	ctx context.Context,
	symbol string,
	quote string,
	interval domain.Interval,
	kline *domain.Kline,
) error {
	res, err := r.getCollection(symbol, quote, interval).UpdateOne(
		ctx,
		bson.M{"start_time": kline.StartTime},
		bson.M{"$set": kline},
	)
	if err != nil {
		return fmt.Errorf("could not update %s%s kline data, err: %w", symbol, quote, err)
	}

	if res.MatchedCount == 0 {
		_, err := r.getCollection(symbol, quote, interval).InsertOne(ctx, kline)
		if err != nil {
			return fmt.Errorf("could not insert %s%s kline data, err: %w", symbol, quote, err)
		}
	}

	return nil
}
