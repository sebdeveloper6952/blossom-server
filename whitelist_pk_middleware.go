package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func whitelistPkMiddleware(whitelistedPks map[string]struct{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := whitelistedPks[c.GetString("pk")]; !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
