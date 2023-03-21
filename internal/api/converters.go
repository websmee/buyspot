package api

import (
	"fmt"

	"websmee/buyspot/internal/domain"
)

func ConvertSpotToMessage(spot *domain.Spot) *Spot {
	chartsData := buildChartsData(spot.HistoryMarketData, spot.ForecastMarketData)

	var news []NewsArticle
	for i := range spot.News {
		news = append(news, NewsArticle{
			Sentiment: NewsArticleSentiment(spot.News[i].Sentiment),
			Title:     spot.News[i].Title,
			Content:   spot.News[i].Content,
			Created:   spot.News[i].Created,
			Views:     spot.News[i].Views,
		})
	}

	var takeProfitOptions []Option
	for i := range spot.BuyOrderSettings.TakeProfitOptions {
		takeProfitOptions = append(takeProfitOptions, Option{
			Value: spot.BuyOrderSettings.TakeProfitOptions[i],
			Text:  fmt.Sprintf("%.2f", spot.BuyOrderSettings.TakeProfitOptions[i]),
		})
	}

	var stopLossOptions []Option
	for i := range spot.BuyOrderSettings.StopLossOptions {
		stopLossOptions = append(stopLossOptions, Option{
			Value: spot.BuyOrderSettings.StopLossOptions[i],
			Text:  fmt.Sprintf("%.2f", spot.BuyOrderSettings.StopLossOptions[i]),
		})
	}

	return &Spot{
		Asset: Asset{
			Name:        spot.Asset.Name,
			Ticker:      spot.Asset.Ticker,
			Description: spot.Asset.Description,
		},
		ActiveOrders:  spot.ActiveOrders,
		PriceForecast: spot.PriceForecast,
		ChartsData:    chartsData,
		News:          news,
		BuyOrderSettings: BuyOrderSettings{
			Amount:            spot.BuyOrderSettings.Amount,
			TakeProfit:        spot.BuyOrderSettings.TakeProfit,
			TakeProfitOptions: takeProfitOptions,
			StopLoss:          spot.BuyOrderSettings.StopLoss,
			StopLossOptions:   stopLossOptions,
		},
	}
}

func buildChartsData(history []domain.Candlestick, forecast []domain.Candlestick) ChartsData {
	var chartsData ChartsData
	for i := range history {
		chartsData.Times = append(chartsData.Times, history[i].Timestamp.Format("15:04"))
		chartsData.Prices = append(chartsData.Prices, history[i].High)
		chartsData.Volumes = append(chartsData.Volumes, history[i].Volume)
	}
	for i := range forecast {
		chartsData.Times = append(chartsData.Times, forecast[i].Timestamp.Format("15:04"))
		chartsData.Forecast = append(chartsData.Forecast, forecast[i].High)
	}

	return chartsData
}
