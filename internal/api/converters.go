package api

import (
	"fmt"

	"websmee/buyspot/internal/domain"
)

func UserToDomain(user *User) *domain.User {
	return &domain.User{
		Email:            user.Email,
		Password:         user.Password,
		BinanceAPIKey:    user.BinanceAPIKey,
		BinanceSecretKey: user.BinanceSecretKey,
	}
}

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
		Quote:           prices.Quote,
		PricesBySymbols: prices.PricesBySymbols,
	}
}

func ConvertSpotToMessage(spot *domain.Spot) *Spot {
	chartsData := buildChartsData(
		spot.HistoryMarketDataByQuotes,
		spot.ForecastMarketDataByQuotes,
		spot.ActualMarketDataByQuotes,
	)

	var news []NewsArticle
	for i := range spot.News {
		news = append(news, NewsArticle{
			Sentiment: ConvertNewsSentiment(spot.News[i].Sentiment),
			Title:     spot.News[i].Title,
			Content:   spot.News[i].Content,
			Created:   spot.News[i].Created,
			URL:       spot.News[i].URL,
			ImgURL:    spot.News[i].ImageURL,
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
		ActiveOrders:       spot.ActiveOrders,
		PriceForecast:      spot.Advice.PriceForecast,
		ChartsDataByQuotes: chartsData,
		News:               news,
		BuyOrderSettings: BuyOrderSettings{
			Amount:            spot.Advice.BuyOrderSettings.Amount,
			TakeProfit:        spot.Advice.BuyOrderSettings.TakeProfit,
			TakeProfitOptions: takeProfitOptions,
			StopLoss:          spot.Advice.BuyOrderSettings.StopLoss,
			StopLossOptions:   stopLossOptions,
		},
	}
}

func ConvertNewsSentiment(sentiment domain.NewsArticleSentiment) NewsArticleSentiment {
	switch sentiment {
	case domain.NewsArticleSentimentNeutral:
		return NewsArticleSentimentNeutral
	case domain.NewsArticleSentimentPositive:
		return NewsArticleSentimentPositive
	case domain.NewsArticleSentimentNegative:
		return NewsArticleSentimentNegative
	default:
		return NewsArticleSentimentNeutral
	}
}

func buildChartsData(
	historyByQuotes map[string][]domain.Kline,
	forecastByQuotes map[string][]domain.Kline,
	actualByQuotes map[string][]domain.Kline,
) map[string]ChartsData {
	chartsDataByQuotes := make(map[string]ChartsData)
	for quote := range historyByQuotes {
		var data ChartsData
		for i := range historyByQuotes[quote] {
			data.Times = append(data.Times, historyByQuotes[quote][i].EndTime.Format("15:04"))
			data.Prices = append(data.Prices, historyByQuotes[quote][i].Close)
			data.Volumes = append(data.Volumes, int64(historyByQuotes[quote][i].Volume))
		}
		for i := range forecastByQuotes[quote] {
			data.Times = append(data.Times, forecastByQuotes[quote][i].EndTime.Format("15:04"))
			data.Forecast = append(data.Forecast, forecastByQuotes[quote][i].Close)
		}
		for i := range actualByQuotes[quote] {
			data.Actual = append(data.Actual, actualByQuotes[quote][i].Close)
		}
		chartsDataByQuotes[quote] = data
	}

	return chartsDataByQuotes
}
