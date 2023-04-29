package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"websmee/buyspot/internal/domain"
)

type NewsRepository struct {
	client *mongo.Client
}

func NewNewsRepository(client *mongo.Client) *NewsRepository {
	return &NewsRepository{client}
}

func (r *NewsRepository) getCollection() *mongo.Collection {
	return r.client.Database("buyspot_main").Collection("news")
}

func (r *NewsRepository) GetNewsBySymbol(ctx context.Context, symbol string, from, to time.Time) ([]domain.NewsArticle, error) {
	cur, err := r.getCollection().Find(ctx, bson.M{
		"$and": []bson.M{
			{"symbols": symbol},
			{"created": bson.M{"$gte": from}},
			{"created": bson.M{"$lte": to}},
		},
	}, options.Find().SetSort(bson.D{{"created", -1}}))
	if err != nil {
		return nil, fmt.Errorf("could not get %s news from mongo, err: %w", symbol, err)
	}

	var news []domain.NewsArticle
	for cur.Next(ctx) {
		var n domain.NewsArticle
		if err := cur.Decode(&n); err != nil {
			return nil, fmt.Errorf("could not decode %s news data, err: %w", symbol, err)
		}
		news = append(news, n)
	}

	return news, nil
}

func (r *NewsRepository) IsArticleExists(ctx context.Context, article *domain.NewsArticle) (bool, error) {
	count, err := r.getCollection().CountDocuments(ctx, bson.M{"url": article.URL})
	if err != nil {
		return false, fmt.Errorf("could not check if news article exists in mongo, err: %w", err)
	}

	return count > 0, nil
}

func (r *NewsRepository) CreateOrUpdate(
	ctx context.Context,
	article *domain.NewsArticle,
) error {
	res, err := r.getCollection().UpdateOne(
		ctx,
		bson.M{"url": article.URL},
		bson.M{"$set": article},
	)
	if err != nil {
		return fmt.Errorf("could not update news article in mongo, err: %w", err)
	}

	if res.MatchedCount == 0 {
		_, err := r.getCollection().InsertOne(ctx, article)
		if err != nil {
			return fmt.Errorf("could not insert news article in mongo, err: %w", err)
		}
	}

	return nil
}
