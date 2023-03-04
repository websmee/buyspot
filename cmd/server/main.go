package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/api/v1/balances/current", getCurrentBalance)
	router.GET("/api/v1/spots/next", getNextSpot)
	router.POST("/api/v1/spots/buy", buySpot)

	router.Run("localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func getCurrentBalance(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Balance{
		Ticker: "USDT",
		Amount: 1234.56,
	})
}

func getNextSpot(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Spot{
		Asset: Asset{
			Name:         "Bitcoin",
			Ticker:       "BTC",
			Description:  "Bitcoin (abbreviation: BTC[a] or XBT[b]; sign: ₿) is a protocol which implements a highly available, public, permanent, and decentralized ledger. In order to add to the ledger, a user must prove they control an entry in the ledger. The protocol specifies that the entry indicates an amount of a token, bitcoin with a minuscule b. The user can update the ledger, assigning some of their bitcoin to another entry in the ledger. Because the token has characteristics of money, it can be thought of as a digital currency.",
			ActiveOrders: 3,
		},
		PriceForecast: 3.45,
		ChartsData: ChartsData{
			Times: []string{
				"00:00", "01:00", "02:00", "03:00", "04:00", "05:00", "06:00", "07:00",
				"08:00", "09:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00",
				"16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00",
			},
			Prices: []float64{
				21234.12, 21224.23, 21214.56, 21264.78, 21214.90, 21134.12, 21154.34, 21164.56,
				21174.56, 21184.56, 21214.56, 21224.56, 21234.56, 21244.56, 21264.56, 21284.56,
				21319.56,
			},
			Forecast: []float64{
				21319.56, 21344.56, 21374.56, 21420.56, 21515.56, 21624.56, 21744.56, 21850.56,
			},
			Volumes: []float64{
				5000, 4000, 6000, 5000, 6000, 4000, 3000, 2000,
				5000, 4000, 6000, 5000, 6000, 7000, 8000, 9000,
				10000,
			},
		},
		CurrentSpotsIndex: 3,
		CurrentSpotsTotal: 10,
		News: []NewsArticle{
			{
				Sentiment: NewsArticleSentimentNeutral,
				Title:     "Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain",
				Content:   "Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain. Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain. Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain. Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain",
				Created:   time.Now().Add(-1 * time.Hour),
				Views:     15678,
			},
			{
				Sentiment: NewsArticleSentimentPositive,
				Title:     "IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed",
				Content:   "IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed. IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed. IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed. IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed",
				Created:   time.Now().Add(-12 * time.Hour),
				Views:     25678,
			},
			{
				Sentiment: NewsArticleSentimentNegative,
				Title:     "Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown",
				Content:   "Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown. Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown. Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown. Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown",
				Created:   time.Now().Add(-48 * time.Hour),
				Views:     178,
			},
		},
		BuyOrderSettings: BuyOrderSettings{
			Amount:     100,
			TakeProfit: 4,
			TakeProfitOptions: []Option{
				{3, "+3%"},
				{4, "+4%"},
				{5, "+5%"},
			},
			StopLoss: -2,
			StopLossOptions: []Option{
				{-1, "-1%"},
				{-2, "-2%"},
				{-3, "-3%"},
			},
		},
	})
}

func buySpot(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, BuySpotResponse{
		UpdatedBalance: Balance{
			Ticker: "USDT",
			Amount: 1134.56,
		},
	})
}

type Balance struct {
	Ticker string  `json:"ticker"`
	Amount float64 `json:"amount"`
}

type Asset struct {
	Name         string `json:"name"`
	Ticker       string `json:"ticker"`
	Description  string `json:"description"`
	ActiveOrders int    `json:"activeOrders"`
}

type ChartsData struct {
	Times    []string  `json:"times"`
	Prices   []float64 `json:"prices"`
	Forecast []float64 `json:"forecast"`
	Volumes  []float64 `json:"volumes"`
}

type NewsArticle struct {
	Sentiment NewsArticleSentiment `json:"sentiment"`
	Title     string               `json:"title"`
	Content   string               `json:"content"`
	Created   time.Time            `json:"created"`
	Views     int                  `json:"views"`
}

type NewsArticleSentiment string

const (
	NewsArticleSentimentNeutral  NewsArticleSentiment = "NEUTRAL"
	NewsArticleSentimentPositive NewsArticleSentiment = "POSITIVE"
	NewsArticleSentimentNegative NewsArticleSentiment = "NEGATIVE"
)

type Option struct {
	Value float64 `json:"value"`
	Text  string  `json:"text"`
}

type BuyOrderSettings struct {
	Amount            float64  `json:"amount"`
	TakeProfit        float64  `json:"takeProfit"`
	TakeProfitOptions []Option `json:"takeProfitOptions"`
	StopLoss          float64  `json:"stopLoss"`
	StopLossOptions   []Option `json:"stopLossOptions"`
}

type Spot struct {
	Asset             Asset            `json:"asset"`
	PriceForecast     float64          `json:"priceForecast"`
	ChartsData        ChartsData       `json:"chartsData"`
	CurrentSpotsIndex int              `json:"currentSpotsIndex"`
	CurrentSpotsTotal int              `json:"currentSpotsTotal"`
	News              []NewsArticle    `json:"news"`
	BuyOrderSettings  BuyOrderSettings `json:"buyOrderSettings"`
}

type BuySpotResponse struct {
	UpdatedBalance Balance `json:"updatedBalance"`
}
