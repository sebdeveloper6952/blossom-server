package main

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
)

func whitelistPkMiddleware(whitelistedPks map[string]struct{}, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := whitelistedPks[c.GetString("pk")]; !ok {
			log.Info("[whitelistPkMiddleware] pubkey not in whitelist")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
