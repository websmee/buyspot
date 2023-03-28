package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"websmee/buyspot/internal/api"
)

func AddBalanceHandlers(
	router *gin.Engine,
) {
	router.GET("/api/v1/balances/current", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, api.Balance{
			Symbol: "USDT",
			Amount: 1234.56,
		})
	})
}
