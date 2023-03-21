package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"websmee/buyspot/internal/api"
)

func AddBalanceHandlers(
	_ context.Context,
	router *gin.Engine,
) {
	router.GET("/api/v1/balances/current", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, api.Balance{
			Ticker: "USDT",
			Amount: 1234.56,
		})
	})
}
