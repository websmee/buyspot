package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Asset struct {
	ID          primitive.ObjectID `json:"_id"`
	Ticker      string             `json:"ticker"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
}
