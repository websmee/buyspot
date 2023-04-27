package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/domain/indicator"
	"websmee/buyspot/internal/infrastructure/example"
	redisInfra "websmee/buyspot/internal/infrastructure/redis"
	"websmee/buyspot/internal/usecases/background"
)

func TestAdviser(t *testing.T) {
	checksCount := 0
	adviceCount := 0
	takeProfitCount := 0
	stopLossCount := 0
	unknownCount := 0
	percentProfit := 0.0
	var advices []domain.Advice
	var times []time.Time
	var asts []domain.Asset
	c, _ := Connect(context.Background(), "mongodb://localhost:27017")
	ar := NewAssetRepository(c)
	adviser := domain.NewAdviser(
		24,
		8,
		5,
		indicator.NewRSI(10, 70),
		indicator.NewVolumeRise(3),
		//indicator.NewMFI(14, 75),
		//indicator.NewPVO(3, 4, 2),
		indicator.NewTrue(),
	)
	assets, _ := ar.GetAvailableAssets(context.Background())
	for _, asset := range assets {
		marketData := getTestMarketData(asset.Symbol, "USDT", domain.IntervalHour)
		for i := range marketData {
			if i > 24 && i < len(marketData)-24 {
				checksCount++
				a, _ := adviser.GetAdvice(context.Background(), marketData[:i+1])
				//fmt.Println(a)
				if a != nil {
					adviceCount++
					check := adviser.CheckAdvice(a, marketData[i+1:])
					switch check {
					case 1:
						takeProfitCount++
						percentProfit += a.PriceForecast
						a.IsProfitable = true
					case -1:
						stopLossCount++
						percentProfit -= a.PriceForecast
					default:
						unknownCount++
					}
					advices = append(advices, *a)
					times = append(times, marketData[i].EndTime.UTC())
					asts = append(asts, asset)
				}
			}
		}
	}
	fmt.Println("TOTAL CHECKS:", checksCount)
	fmt.Println("FREQUENCY:", float64(adviceCount)/30, " per day")
	fmt.Println("PROFIT RATE:", float64(takeProfitCount)/float64(adviceCount)*100, "%")
	fmt.Println("PROFIT AMOUNT:", percentProfit, "%")

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

	sort.Slice(spots, func(i, j int) bool {
		return spots[i].ForecastMarketDataByQuotes["USDT"][0].StartTime.Before(spots[j].ForecastMarketDataByQuotes["USDT"][0].StartTime)
	})

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
		}, options.Find().SetSort(bson.D{{"start_time", 1}}))

	var klines []domain.Kline
	for cur.Next(ctx) {
		var kline domain.Kline
		cur.Decode(&kline)
		klines = append(klines, kline)
	}

	return klines
}

//func TestDeepLearning(t *testing.T) {
//	var data = getNNTrainingData("BTC", "USDT", domain.IntervalHour)
//
//	n := deep.NewNeural(&deep.Config{
//		/* Input dimensionality */
//		Inputs: 48,
//		/* Two hidden layers consisting of two neurons each, and a single output */
//		Layout: []int{24, 8},
//		/* Activation functions: Sigmoid, Tanh, ReLU, Linear */
//		Activation: deep.ActivationReLU,
//		/* Determines output layer activation & loss function:
//		ModeRegression: linear outputs with MSE loss
//		ModeMultiClass: softmax output with Cross Entropy loss
//		ModeMultiLabel: sigmoid output with Cross Entropy loss
//		ModeBinary: sigmoid output with binary CE loss */
//		Mode: deep.ModeMultiLabel,
//		/* Weight initializers: {deep.NewNormal(μ, σ), deep.NewUniform(μ, σ)} */
//		Weight: deep.NewNormal(0.6, 0.1),
//		/* Apply bias */
//		Bias: true,
//	})
//
//	// params: learning rate, momentum, alpha decay, nesterov
//	optimizer := training.NewSGD(0.05, 0.1, 1e-6, true)
//	// params: optimizer, verbosity (print stats at every 50th iteration)
//	trainer := training.NewTrainer(optimizer, 10)
//
//	training, heldout := data.Split(0.5)
//	trainer.Train(n, training, heldout, 100) // training, validation, iterations
//
//	fmt.Println(data[78].Input, data[78].Response, "=>", n.Predict(data[78].Input))
//	fmt.Println(data[123].Input, data[123].Response, "=>", n.Predict(data[123].Input))
//	fmt.Println(data[100].Input, data[100].Response, "=>", n.Predict(data[100].Input))
//	fmt.Println(data[234].Input, data[234].Response, "=>", n.Predict(data[234].Input))
//	fmt.Println(data[300].Input, data[300].Response, "=>", n.Predict(data[300].Input))
//}
//
//func getNNTrainingData(
//	symbol string,
//	quote string,
//	interval domain.Interval,
//) training.Examples {
//	ctx := context.Background()
//	client, _ := Connect(ctx, "mongodb://localhost:27017")
//	cur, _ := client.
//		Database("buyspot_market_data").
//		Collection(fmt.Sprintf("%s%s_%s", symbol, quote, interval)).
//		Find(ctx, bson.M{
//			"$and": []bson.M{
//				{"start_time": bson.M{"$gte": time.Now().AddDate(0, -2, 0)}},
//				{"end_time": bson.M{"$lte": time.Now()}},
//			},
//		})
//
//	var klines []domain.Kline
//	for cur.Next(ctx) {
//		var kline domain.Kline
//		cur.Decode(&kline)
//		klines = append(klines, kline)
//	}
//
//	var data training.Examples
//	for i := range klines {
//		if i < 24 {
//			continue
//		}
//		if i > len(klines)-9 {
//			break
//		}
//
//		var inputs []float64
//		for j := i - 24; j < i; j++ {
//			inputs = append(inputs, klines[j].Close)
//			inputs = append(inputs, klines[j].Volume)
//		}
//
//		//diff := (klines[i-8].Close - klines[i].Close) / klines[i-8].Close
//		//classes := make([]float64, 3)
//		//if diff > 0.02 {
//		//	classes[0] = 1
//		//} else if diff < -0.02 {
//		//	classes[1] = 1
//		//} else {
//		//	classes[2] = 1
//		//}
//
//		var outputs []float64
//		for j := i; j < i+8; j++ {
//			outputs = append(outputs, klines[j].Close)
//		}
//
//		data = append(data, training.Example{Input: inputs, Response: outputs})
//	}
//
//	return data
//}
