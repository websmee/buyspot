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
		GetMonth(ctx context.Context, symbol string, interval domain.Interval) ([]domain.Kline, error)
		CreateOrUpdate(
			ctx context.Context,
			symbol string,
			interval domain.Interval,
			kline *domain.Kline,
		) error
	}

	NewsRepository interface {
		GetFreshNewsBySymbol(ctx context.Context, symbol string, from time.Time) ([]domain.NewsArticle, error)
	}

	AssetRepository interface {
		GetAvailableAssets(ctx context.Context) ([]domain.Asset, error)
		GetAssetByTicket(ctx context.Context, symbol string) (*domain.Asset, error)
	}

	Adviser interface {
		GetAdviceBySymbol(ctx context.Context, symbol string) (*domain.Advice, error)
	}

	OrderRepository interface {
		GetUserOrderByID(ctx context.Context, userID, orderID string) (*domain.Order, error)
		GetUserActiveOrders(ctx context.Context, userID string) ([]domain.Order, error)
		GetUserActiveOrdersCountBySymbol(ctx context.Context, userID string, symbol string) (int, error)
		GetActiveOrdersToSell(
			ctx context.Context,
			fromSymbol string,
			toSymbol string,
			toSymbolCurrentPrice float64,
		) ([]domain.Order, error)
		SaveOrder(ctx context.Context, order *domain.Order) error
	}

	FundsConverter interface {
		ConvertUserFunds(
			ctx context.Context,
			user *domain.User,
			fromSymbol string,
			toSymbol string,
			amount float64,
		) (float64, error)
	}

	CurrentPricesRepository interface {
		GetCurrentPrices(ctx context.Context, inSymbol string) (*domain.Prices, error)
		SaveCurrentPrices(ctx context.Context, prices *domain.Prices, inSymbol string) error
	}

	BalanceService interface {
		GetUserActiveBalance(ctx context.Context, user *domain.User) (*domain.Balance, error)
		GetUserBalances(ctx context.Context, user *domain.User) ([]domain.Balance, error)
		GetAvailableSymbols(ctx context.Context) ([]string, error)
	}

	PricesService interface {
		GetCurrentPrices(ctx context.Context, inSymbol string) (*domain.Prices, error)
	}

	ConverterService interface {
		Convert(ctx context.Context, user *domain.User, amount float64, fromSymbol, toSymbol string) (float64, error)
	}

	UserRepository interface {
		GetUserByID(ctx context.Context, userID string) (*domain.User, error)
	}

	MarketDataStream interface {
		Subscribe(
			ctx context.Context,
			symbol string,
			interval domain.Interval,
			handler func(kline *domain.Kline),
			errorHandler func(err error),
		) error
	}
)
