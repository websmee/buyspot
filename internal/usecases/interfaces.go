package usecases

import (
	"context"
	"time"

	"websmee/buyspot/internal/domain"
)

type (
	CurrentSpotsRepository interface {
		GetSpotByIndex(ctx context.Context, index int) (*domain.Spot, error)
		GetSpotsCount(ctx context.Context) (int, error)
		SaveSpots(ctx context.Context, spots []domain.Spot, expiration time.Duration) error
	}

	SpotRepository interface {
		SaveSpot(ctx context.Context, spot *domain.Spot) error
		GetSpotByID(ctx context.Context, id string) (*domain.Spot, error)
	}

	MarketDataRepository interface {
		GetKlines(
			ctx context.Context,
			symbol string,
			quote string,
			from time.Time,
			to time.Time,
			interval domain.Interval,
		) ([]domain.Kline, error)
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
		IsArticleExists(ctx context.Context, article *domain.NewsArticle) (bool, error)
		CreateOrUpdate(
			ctx context.Context,
			article *domain.NewsArticle,
		) error
	}

	AssetRepository interface {
		GetAvailableAssets(ctx context.Context) ([]domain.Asset, error)
		GetAssetBySymbol(ctx context.Context, symbol string) (*domain.Asset, error)
		CreateOrUpdate(ctx context.Context, asset *domain.Asset) error
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

	BalanceRepository interface {
		CreateBalance(ctx context.Context, balance *domain.Balance) error
	}

	BalanceService interface {
		GetUserActiveBalance(ctx context.Context, user *domain.User) (*domain.Balance, error)
		GetUserBalances(ctx context.Context, user *domain.User) ([]domain.Balance, error)
		GetAvailableSymbols(ctx context.Context) ([]string, error)
	}

	PricesService interface {
		GetCurrentPrices(ctx context.Context, inSymbol string) (*domain.Prices, error)
	}

	TradingService interface {
		Buy(
			ctx context.Context,
			user *domain.User,
			balanceSymbol string,
			balanceAmount float64,
			tradeSymbol string,
		) (string, error)
		Sell(
			ctx context.Context,
			user *domain.User,
			tradeSymbol string,
			tradeAmount string,
			balanceSymbol string,
		) (float64, error)
	}

	UserRepository interface {
		CreateOrUpdate(ctx context.Context, user *domain.User) error
		GetByID(ctx context.Context, userID string) (*domain.User, error)
		GetUsers(ctx context.Context) ([]domain.User, error)
	}

	MarketDataStream interface {
		Subscribe(
			ctx context.Context,
			symbol string,
			quote string,
			interval domain.Interval,
			handler func(kline *domain.Kline),
			errorHandler func(err error),
		) error
	}

	MarketDataService interface {
		GetKlines(
			ctx context.Context,
			symbol string,
			quote string,
			from time.Time,
			to time.Time,
			interval domain.Interval,
		) ([]domain.Kline, error)
	}

	NewsService interface {
		GetNews(ctx context.Context, symbols []string, period domain.NewsPeriod) ([]domain.NewsArticle, error)
	}

	Summarizer interface {
		GetSummary(ctx context.Context, url string) (string, error)
	}

	Notifier interface {
		Notify(ctx context.Context, user *domain.User, title, message string) error
	}

	AdviserRepository interface {
		SaveAdvisers(ctx context.Context, advisers []domain.Adviser) error
		GetLatestAdvisers(ctx context.Context) ([]domain.Adviser, error)
		MarkAllAdvisersAsNotLatest(ctx context.Context) error
	}

	ExchangeInfoService interface {
		GetExchangeInfo(ctx context.Context, symbols []string) ([]domain.AssetExchangeInfo, error)
	}
)
