package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebdeveloper6952/blossom-server/src/core"
)

type apiStat struct {
	BytesStored int `json:"bytes_stored"`
	BlobCount   int `json:"blob_count"`
	PubkeyCount int `json:"pubkey_count"`
}

func fromCoreStat(s *core.Stats) apiStat {
	return apiStat{
		BytesStored: s.BytesStored,
		BlobCount:   s.BlobCount,
		PubkeyCount: s.PubkeyCount,
	}
}

func getStats(
	statService core.StatService,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		stats, err := statService.Get(ctx.Request.Context())
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.JSON(
			http.StatusOK,
			fromCoreStat(stats),
		)
	}
}
