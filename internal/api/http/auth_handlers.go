package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddAuthHandlers(
	router *gin.Engine,
	auth *Auth,
) {
	router.POST("/api/v1/login", func(c *gin.Context) {
		var request LoginRequest
		if err := c.BindJSON(&request); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		user, err := auth.CheckCredentials(c, request.Email, request.Password)
		if err != nil {
			if errors.Is(err, ErrInvalidCredentials) {
				c.Status(http.StatusUnauthorized)
				return
			}

			c.Error(fmt.Errorf("could not check credentials, err: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		token, err := auth.GetToken(user)
		if err != nil {
			c.Error(fmt.Errorf("could not generate token, err: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusOK, token)
	})
}
