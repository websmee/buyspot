package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"websmee/buyspot/internal/domain"
)

type OrderRepository struct {
	client *mongo.Client
}

func NewOrderRepository(client *mongo.Client) *OrderRepository {
	return &OrderRepository{client}
}

func (r *OrderRepository) getCollection() *mongo.Collection {
	return r.client.Database("buyspot").Collection("orders")
}

func (r *OrderRepository) GetUserActiveOrders(ctx context.Context, userID string) ([]domain.Order, error) {
	cur, err := r.getCollection().Find(ctx, primitive.M{"user_id": userID, "status": domain.OrderStatusActive})
	if err != nil {
		return nil, fmt.Errorf("could not get active orders from mongo, err: %w", err)
	}

	var orders []domain.Order
	for cur.Next(ctx) {
		var order domain.Order
		if err := cur.Decode(&order); err != nil {
			return nil, fmt.Errorf("could not decode order from mongo, err: %w", err)
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) GetUserOrderByID(ctx context.Context, userID, orderID string) (*domain.Order, error) {
	id, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return nil, domain.ErrInvalidArgument
	}

	var order domain.Order
	if err := r.getCollection().FindOne(ctx, primitive.M{"_id": id, "user_id": userID}).Decode(&order); err != nil {
		return nil, fmt.Errorf("could not get order from mongo, err: %w", err)
	}

	return &order, nil
}

func (r *OrderRepository) GetUserActiveOrdersCountBySymbol(ctx context.Context, userID, symbol string) (int, error) {
	count, err := r.getCollection().CountDocuments(
		ctx,
		primitive.M{"user_id": userID, "status": domain.OrderStatusActive, "to_symbol": symbol},
	)
	if err != nil {
		return 0, fmt.Errorf("could not get active orders count from mongo, err: %w", err)
	}

	return int(count), nil
}

func (r *OrderRepository) GetActiveOrdersToSell(
	ctx context.Context,
	fromSymbol string,
	toSymbol string,
	toSymbolCurrentPrice float64,
) ([]domain.Order, error) {
	cur, err := r.getCollection().Find(ctx, primitive.M{
		"status":      domain.OrderStatusActive,
		"from_symbol": fromSymbol,
		"to_symbol":   toSymbol,
		"$where": fmt.Sprintf(
			"function() {"+
				"const pnl = %f / this.to_symbol_price * 100 - 100; "+
				"return this.take_profit <= pnl || this.stop_loss >= pnl;"+
				"}",
			toSymbolCurrentPrice,
		),
	})
	if err != nil {
		return nil, fmt.Errorf("could not get active orders to sell from mongo, err: %w", err)
	}

	var orders []domain.Order
	for cur.Next(ctx) {
		var order domain.Order
		if err := cur.Decode(&order); err != nil {
			return nil, fmt.Errorf("could not decode order from mongo, err: %w", err)
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) SaveOrder(ctx context.Context, order *domain.Order) error {
	if order.ID.IsZero() {
		order.ID = primitive.NewObjectID()

		if _, err := r.getCollection().InsertOne(ctx, order); err != nil {
			return fmt.Errorf("could not save new order to mongo, err: %w", err)
		}

		return nil
	}

	if _, err := r.getCollection().UpdateByID(ctx, order.ID, bson.M{"$set": order}); err != nil {
		return fmt.Errorf("could not save order to mongo, err: %w", err)
	}

	return nil
}
