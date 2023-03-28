package api

import (
	"fmt"

	"websmee/buyspot/internal/domain"
)

func ConvertOrderToMessages(order *domain.Order) *Order {
	return &Order{
		ID:          order.ID.Hex(),
		FromAmount:  order.FromAmount,
		FromSymbol:  order.FromSymbol,
		ToAmount:    order.ToAmount,
		ToSymbol:    order.ToSymbol,
		ToAssetName: order.ToAssetName,
		TakeProfit:  order.TakeProfit,
		StopLoss:    order.StopLoss,
		Created:     order.Created,
	}
}

func ConvertPricesToMessage(prices *domain.Prices) *Prices {
	return &Prices{
		InSymbol:        prices.InSymbol,
		PricesBySymbols: prices.PricesBySymbols,
	}
}

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
	for i := range spot.Advice.BuyOrderSettings.TakeProfitOptions {
		takeProfitOptions = append(takeProfitOptions, Option{
			Value: spot.Advice.BuyOrderSettings.TakeProfitOptions[i],
			Text:  fmt.Sprintf("%.2f", spot.Advice.BuyOrderSettings.TakeProfitOptions[i]),
		})
	}

	var stopLossOptions []Option
	for i := range spot.Advice.BuyOrderSettings.StopLossOptions {
		stopLossOptions = append(stopLossOptions, Option{
			Value: spot.Advice.BuyOrderSettings.StopLossOptions[i],
			Text:  fmt.Sprintf("%.2f", spot.Advice.BuyOrderSettings.StopLossOptions[i]),
		})
	}

	return &Spot{
		Asset: Asset{
			Name:        spot.Asset.Name,
			Symbol:      spot.Asset.Symbol,
			Description: spot.Asset.Description,
		},
		ActiveOrders:  spot.ActiveOrders,
		PriceForecast: spot.Advice.PriceForecast,
		ChartsData:    chartsData,
		News:          news,
		BuyOrderSettings: BuyOrderSettings{
			Amount:            spot.Advice.BuyOrderSettings.Amount,
			TakeProfit:        spot.Advice.BuyOrderSettings.TakeProfit,
			TakeProfitOptions: takeProfitOptions,
			StopLoss:          spot.Advice.BuyOrderSettings.StopLoss,
			StopLossOptions:   stopLossOptions,
		},
	}
}

func buildChartsData(history []domain.Kline, forecast []domain.Kline) ChartsData {
	var chartsData ChartsData
	for i := range history {
		chartsData.Times = append(chartsData.Times, history[i].EndTime.Format("15:04"))
		chartsData.Prices = append(chartsData.Prices, history[i].High)
		chartsData.Volumes = append(chartsData.Volumes, int64(history[i].Volume))
	}
	for i := range forecast {
		chartsData.Times = append(chartsData.Times, forecast[i].EndTime.Format("15:04"))
		chartsData.Forecast = append(chartsData.Forecast, forecast[i].High)
	}

	return chartsData
}
