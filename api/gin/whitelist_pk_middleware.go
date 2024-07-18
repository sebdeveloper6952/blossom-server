package gin

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func whitelistPkMiddleware(whitelistedPks map[string]struct{}, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := whitelistedPks[c.GetString("pk")]; len(whitelistedPks) > 0 && !ok {
			log.Debug("[whitelistPkMiddleware] pubkey not in whitelist")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
