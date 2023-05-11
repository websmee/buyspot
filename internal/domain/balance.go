package domain

type Balance struct {
	UserID   string  `bson:"user_id"`
	Symbol   string  `bson:"symbol"`
	Amount   float64 `bson:"amount"`
	IsActive bool    `bson:"is_active"`
}
