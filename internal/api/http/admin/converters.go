package admin

import (
	"websmee/buyspot/internal/domain"
)

func UserToDomain(user *User) *domain.User {
	return &domain.User{
		Email:            user.Email,
		Password:         user.Password,
		BinanceAPIKey:    user.BinanceAPIKey,
		BinanceSecretKey: user.BinanceSecretKey,
		NotificationsKey: user.NotificationsKey,
		IsDemo:           user.IsDemo,
	}
}

func BalanceToDomain(balance *Balance) *domain.Balance {
	return &domain.Balance{
		UserID:   balance.UserID,
		Symbol:   balance.Symbol,
		Amount:   balance.Amount,
		IsActive: balance.IsActive,
	}
}
