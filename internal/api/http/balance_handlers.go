package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

func AddBalanceHandlers(
	router *gin.Engine,
	balanceReader *usecases.BalanceReader,
) {
	router.GET("/api/v1/balances/current", func(c *gin.Context) {
		balance, err := balanceReader.GetActiveBalance(c)
		if err != nil {
			if errors.Is(err, domain.ErrUnauthorized) {
				c.Status(http.StatusUnauthorized)
				return
			}

			c.Error(fmt.Errorf("could not get user balance, err: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusOK, ConvertBalanceToMessage(balance))
	})
}
