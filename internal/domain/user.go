package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const CtxKeyUserID = "user"

type User struct {
	ID               primitive.ObjectID `bson:"_id"`
	Email            string             `bson:"email"`
	Password         string             `bson:"password"`
	BinanceAPIKey    string             `bson:"binance_api_key"`
	BinanceSecretKey string             `bson:"binance_secret_key"`
}

func GetCtxUserID(ctx context.Context) string {
	return ctx.Value(CtxKeyUserID).(string)
}

func GetPasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
