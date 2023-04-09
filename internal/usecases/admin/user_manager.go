package admin

import (
	"context"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

type UserManager struct {
	secretKey      string
	userRepository usecases.UserRepository
}

func NewUserManager(secretKey string, userRepository usecases.UserRepository) *UserManager {
	return &UserManager{secretKey, userRepository}
}

func (r *UserManager) Save(ctx context.Context, secretKey string, user *domain.User) error {
	if secretKey != r.secretKey {
		return domain.ErrForbidden
	}

	var err error
	user.Password, err = domain.GetPasswordHash(user.Password)
	if err != nil {
		return domain.ErrInvalidArgument
	}

	return r.userRepository.CreateOrUpdate(ctx, user)
}
