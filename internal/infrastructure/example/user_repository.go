package example

import (
	"context"

	"websmee/buyspot/internal/domain"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	return &domain.User{ID: "test"}, nil
}
