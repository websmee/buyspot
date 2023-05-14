package main

import (
	"context"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	httpAPI "websmee/buyspot/internal/api/http"
	adminHTTPAPI "websmee/buyspot/internal/api/http/admin"
	binanceInfra "websmee/buyspot/internal/infrastructure/binance"
	"websmee/buyspot/internal/infrastructure/cryptonews"
	"websmee/buyspot/internal/infrastructure/local"
	mongoInfra "websmee/buyspot/internal/infrastructure/mongo"
	redisInfra "websmee/buyspot/internal/infrastructure/redis"
	"websmee/buyspot/internal/usecases"
	"websmee/buyspot/internal/usecases/admin"
)

var secretKey = os.Getenv("BUYSPOT_SECRET_KEY")
var binanceAPIKey = os.Getenv("BUYSPOT_BINANCE_API_KEY")
var binanceSecretKey = os.Getenv("BUYSPOT_BINANCE_SECRET_KEY")
var redisAddr = os.Getenv("BUYSPOT_REDIS")
var redisPassword = os.Getenv("BUYSPOT_REDIS_PASSWORD")
var mongoURI = os.Getenv("BUYSPOT_MONGO")
var mongoUser = os.Getenv("BUYSPOT_MONGO_USER")
var mongoPwd = os.Getenv("BUYSPOT_MONGO_PWD")
var webAddr = os.Getenv("BUYSPOT_WEB_ADDR")
var cryptonewsAPIToken = os.Getenv("BUYSPOT_CRYPTONEWS_API_TOKEN")

func main() {
	ctx := context.Background()

	// dependencies
	logger := newLogger("[MAIN]")

	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr, Password: redisPassword})
	defer func() {
		if err := redisClient.Close(); err != nil {
			logger.Fatalln(err)
		}
	}()

	mongoClient, err := mongoInfra.Connect(ctx, mongoURI, mongoUser, mongoPwd)
	if err != nil {
		logger.Fatalln(err)
	}
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			logger.Fatalln(err)
		}
	}()

	userRepository := mongoInfra.NewUserRepository(mongoClient)
	marketDataRepository := mongoInfra.NewMarketDataRepository(mongoClient)
	newsRepository := mongoInfra.NewNewsRepository(mongoClient)
	assetRepository := mongoInfra.NewAssetRepository(mongoClient)
	orderRepository := mongoInfra.NewOrderRepository(mongoClient)
	spotRepository := mongoInfra.NewSpotRepository(mongoClient)
	balanceService := mongoInfra.NewBalanceService(mongoClient)
	currentSpotsRepository := redisInfra.NewCurrentSpotsRepository(redisClient)
	currentPricesRepository := redisInfra.NewCurrentPricesRepository(redisClient)
	tradingService := local.NewTradingService(currentPricesRepository, balanceService)
	marketDataService := binanceInfra.NewMarketDataService(binance.NewClient(binanceAPIKey, binanceSecretKey))
	newsService := cryptonews.NewNewsService(cryptonewsAPIToken)
	spotReader := usecases.NewSpotReader(currentSpotsRepository, orderRepository, marketDataRepository)
	spotBuyer := usecases.NewSpotBuyer(
		orderRepository,
		spotRepository,
		tradingService,
		balanceService,
		assetRepository,
	)
	orderReader := usecases.NewOrderReader(orderRepository)
	balanceReader := usecases.NewBalanceReader(balanceService)
	orderSeller := usecases.NewOrderSeller(
		orderRepository,
		tradingService,
		balanceService,
	)
	pricesReader := usecases.NewPricesReader(assetRepository, currentPricesRepository, balanceService)
	userManager := admin.NewUserManager(secretKey, userRepository, balanceService)
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

	// web server
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	auth := httpAPI.NewAuth(secretKey, userRepository)
	router.Use(httpAPI.CORSMiddleware())
	router.Use(httpAPI.AuthMiddleware(auth))
	httpAPI.AddBalanceHandlers(router, balanceReader)
	httpAPI.AddSpotHandlers(router, spotReader, spotBuyer)
	httpAPI.AddOrderHandlers(router, orderReader, orderSeller)
	httpAPI.AddPricesHandlers(router, pricesReader)
	adminHTTPAPI.AddAdminHandlers(router, marketDataUpdater, newsUpdater, userManager)
	httpAPI.AddAuthHandlers(router, auth)
	_ = router.Run(webAddr)
}

func newLogger(name string) *log.Logger {
	return log.New(os.Stdout, name+" ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
