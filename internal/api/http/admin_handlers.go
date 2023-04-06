package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases/admin"
)

func AddAdminHandlers(
	router *gin.Engine,
	marketDataUpdater *admin.MarketDataUpdater,
	newsUpdater *admin.NewsUpdater,
) {
	router.POST("/api/v1/admin/market_data/update", func(c *gin.Context) {
		if err := marketDataUpdater.Update(c, c.GetHeader("X-BUYSPOT-SECRET-KEY")); err != nil {
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

	router.POST("/api/v1/admin/news/update", func(c *gin.Context) {
		if err := newsUpdater.Update(c, c.GetHeader("X-BUYSPOT-SECRET-KEY")); err != nil {
			if errors.Is(err, domain.ErrForbidden) {
				c.Status(http.StatusForbidden)
				return
			}

			c.Error(fmt.Errorf("could not update news, err: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusOK, nil)
	})
}
