package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/redis/go-redis/v9"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/domain/indicator"
	binanceInfra "websmee/buyspot/internal/infrastructure/binance"
	"websmee/buyspot/internal/infrastructure/cryptonews"
	"websmee/buyspot/internal/infrastructure/local"
	mongoInfra "websmee/buyspot/internal/infrastructure/mongo"
	"websmee/buyspot/internal/infrastructure/openai"
	redisInfra "websmee/buyspot/internal/infrastructure/redis"
	"websmee/buyspot/internal/infrastructure/simplepush"
	"websmee/buyspot/internal/usecases/background"
)

var redisAddr = os.Getenv("BUYSPOT_REDIS")
var redisPassword = os.Getenv("BUYSPOT_REDIS_PASSWORD")
var mongoURI = os.Getenv("BUYSPOT_MONGO")
var mongoUser = os.Getenv("BUYSPOT_MONGO_USER")
var mongoPwd = os.Getenv("BUYSPOT_MONGO_PWD")
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
	balanceService := mongoInfra.NewBalanceService(mongoClient)
	currentSpotsRepository := redisInfra.NewCurrentSpotsRepository(redisClient)
	currentPricesRepository := redisInfra.NewCurrentPricesRepository(redisClient)
	tradingService := local.NewTradingService(currentPricesRepository, balanceService)
	newsService := cryptonews.NewNewsService(cryptonewsAPIToken)
	summarizer := openai.NewSummarizer(openai.NewClient(openaiAPIKey, openaiOrgID))

	spotMaker := background.NewSpotMaker(
		balanceService,
		currentSpotsRepository,
		spotRepository,
		marketDataRepository,
		newsRepository,
		assetRepository,
		adviser,
		userRepository,
		simplepush.NewNotifier(),
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
		tradingService,
		simplepush.NewNotifier(),
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

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	<-done
}

func newLogger(name string) *log.Logger {
	return log.New(os.Stdout, name+" ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
