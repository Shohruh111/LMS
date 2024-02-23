package handler

import (
	"app/pkg/helper"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (h *handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerToken := c.GetHeader("Authorization")
		if len(bearerToken) <= 0 {
			c.AbortWithError(http.StatusUnauthorized, errors.New("User not authentication"))
			return
		}

		token, err := helper.ExtractToken(cast.ToString(bearerToken))
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		authInfo, err := helper.ParseClaims(token, h.cfg.SecretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("user_id", authInfo["user_id"])
		// c.Set("client_type", authInfo["client_type"])

		c.Next()
	}
}
