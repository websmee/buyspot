package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	httpAPI "websmee/buyspot/internal/api/http"
	"websmee/buyspot/internal/infrastructure/binance"
	"websmee/buyspot/internal/infrastructure/example"
	mongoInfra "websmee/buyspot/internal/infrastructure/mongo"
	redisInfra "websmee/buyspot/internal/infrastructure/redis"
	"websmee/buyspot/internal/usecases"
	"websmee/buyspot/internal/usecases/background"
)

func main() {
	ctx := context.Background()

	// dependencies
	logger := newLogger("[MAIN]")

	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer func() {
		if err := redisClient.Close(); err != nil {
			logger.Fatalln(err)
		}
	}()

	mongoClient, err := mongoInfra.Connect(ctx, "mongodb://localhost:27017")
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
	newsRepository := example.NewNewsRepository()
	assetRepository := example.NewAssetRepository()
	adviser := example.NewAdviser()
	orderRepository := mongoInfra.NewOrderRepository(mongoClient)
	balanceService := example.NewBalanceService()
	currentSpotsRepository := redisInfra.NewCurrentSpotsRepository(redisClient)
	currentPricesRepository := redisInfra.NewCurrentPricesRepository(redisClient)
	converterService := example.NewConverterService(currentPricesRepository)
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

	marketDataUpdater := background.NewMarketDataUpdater(
		balanceService,
		assetRepository,
		binance.NewMarketDataStream(),
		marketDataRepository,
		currentPricesRepository,
		newLogger("[MARKET DATA UPDATER]"),
	)
	if err := marketDataUpdater.Run(ctx); err != nil {
		logger.Fatalln(fmt.Errorf("could not run market data updater, err: %w", err))
	}
	defer marketDataUpdater.Close()

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

	// web server
	router := gin.Default()
	router.Use(httpAPI.CORSMiddleware())
	router.Use(httpAPI.AuthMiddleware())
	httpAPI.AddBalanceHandlers(router)
	httpAPI.AddSpotHandlers(router, spotReader, spotBuyer)
	httpAPI.AddOrderHandlers(router, orderReader, orderSeller)
	httpAPI.AddPricesHandlers(router, pricesReader)
	_ = router.Run("localhost:8080")
}

func newLogger(name string) *log.Logger {
	return log.New(os.Stdout, name+" ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
