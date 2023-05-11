package admin

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases/admin"
)

func AddAdminHandlers(
	router *gin.Engine,
	marketDataUpdater *admin.MarketDataUpdater,
	newsUpdater *admin.NewsUpdater,
	userManager *admin.UserManager,
) {
	router.POST("/api/v1/admin/market_data/update/:period", func(c *gin.Context) {
		period, err := strconv.Atoi(c.Param("period"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := marketDataUpdater.Update(c, c.GetHeader("X-BUYSPOT-SECRET-KEY"), period); err != nil {
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

	router.POST("/api/v1/admin/users", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := userManager.Save(c, c.GetHeader("X-BUYSPOT-SECRET-KEY"), UserToDomain(&user)); err != nil {
			if errors.Is(err, domain.ErrForbidden) {
				c.Status(http.StatusForbidden)
				return
			}
			if errors.Is(err, domain.ErrInvalidArgument) {
				c.Status(http.StatusBadRequest)
				return
			}

			c.Error(fmt.Errorf("could not save user, err: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusOK, nil)
	})

	router.POST("/api/v1/admin/users/balances", func(c *gin.Context) {
		var balance Balance
		if err := c.BindJSON(&balance); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := userManager.CreateBalance(c, c.GetHeader("X-BUYSPOT-SECRET-KEY"), BalanceToDomain(&balance)); err != nil {
			if errors.Is(err, domain.ErrForbidden) {
				c.Status(http.StatusForbidden)
				return
			}

			c.Error(fmt.Errorf("could not create balance, err: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusOK, nil)
	})
}
