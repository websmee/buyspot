package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"websmee/buyspot/internal/domain"
)

type DemoBalanceService struct {
	client *mongo.Client
}

func NewDemoBalanceService(client *mongo.Client) *DemoBalanceService {
	return &DemoBalanceService{client}
}

func (s *DemoBalanceService) getCollection() *mongo.Collection {
	return s.client.Database("buyspot_main").Collection("balances")
}

func (s *DemoBalanceService) CreateBalance(ctx context.Context, balance *domain.Balance) error {
	if _, err := s.getCollection().InsertOne(ctx, balance); err != nil {
		return fmt.Errorf("could not create balance in mongo, err: %w", err)
	}

	return nil
}

func (s *DemoBalanceService) ChangeBalance(ctx context.Context, userID, symbol string, value float64) error {
	if _, err := s.getCollection().UpdateOne(ctx, primitive.M{
		"user_id":   userID,
		"symbol":    symbol,
		"is_active": true,
	}, primitive.M{
		"$inc": primitive.M{
			"amount": value,
		},
	}); err != nil {
		return fmt.Errorf("could not change balance by user id = '%s' in mongo, err: %w", userID, err)
	}

	return nil
}

func (s *DemoBalanceService) GetUserActiveBalance(ctx context.Context, user *domain.User) (*domain.Balance, error) {
	var balance domain.Balance
	if err := s.getCollection().FindOne(ctx, primitive.M{
		"user_id":   user.ID.Hex(),
		"is_active": true,
	}).Decode(&balance); err != nil {
		return nil, fmt.Errorf("could not get balance by user id = '%s' from mongo, err: %w", user.ID.Hex(), err)
	}

	return &balance, nil
}

func (s *DemoBalanceService) GetUserBalances(ctx context.Context, user *domain.User) ([]domain.Balance, error) {
	cur, err := s.getCollection().Find(ctx, primitive.M{
		"user_id": user.ID.Hex(),
	})
	if err != nil {
		return nil, fmt.Errorf("could not get balances by user id = '%s' from mongo, err: %w", user.ID.Hex(), err)
	}

	var balances []domain.Balance
	for cur.Next(ctx) {
		var balance domain.Balance
		if err := cur.Decode(&balance); err != nil {
			return nil, fmt.Errorf("could not decode balance data, err: %w", err)
		}
		balances = append(balances, balance)
	}

	return balances, nil
}

func (s *DemoBalanceService) GetAvailableSymbols(_ context.Context) ([]string, error) {
	return []string{"USDT"}, nil
}
