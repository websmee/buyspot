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
	"websmee/buyspot/internal/domain/indicator"
	binanceInfra "websmee/buyspot/internal/infrastructure/binance"
	"websmee/buyspot/internal/infrastructure/cryptonews"
	"websmee/buyspot/internal/infrastructure/example"
	mongoInfra "websmee/buyspot/internal/infrastructure/mongo"
	"websmee/buyspot/internal/infrastructure/openai"
	redisInfra "websmee/buyspot/internal/infrastructure/redis"
	"websmee/buyspot/internal/infrastructure/simplepush"
	"websmee/buyspot/internal/usecases"
	"websmee/buyspot/internal/usecases/admin"
	"websmee/buyspot/internal/usecases/background"
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
var openaiAPIKey = os.Getenv("BUYSPOT_OPENAI_API_KEY")
var openaiOrgID = os.Getenv("BUYSPOT_OPENAI_ORG_ID")

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
	spotRepository := mongoInfra.NewSpotRepository(mongoClient)
	marketDataRepository := mongoInfra.NewMarketDataRepository(mongoClient)
	newsRepository := mongoInfra.NewNewsRepository(mongoClient)
	assetRepository := mongoInfra.NewAssetRepository(mongoClient)
	adviser := domain.NewAdviser(
		24,
		8,
		4,
		indicator.NewRSI(10, 65),
		indicator.NewVolumeRise(3),
	)
	orderRepository := mongoInfra.NewOrderRepository(mongoClient)
	balanceService := example.NewBalanceService()
	currentSpotsRepository := redisInfra.NewCurrentSpotsRepository(redisClient)
	currentPricesRepository := redisInfra.NewCurrentPricesRepository(redisClient)
	converterService := example.NewConverterService(currentPricesRepository)
	marketDataService := binanceInfra.NewMarketDataService(binance.NewClient(binanceAPIKey, binanceSecretKey))
	newsService := cryptonews.NewNewsService(cryptonewsAPIToken)
	summarizer := openai.NewSummarizer(openai.NewClient(openaiAPIKey, openaiOrgID))
	spotReader := usecases.NewSpotReader(currentSpotsRepository, orderRepository, marketDataRepository)
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
	userManager := admin.NewUserManager(secretKey, userRepository)
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
		spotRepository,
		marketDataRepository,
		newsRepository,
		assetRepository,
		adviser,
		userRepository,
		simplepush.NewNewSpotsNotifier(simplepush.NewClient()),
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
		summarizer,
		newLogger("[NEWS UPDATER]"),
	)
	if err := newsBackgroundUpdater.Run(ctx); err != nil {
		logger.Fatalln(fmt.Errorf("could not run news updater, err: %w", err))
	}

	// web server
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	auth := httpAPI.NewAuth(secretKey, userRepository)
	router.Use(httpAPI.CORSMiddleware())
	router.Use(httpAPI.AuthMiddleware(auth))
	httpAPI.AddBalanceHandlers(router)
	httpAPI.AddSpotHandlers(router, spotReader, spotBuyer)
	httpAPI.AddOrderHandlers(router, orderReader, orderSeller)
	httpAPI.AddPricesHandlers(router, pricesReader)
	httpAPI.AddAdminHandlers(router, marketDataUpdater, newsUpdater, userManager)
	httpAPI.AddAuthHandlers(router, auth)
	_ = router.Run(webAddr)
}

func newLogger(name string) *log.Logger {
	return log.New(os.Stdout, name+" ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
