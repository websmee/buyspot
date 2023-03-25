package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"websmee/buyspot/internal/api"
	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

func AddPricesHandlers(
	router *gin.Engine,
	pricesReader *usecases.PricesReader,
) {
	router.GET("/api/v1/prices/current", func(c *gin.Context) {
		prices, err := pricesReader.GetCurrentPrices(c)
		if err != nil {
			if errors.Is(err, domain.ErrUnauthorized) {
				c.Status(http.StatusUnauthorized)
				return
			}

			c.Error(fmt.Errorf("could not get current prices, err: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusOK, api.ConvertPricesToMessage(prices))
	})
}
