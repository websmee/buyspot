package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	httpAPI "websmee/buyspot/internal/api/http"
	"websmee/buyspot/internal/domain"
	binanceInfra "websmee/buyspot/internal/infrastructure/binance"
	"websmee/buyspot/internal/infrastructure/cryptonews"
	"websmee/buyspot/internal/infrastructure/example"
	mongoInfra "websmee/buyspot/internal/infrastructure/mongo"
	redisInfra "websmee/buyspot/internal/infrastructure/redis"
	"websmee/buyspot/internal/usecases"
	"websmee/buyspot/internal/usecases/admin"
	"websmee/buyspot/internal/usecases/background"
)

var secretKey = os.Getenv("BUYSPOT_SECRET_KEY")
var binanceAPIKey = os.Getenv("BUYSPOT_BINANCE_API_KEY")
var binanceSecretKey = os.Getenv("BUYSPOT_BINANCE_SECRET_KEY")
var authUsers = os.Getenv("BUYSPOT_AUTH_USERS")
var redisAddr = os.Getenv("BUYSPOT_REDIS")
var mongoURI = os.Getenv("BUYSPOT_MONGO")
var webAddr = os.Getenv("BUYSPOT_WEB_ADDR")
var cryptonewsAPIToken = os.Getenv("BUYSPOT_CRYPTONEWS_API_TOKEN")

func main() {
	ctx := context.Background()

	// dependencies
	logger := newLogger("[MAIN]")

	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	defer func() {
		if err := redisClient.Close(); err != nil {
			logger.Fatalln(err)
		}
	}()

	mongoClient, err := mongoInfra.Connect(ctx, mongoURI)
	if err != nil {
		logger.Fatalln(err)
	}
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			logger.Fatalln(err)
		}
	}()

	userRepository := example.NewUserRepository()
	marketDataRepository := mongoInfra.NewMarketDataRepository(mongoClient)
	newsRepository := mongoInfra.NewNewsRepository(mongoClient)
	assetRepository := mongoInfra.NewAssetRepository(mongoClient)
	adviser := domain.NewAdviser(3, 2.5)
	orderRepository := mongoInfra.NewOrderRepository(mongoClient)
	balanceService := example.NewBalanceService()
	currentSpotsRepository := redisInfra.NewCurrentSpotsRepository(redisClient)
	currentPricesRepository := redisInfra.NewCurrentPricesRepository(redisClient)
	converterService := example.NewConverterService(currentPricesRepository)
	marketDataService := binanceInfra.NewMarketDataService(binance.NewClient(binanceAPIKey, binanceSecretKey))
	newsService := cryptonews.NewNewsService(cryptonewsAPIToken)
	spotReader := usecases.NewSpotReader(currentSpotsRepository, orderRepository)
	spotBuyer := usecases.NewSpotBuyer(
		orderRepository,
		converterService,
		balanceService,
		assetRepository,
	)
	orderReader := usecases.NewOrderReader(orderRepository)
	orderSeller := usecases.NewOrderSeller(
		orderRepository,
		converterService,
		balanceService,
	)
	pricesReader := usecases.NewPricesReader(assetRepository, currentPricesRepository, balanceService)
	marketDataUpdater := admin.NewMarketDataUpdater(
		secretKey,
		balanceService,
		assetRepository,
		marketDataRepository,
		marketDataService,
	)
	newsUpdater := admin.NewNewsUpdater(
		secretKey,
		assetRepository,
		newsRepository,
		newsService,
	)

	// background processed
	spotMaker := background.NewSpotMaker(
		balanceService,
		currentSpotsRepository,
		marketDataRepository,
		newsRepository,
		assetRepository,
		adviser,
		newLogger("[SPOT MAKER]"),
	)
	if err := spotMaker.Run(ctx); err != nil {
		logger.Fatalln(fmt.Errorf("could not run spot maker, err: %w", err))
	}

	marketDataBackgroundUpdater := background.NewMarketDataUpdater(
		balanceService,
		assetRepository,
		binanceInfra.NewMarketDataStream(),
		marketDataRepository,
		currentPricesRepository,
		newLogger("[MARKET DATA UPDATER]"),
	)
	if err := marketDataBackgroundUpdater.Run(ctx); err != nil {
		logger.Fatalln(fmt.Errorf("could not run market data updater, err: %w", err))
	}
	defer marketDataBackgroundUpdater.Close()

	orderBackgroundSeller := background.NewOrderSeller(
		userRepository,
		balanceService,
		assetRepository,
		currentPricesRepository,
		orderRepository,
		converterService,
		newLogger("[ORDER SELLER]"),
	)
	if err := orderBackgroundSeller.Run(ctx); err != nil {
		logger.Fatalln(fmt.Errorf("could not run order seller, err: %w", err))
	}

	newsBackgroundUpdater := background.NewNewsUpdater(
		assetRepository,
		newsRepository,
		newsService,
		newLogger("[NEWS UPDATER]"),
	)
	if err := newsBackgroundUpdater.Run(ctx); err != nil {
		logger.Fatalln(fmt.Errorf("could not run news updater, err: %w", err))
	}

	// web server
	router := gin.Default()
	auth := httpAPI.NewSimpleAuth(authUsers)
	router.Use(httpAPI.CORSMiddleware())
	router.Use(httpAPI.AuthMiddleware(auth))
	httpAPI.AddBalanceHandlers(router)
	httpAPI.AddSpotHandlers(router, spotReader, spotBuyer)
	httpAPI.AddOrderHandlers(router, orderReader, orderSeller)
	httpAPI.AddPricesHandlers(router, pricesReader)
	httpAPI.AddAdminHandlers(router, marketDataUpdater, newsUpdater)
	httpAPI.AddAuthHandlers(router, auth)
	_ = router.Run(webAddr)
}

func newLogger(name string) *log.Logger {
	return log.New(os.Stdout, name+" ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
