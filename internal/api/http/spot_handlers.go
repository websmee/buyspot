package http

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

func AddSpotHandlers(
	router *gin.Engine,
	spotReader *usecases.SpotReader,
	spotBuyer *usecases.SpotBuyer,
) {
	router.GET("/api/v1/spots/data", func(c *gin.Context) {
		count, err := spotReader.GetSpotsCount(c)
		if err != nil {
			c.Error(fmt.Errorf("could not get spots count, err: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusOK, SpotsData{CurrentSpotsTotal: count})
	})

	router.GET("/api/v1/spots/:index", func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		spot, err := spotReader.GetSpotByIndex(c, index)
		if err != nil {
			if errors.Is(err, domain.ErrUnauthorized) {
				c.Status(http.StatusUnauthorized)
				return
			}

			if errors.Is(err, domain.ErrSpotNotFound) {
				c.Status(http.StatusNotFound)
				return
			}

			c.Error(fmt.Errorf("could not get spot by index %d, err: %w", index, err))
			c.Status(http.StatusInternalServerError)
			return
		}

		msg := ConvertSpotToMessage(spot)
		msg.Index = index
		c.IndentedJSON(http.StatusOK, msg)
	})

	router.POST("/api/v1/spots/buy", func(c *gin.Context) {
		var buySpotRequest BuySpotRequest
		if err := c.BindJSON(&buySpotRequest); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		updatedBalance, err := spotBuyer.BuySpot(
			c,
			buySpotRequest.SpotID,
			buySpotRequest.Amount,
			buySpotRequest.Symbol,
			buySpotRequest.TakeProfit,
			buySpotRequest.StopLoss,
		)
		if err != nil {
			if errors.Is(err, domain.ErrUnauthorized) {
				c.Status(http.StatusUnauthorized)
				return
			}

			c.Error(fmt.Errorf("could not buy spot, err: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusOK, BuySpotResponse{
			UpdatedBalance: Balance{
				Symbol: updatedBalance.Symbol,
				Amount: updatedBalance.Amount,
			},
		})
	})
}
