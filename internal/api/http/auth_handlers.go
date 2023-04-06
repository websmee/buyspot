package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"websmee/buyspot/internal/api"
)

func AddAuthHandlers(
	router *gin.Engine,
	auth *SimpleAuth,
) {
	router.POST("/api/v1/login", func(c *gin.Context) {
		var request api.LoginRequest
		if err := c.BindJSON(&request); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if !auth.CheckCredentials(request.Username, request.Password) {
			c.Status(http.StatusUnauthorized)
			return
		}

		c.IndentedJSON(http.StatusOK, auth.GetToken(request.Username))
	})
}
