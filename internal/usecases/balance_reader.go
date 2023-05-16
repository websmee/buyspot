package usecases

import (
	"context"
	"fmt"

	"websmee/buyspot/internal/domain"
)

type BalanceReader struct {
	userRepository     UserRepository
	balanceService     BalanceService
	demoBalanceService BalanceService
}

func NewBalanceReader(
	userRepository UserRepository,
	balanceService BalanceService,
	demoBalanceService BalanceService,
) *BalanceReader {
	return &BalanceReader{
		userRepository,
		balanceService,
		demoBalanceService,
	}
}

func (r *BalanceReader) GetActiveBalance(ctx context.Context) (*domain.Balance, error) {
	userID := domain.GetCtxUserID(ctx)
	if userID == "" {
		return nil, domain.ErrUnauthorized
	}

	user, err := r.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not get user by ID = '%s', err: %w", userID, err)
	}

	var balance *domain.Balance
	if user.IsDemo {
		balance, err = r.demoBalanceService.GetUserActiveBalance(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("could not get demo user active balance, err: %w", err)
		}
	} else {
		balance, err = r.balanceService.GetUserActiveBalance(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("could not get user active balance, err: %w", err)
		}
	}

	return balance, nil
}
