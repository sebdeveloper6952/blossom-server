package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type meResponse struct {
	Pubkey  string `json:"pubkey"`
	IsAdmin bool   `json:"is_admin"`
}

func getMe(adminPubkey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pk := ctx.GetString("pk")
		ctx.JSON(http.StatusOK, meResponse{
			Pubkey:  pk,
			IsAdmin: adminPubkey != "" && pk == adminPubkey,
		})
	}
}
