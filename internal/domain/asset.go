package domain

type Asset struct {
	Symbol      string `bson:"symbol"`
	Order       int    `bson:"order"`
	IsAvailable bool   `bson:"is_available"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
}
