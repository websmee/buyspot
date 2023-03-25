package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	httpAPI "websmee/buyspot/internal/api/http"
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

	marketDataRepository := example.NewMarketDataRepository()
	newsRepository := example.NewNewsRepository()
	assetRepository := example.NewAssetRepository()
	adviser := example.NewAdviser()
	orderRepository := example.NewOrderRepository()
	balanceService := example.NewBalanceService()
	pricesService := example.NewPricesService()
	currentSpotsRepository := redisInfra.NewCurrentSpotsRepository(redisClient)
	currentPricesRepository := redisInfra.NewCurrentPricesRepository(redisClient)
	spotReader := usecases.NewSpotReader(currentSpotsRepository)
	orderReader := usecases.NewOrderReader(orderRepository)
	pricesReader := usecases.NewPricesReader(currentPricesRepository, balanceService)

	// background processed
	spotMaker := background.NewSpotMaker(
		currentSpotsRepository,
		marketDataRepository,
		newsRepository,
		assetRepository,
		adviser,
		orderRepository,
		newLogger("[SPOT MAKER]"),
	)
	if err := spotMaker.Run(ctx); err != nil {
		logger.Fatalln(fmt.Errorf("could not run spot maker, err: %w", err))
	}

	currentPricesUpdater := background.NewCurrentPricesUpdater(
		currentPricesRepository,
		balanceService,
		pricesService,
		newLogger("[CURRENT PRICES UPDATER]"),
	)
	if err := currentPricesUpdater.Run(ctx); err != nil {
		logger.Fatalln(fmt.Errorf("could not run current prices updater, err: %w", err))
	}

	// web server
	router := gin.Default()
	router.Use(httpAPI.CORSMiddleware())
	router.Use(httpAPI.AuthMiddleware())
	httpAPI.AddBalanceHandlers(router)
	httpAPI.AddSpotHandlers(router, spotReader)
	httpAPI.AddOrderHandlers(router, orderReader)
	httpAPI.AddPricesHandlers(router, pricesReader)
	_ = router.Run("localhost:8080")
}

func newLogger(name string) *log.Logger {
	return log.New(os.Stdout, name+" ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
