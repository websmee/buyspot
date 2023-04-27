package http

import (
	"github.com/gin-gonic/gin"

	"websmee/buyspot/internal/domain"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthMiddleware(auth *Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.Set(domain.CtxKeyUserID, "")
			return
		}

		userID, err := auth.GetUserIDByToken(token)
		if err != nil {
			c.Error(err)
		}

		c.Set(domain.CtxKeyUserID, userID)
	}
}
