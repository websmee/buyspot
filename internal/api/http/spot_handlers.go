package http

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"websmee/buyspot/internal/api"
	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

func AddSpotHandlers(
	router *gin.Engine,
	spotReader *usecases.SpotReader,
) {
	router.GET("/api/v1/spots/data", func(c *gin.Context) {
		count, err := spotReader.GetSpotsCount(c)
		if err != nil {
			c.Error(fmt.Errorf("could get spots count, err: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusOK, api.SpotsData{CurrentSpotsTotal: count})
	})

	router.GET("/api/v1/spots/:index", func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		spot, err := spotReader.GetSpotByIndex(c, index)
		if err != nil {
			if errors.Is(err, domain.ErrSpotNotFound) {
				c.Status(http.StatusNotFound)
				return
			}

			c.Error(fmt.Errorf("could get spot by index %d, err: %w", index, err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusOK, api.ConvertSpotToMessage(spot))
	})

	router.POST("/api/v1/spots/buy", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, api.BuySpotResponse{
			UpdatedBalance: api.Balance{
				Ticker: "USDT",
				Amount: 1134.56,
			},
		})
	})
}
