package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/adshao/go-binance/v2"
	"github.com/redis/go-redis/v9"

	binanceInfra "websmee/buyspot/internal/infrastructure/binance"
	"websmee/buyspot/internal/infrastructure/cryptonews"
	mongoInfra "websmee/buyspot/internal/infrastructure/mongo"
	"websmee/buyspot/internal/infrastructure/openai"
	redisInfra "websmee/buyspot/internal/infrastructure/redis"
	"websmee/buyspot/internal/infrastructure/telegram"
	"websmee/buyspot/internal/usecases/background"
)

var redisAddr = os.Getenv("BUYSPOT_REDIS")
var binanceAPIKey = os.Getenv("BUYSPOT_BINANCE_API_KEY")
var binanceSecretKey = os.Getenv("BUYSPOT_BINANCE_SECRET_KEY")
var redisPassword = os.Getenv("BUYSPOT_REDIS_PASSWORD")
var mongoURI = os.Getenv("BUYSPOT_MONGO")
var mongoUser = os.Getenv("BUYSPOT_MONGO_USER")
var mongoPwd = os.Getenv("BUYSPOT_MONGO_PWD")
var cryptonewsAPIToken = os.Getenv("BUYSPOT_CRYPTONEWS_API_TOKEN")
var openaiAPIKey = os.Getenv("BUYSPOT_OPENAI_API_KEY")
var openaiOrgID = os.Getenv("BUYSPOT_OPENAI_ORG_ID")
var telegramBotAPIToken = os.Getenv("BUYSPOT_TELEGRAM_BOT_API_TOKEN")

func main() {
	ctx, cancel := context.WithCancel(context.Background())

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
	adviserRepository := mongoInfra.NewAdviserRepository(mongoClient)
	marketDataRepository := mongoInfra.NewMarketDataRepository(mongoClient)
	newsRepository := mongoInfra.NewNewsRepository(mongoClient)
	assetRepository := mongoInfra.NewAssetRepository(mongoClient)
	orderRepository := mongoInfra.NewOrderRepository(mongoClient)
	balanceService := mongoInfra.NewDemoBalanceService(mongoClient)
	exchangeInfoService := binanceInfra.NewExchangeInfoService(binance.NewClient(binanceAPIKey, binanceSecretKey))
	currentSpotsRepository := redisInfra.NewCurrentSpotsRepository(redisClient)
	currentPricesRepository := redisInfra.NewCurrentPricesRepository(redisClient)
	tradingService := binanceInfra.NewTradingService()
	demoTradingService := mongoInfra.NewDemoTradingService(currentPricesRepository, balanceService)
	newsService := cryptonews.NewNewsService(cryptonewsAPIToken)
	summarizer := openai.NewSummarizer(openai.NewClient(openaiAPIKey, openaiOrgID))
	notifier := telegram.NewNotifier(telegramBotAPIToken)
	notificationsSubscriber := telegram.NewNotificationsSubscriber(telegramBotAPIToken)

	spotMaker := background.NewSpotMaker(
		balanceService,
		currentSpotsRepository,
		spotRepository,
		marketDataRepository,
		newsRepository,
		assetRepository,
		adviserRepository,
		notifier,
		newLogger("[SPOT MAKER]"),
	)
	if err := spotMaker.Run(ctx); err != nil {
		logger.Fatalln(fmt.Errorf("could not run spot maker, err: %w", err))
	}

	assetUpdater := background.NewAssetUpdater(
		assetRepository,
		exchangeInfoService,
		newLogger("[ASSET UPDATER]"),
	)
	go func() {
		if err := assetUpdater.Run(ctx); err != nil {
			logger.Fatalln(fmt.Errorf("could not run asset updater, err: %w", err))
		}
	}()

	adviserCreator := background.NewAdviserCreator(
		assetRepository,
		marketDataRepository,
		adviserRepository,
		newLogger("[ADVISER CREATOR]"),
	)
	go func() {
		if err := adviserCreator.Run(ctx); err != nil {
			logger.Fatalln(fmt.Errorf("could not run adviser creator, err: %w", err))
		}
	}()

	marketDataUpdater := background.NewMarketDataUpdater(
		balanceService,
		assetRepository,
		binanceInfra.NewMarketDataStream(),
		marketDataRepository,
		currentPricesRepository,
		newLogger("[MARKET DATA UPDATER]"),
	)
	go func() {
		if err := marketDataUpdater.Run(ctx); err != nil {
			logger.Fatalln(fmt.Errorf("could not run market data updater, err: %w", err))
		}
	}()

	orderSeller := background.NewOrderSeller(
		userRepository,
		balanceService,
		assetRepository,
		currentPricesRepository,
		orderRepository,
		tradingService,
		demoTradingService,
		notifier,
		newLogger("[ORDER SELLER]"),
	)
	if err := orderSeller.Run(ctx); err != nil {
		logger.Fatalln(fmt.Errorf("could not run order seller, err: %w", err))
	}

	newsUpdater := background.NewNewsUpdater(
		assetRepository,
		newsRepository,
		newsService,
		summarizer,
		newLogger("[NEWS UPDATER]"),
	)
	if err := newsUpdater.Run(ctx); err != nil {
		logger.Fatalln(fmt.Errorf("could not run news updater, err: %w", err))
	}

	background.NewNotificationKeyUpdater(
		notificationsSubscriber,
		userRepository,
		newLogger("[NOTIFICATION KEY UPDATER]"),
	).Run(ctx)

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
	cancel()
}

func newLogger(name string) *log.Logger {
	return log.New(os.Stdout, name+" ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
