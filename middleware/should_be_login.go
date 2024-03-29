package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShouldBeLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get(CTX_AUTH_KEY); exists {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": "Unauthorized",
			})
		}
	}
}
