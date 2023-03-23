package usecases

import (
	"context"
	"time"

	"websmee/buyspot/internal/domain"
)

type (
	CurrentSpotsRepository interface {
		SaveSpots(ctx context.Context, spots []domain.Spot, expiration time.Duration) error
	}

	MarketDataRepository interface {
		GetAllHours(ctx context.Context, assetTicker string) ([]domain.Candlestick, error)
	}

	NewsRepository interface {
		GetFreshNewsByTicker(ctx context.Context, ticker string, from time.Time) ([]domain.NewsArticle, error)
	}

	AssetRepository interface {
		GetAvailableAssets(ctx context.Context) ([]domain.Asset, error)
	}

	Adviser interface {
		GetAdviceByTicker(ctx context.Context, ticker string) (*domain.Advice, error)
	}

	OrderRepository interface {
		GetUserOrders(ctx context.Context, userID string) ([]domain.Order, error)
		GetActiveOrdersCountByTicker(ctx context.Context, ticker string) (int, error)
		SaveOrder(ctx context.Context, order *domain.Order) error
	}

	FundsConverter interface {
		ConvertUserFunds(
			ctx context.Context,
			user *domain.User,
			fromTicker string,
			toTicker string,
			amount float64,
		) (float64, error)
	}

	PricesRepository interface {
		GetCurrentPrices(ctx context.Context, inTicker string) (*domain.Prices, error)
		SaveCurrentPrices(ctx context.Context, prices *domain.Prices) error
	}

	BalanceService interface {
		GetUserBalance(ctx context.Context, user *domain.User) (*domain.Balance, error)
	}
)
