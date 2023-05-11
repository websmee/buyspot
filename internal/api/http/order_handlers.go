package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

func AddOrderHandlers(
	router *gin.Engine,
	orderReader *usecases.OrderReader,
	orderSeller *usecases.OrderSeller,
) {
	router.GET("/api/v1/orders", func(c *gin.Context) {
		orders, err := orderReader.GetUserOrders(c)
		if err != nil {
			if errors.Is(err, domain.ErrUnauthorized) {
				c.Status(http.StatusUnauthorized)
				return
			}

			c.Error(fmt.Errorf("could not get orders, err: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		ordersMessages := make([]Order, 0, len(orders))
		for i := range orders {
			ordersMessages = append(ordersMessages, *ConvertOrderToMessages(&orders[i]))
		}

		c.IndentedJSON(http.StatusOK, ordersMessages)
	})

	router.POST("/api/v1/orders/:orderID/sell", func(c *gin.Context) {
		orderID := c.Param("orderID")
		updatedBalance, err := orderSeller.SellOrder(c, orderID)
		if err != nil {
			if errors.Is(err, domain.ErrUnauthorized) {
				c.Status(http.StatusUnauthorized)
				return
			}

			c.Error(fmt.Errorf("could not sell order '%s', err: %w", orderID, err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusOK, SellOrderResponse{
			OrderID: orderID,
			UpdatedBalance: Balance{
				Symbol: updatedBalance.Symbol,
				Amount: updatedBalance.Amount,
			},
		})
	})
}
