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
		GetMonth(ctx context.Context, symbol, quote string, before time.Time, interval domain.Interval) ([]domain.Kline, error)
		CreateOrUpdate(
			ctx context.Context,
			symbol string,
			quote string,
			interval domain.Interval,
			kline *domain.Kline,
		) error
	}

	NewsRepository interface {
		GetNewsBySymbol(ctx context.Context, symbol string, from, to time.Time) ([]domain.NewsArticle, error)
		CreateOrUpdate(
			ctx context.Context,
			article *domain.NewsArticle,
		) error
	}

	AssetRepository interface {
		GetAvailableAssets(ctx context.Context) ([]domain.Asset, error)
		GetAssetBySymbol(ctx context.Context, symbol string) (*domain.Asset, error)
	}

	Adviser interface {
		GetAdvice(ctx context.Context, marketData []domain.Kline) (*domain.Advice, error)
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
		GetPrice(ctx context.Context, symbol, quote string) (float64, error)
		GetPrices(ctx context.Context, symbols []string, quote string) (*domain.Prices, error)
		UpdatePrice(ctx context.Context, price float64, symbol, quote string, expiration time.Duration) error
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
			quote string,
			interval domain.Interval,
			handler func(kline *domain.Kline),
			errorHandler func(err error),
		) (done chan struct{}, err error)
	}

	MarketDataService interface {
		GetMonth(ctx context.Context, symbol, quote string, interval domain.Interval) ([]domain.Kline, error)
	}

	NewsService interface {
		GetNews(ctx context.Context, symbols []string, period domain.NewsPeriod) ([]domain.NewsArticle, error)
	}
)
