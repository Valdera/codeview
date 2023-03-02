package middleware

import (
	"codeview/config"
	"codeview/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SessionAuth(cfg config.AppConfig) gin.HandlerFunc {

	return func(c *gin.Context) {
		_, err := util.GetTokenSession(c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}

func JWTAuth(cfg config.AppConfig) gin.HandlerFunc {
	jwtManager := util.NewJWTManager(cfg)

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		claims, err := jwtManager.Verify(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("claims", claims)
		c.Next()
	}
}
