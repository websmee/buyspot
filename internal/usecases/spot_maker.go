package usecases

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"websmee/buyspot/internal/domain"
)

type CurrentSpotsWriter interface {
	SaveSpots(ctx context.Context, spots []domain.Spot, expiration time.Duration) error
}

type MarketDataReader interface {
	GetAllHours(ctx context.Context, assetTicker string) ([]domain.Candlestick, error)
}

type MarketDataForecastBuilder interface {
	BuildForecastHours(
		ctx context.Context,
		assetTicker string,
		currentPrice float64,
		priceForecast float64,
		hours int,
	) ([]domain.Candlestick, error)
}

type SpotMaker struct {
	currentSpotsWriter CurrentSpotsWriter
	marketDataReader   MarketDataReader
	logger             *log.Logger
}

func NewSpotMaker(
	spotWriter CurrentSpotsWriter,
	marketDataReader MarketDataReader,
	logger *log.Logger,
) *SpotMaker {
	return &SpotMaker{
		spotWriter,
		marketDataReader,
		logger,
	}
}

func (m *SpotMaker) Run(ctx context.Context) error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(time.Minute).Do(func() {
		spots, err := m.makeSpots(ctx)
		if err != nil {
			m.logger.Println(fmt.Errorf("could not get new spots, err: %w", err))
			return
		}

		err = m.currentSpotsWriter.SaveSpots(ctx, spots, time.Minute)
		if err != nil {
			m.logger.Println(fmt.Errorf("could not save new spots, err: %w", err))
			return
		}
	})

	s.StartAsync()

	return err
}

func (m *SpotMaker) makeSpots(ctx context.Context) ([]domain.Spot, error) {
	asset := domain.Asset{
		Ticker: "BTC",
		Name:   "Bitcoin",
		Description: "Bitcoin (abbreviation: BTC[a] or XBT[b]; sign: ₿) " +
			"is a protocol which implements a highly available, public, permanent, and decentralized ledger. " +
			"In order to add to the ledger, a user must prove they control an entry in the ledger. " +
			"The protocol specifies that the entry indicates an amount of a token, bitcoin with a minuscule b. " +
			"The user can update the ledger, assigning some of their bitcoin to another entry in the ledger. " +
			"Because the token has characteristics of money, it can be thought of as a digital currency.",
	}
	priceForecast := 3.0
	historyMarketData, _ := m.marketDataReader.GetAllHours(ctx, asset.Ticker)
	forecastMarketData, _ := buildForecastHours(
		ctx,
		asset.Ticker,
		historyMarketData[len(historyMarketData)-1].High,
		priceForecast,
		8,
	)
	news := []domain.NewsArticle{
		{
			Sentiment: domain.NewsArticleSentimentNeutral,
			Title:     "Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain",
			Content:   "Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain. Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain. Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain. Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain",
			Created:   time.Now().Add(-1 * time.Hour),
			Views:     15678,
		},
		{
			Sentiment: domain.NewsArticleSentimentPositive,
			Title:     "IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed",
			Content:   "IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed. IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed. IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed. IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed",
			Created:   time.Now().Add(-12 * time.Hour),
			Views:     25678,
		},
		{
			Sentiment: domain.NewsArticleSentimentNegative,
			Title:     "Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown",
			Content:   "Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown. Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown. Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown. Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown",
			Created:   time.Now().Add(-48 * time.Hour),
			Views:     178,
		},
	}
	buyOrderSettings := domain.BuyOrderSettings{
		Amount:            100,
		TakeProfit:        3,
		TakeProfitOptions: []float64{1, 2, 3},
		StopLoss:          -2,
		StopLossOptions:   []float64{-1, -2, -3},
	}

	return []domain.Spot{
		{
			ID:                 primitive.NewObjectID(),
			Asset:              asset,
			ActiveOrders:       3,
			PriceForecast:      priceForecast,
			HistoryMarketData:  historyMarketData,
			ForecastMarketData: forecastMarketData,
			News:               news,
			BuyOrderSettings:   buyOrderSettings,
		},
		{
			ID:                 primitive.NewObjectID(),
			Asset:              asset,
			ActiveOrders:       3,
			PriceForecast:      priceForecast,
			HistoryMarketData:  historyMarketData,
			ForecastMarketData: forecastMarketData,
			News:               news,
			BuyOrderSettings:   buyOrderSettings,
		},
		{
			ID:                 primitive.NewObjectID(),
			Asset:              asset,
			ActiveOrders:       3,
			PriceForecast:      priceForecast,
			HistoryMarketData:  historyMarketData,
			ForecastMarketData: forecastMarketData,
			News:               news,
			BuyOrderSettings:   buyOrderSettings,
		},
	}, nil
}

func buildForecastHours(
	_ context.Context,
	assetTicker string,
	currentPrice float64,
	priceForecast float64,
	hours int,
) ([]domain.Candlestick, error) {
	price := currentPrice
	endPrice := currentPrice + (currentPrice * priceForecast / 100)
	diff := (endPrice - currentPrice) / float64(hours)
	curvature := diff * 0.9
	var candlesticks []domain.Candlestick
	for i := 1; i <= hours; i++ {
		candlesticks = append(
			candlesticks,
			getForecastCandlestick(assetTicker, price, 0, time.Now().Add(time.Duration(i)*time.Hour)),
		)
		p := float64(i) / float64(hours)
		price += diff - curvature + (2 * curvature * p)
	}

	return candlesticks, nil
}

func getForecastCandlestick(assetTicker string, price float64, volume int64, t time.Time) domain.Candlestick {
	return domain.Candlestick{
		Open:        price,
		Low:         price,
		High:        price,
		Close:       price,
		AdjClose:    price,
		Volume:      volume,
		Timestamp:   t,
		Interval:    "1h",
		AssetTicker: assetTicker,
	}
}
