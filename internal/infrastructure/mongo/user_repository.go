package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"websmee/buyspot/internal/domain"
)

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{client}
}

func (r *UserRepository) getCollection() *mongo.Collection {
	return r.client.Database("buyspot_main").Collection("users")
}

func (r *UserRepository) CreateOrUpdate(ctx context.Context, user *domain.User) error {
	res, err := r.getCollection().UpdateOne(
		ctx,
		bson.M{"email": user.Email},
		bson.M{"$set": bson.M{
			"email":              user.Email,
			"password":           user.Password,
			"binance_api_key":    user.BinanceAPIKey,
			"binance_secret_key": user.BinanceSecretKey,
		}},
	)
	if err != nil {
		return fmt.Errorf("could not update user %s in mongo, err: %w", user.Email, err)
	}

	if res.MatchedCount == 0 {
		user.ID = primitive.NewObjectID()
		_, err := r.getCollection().InsertOne(ctx, user)
		if err != nil {
			return fmt.Errorf("could not insert user %s in mongo, err: %w", user.Email, err)
		}
	}

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, userID string) (*domain.User, error) {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, domain.ErrInvalidArgument
	}

	var user domain.User
	if err := r.getCollection().FindOne(ctx, primitive.M{"_id": id}).Decode(&user); err != nil {
		return nil, fmt.Errorf("could not get user by id = '%s' from mongo, err: %w", userID, err)
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := r.getCollection().FindOne(ctx, primitive.M{"email": email}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, fmt.Errorf("could not find user '%s' in mongo, err: %w", email, err)
	}

	return &user, nil
}
