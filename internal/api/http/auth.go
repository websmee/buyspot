package http

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"

	"websmee/buyspot/internal/domain"
)

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInvalidToken = errors.New("invalid token")

type UserFinder interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}

type Auth struct {
	secretKey  string
	userFinder UserFinder
}

func NewAuth(secretKey string, userFinder UserFinder) *Auth {
	return &Auth{
		secretKey,
		userFinder,
	}
}

func (s *Auth) CheckCredentials(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := s.userFinder.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrInvalidCredentials
	}

	if !domain.CheckPasswordHash(password, user.Password) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (s *Auth) GetUserIDByToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrInvalidToken
	}

	id, ok := claims["id"].(string)
	if !ok {
		return "", ErrInvalidToken
	}

	if id == "" {
		return "", ErrInvalidToken
	}

	return id, nil
}

func (s *Auth) GetToken(user *domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claims["id"] = user.ID.Hex()

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
