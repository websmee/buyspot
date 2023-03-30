package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

func AddMarketDataHandlers(
	router *gin.Engine,
	updater *usecases.MarketDataUpdater,
) {
	router.POST("/api/v1/market_data/update", func(c *gin.Context) {
		if err := updater.Update(c, c.GetHeader("X-BUYSPOT-SECRET-KEY")); err != nil {
			if errors.Is(err, domain.ErrForbidden) {
				c.Status(http.StatusForbidden)
				return
			}

			c.Error(fmt.Errorf("could not update market data, err: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusOK, nil)
	})
}
