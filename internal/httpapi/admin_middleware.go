package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func adminOnlyMiddleware(adminPubkey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if adminPubkey == "" || c.GetString("pk") != adminPubkey {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}
