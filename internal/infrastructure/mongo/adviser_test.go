package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/infrastructure/example"
	redisInfra "websmee/buyspot/internal/infrastructure/redis"
	"websmee/buyspot/internal/usecases/background"
)

func TestAdviser(t *testing.T) {
	riseInARow := 3
	checksCount := 0
	adviceCount := 0
	takeProfitCount := 0
	stopLossCount := 0
	unknownCount := 0
	var advices []domain.Advice
	var times []time.Time
	var asts []domain.Asset
	c, _ := Connect(context.Background(), "mongodb://localhost:27017")
	ar := NewAssetRepository(c)
	adviser := domain.NewAdviser(riseInARow, 3, 0, 2.5, 999)
	assets, _ := ar.GetAvailableAssets(context.Background())
	for _, asset := range assets {
		marketData := getTestMarketData(asset.Symbol, "USDT", domain.IntervalHour)
		for i := range marketData {
			if i > riseInARow {
				checksCount++
				a, _ := adviser.GetAdvice(context.Background(), marketData[:i+1])
				//fmt.Println(a)
				if a != nil {
					advices = append(advices, *a)
					times = append(times, marketData[i].EndTime)
					asts = append(asts, asset)
					adviceCount++
					if i < len(marketData)-riseInARow {
						check := adviser.CheckAdvice(a, marketData[i+1:])
						switch check {
						case 1:
							takeProfitCount++
						case -1:
							stopLossCount++
						default:
							unknownCount++
						}
					} else {
						unknownCount++
					}
				}
			}
		}
	}
	fmt.Println(checksCount)
	fmt.Println(adviceCount)
	fmt.Println(takeProfitCount)
	fmt.Println(stopLossCount)
	fmt.Println(unknownCount)

	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	marketDataRepository := NewMarketDataRepository(c)
	newsRepository := NewNewsRepository(c)
	assetRepository := NewAssetRepository(c)
	balanceService := example.NewBalanceService()
	currentSpotsRepository := redisInfra.NewCurrentSpotsRepository(redisClient)
	spotMaker := background.NewSpotMaker(
		balanceService,
		currentSpotsRepository,
		marketDataRepository,
		newsRepository,
		assetRepository,
		adviser,
		log.New(os.Stdout, "[ADVISER TEST] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
	)

	var spots []domain.Spot
	for i := range advices {
		spot := spotMaker.GetSpot(context.Background(), &advices[i], times[i], &asts[i], []string{"USDT"})
		spots = append(spots, *spot)
	}

	currentSpotsRepository.ClearSpots(context.Background())
	currentSpotsRepository.SaveSpots(context.Background(), spots, time.Hour)
}

func getTestMarketData(
	symbol string,
	quote string,
	interval domain.Interval,
) []domain.Kline {
	ctx := context.Background()
	client, _ := Connect(ctx, "mongodb://localhost:27017")
	cur, _ := client.
		Database("buyspot_market_data").
		Collection(fmt.Sprintf("%s%s_%s", symbol, quote, interval)).
		Find(ctx, bson.M{
			"$and": []bson.M{
				{"start_time": bson.M{"$gte": time.Now().AddDate(0, -1, 0)}},
				{"end_time": bson.M{"$lte": time.Now()}},
			},
		})

	var klines []domain.Kline
	for cur.Next(ctx) {
		var kline domain.Kline
		cur.Decode(&kline)
		klines = append(klines, kline)
	}

	return klines
}
