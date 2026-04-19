package httpapi

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	bud02 "github.com/sebdeveloper6952/blossom-server/v2/internal/bud02"
	"github.com/sebdeveloper6952/blossom-server/v2/internal/core"
)

type adminBlobDescriptor struct {
	Pubkey   string `json:"pubkey"`
	Url      string `json:"url"`
	Sha256   string `json:"sha256"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
	Uploaded int64  `json:"uploaded"`
}

func fromDomainAdminBlobDescriptor(blob *core.Blob) adminBlobDescriptor {
	return adminBlobDescriptor{
		Pubkey:   blob.Pubkey,
		Url:      blob.Url,
		Sha256:   blob.Sha256,
		Size:     blob.Size,
		Type:     blob.Type,
		Uploaded: blob.Uploaded,
	}
}

func listAllBlobs(services core.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var since int64
		var until int64 = math.MaxInt64

		if s := ctx.Query("since"); s != "" {
			if v, err := strconv.ParseInt(s, 10, 64); err == nil {
				since = v
			}
		}
		if u := ctx.Query("until"); u != "" {
			if v, err := strconv.ParseInt(u, 10, 64); err == nil {
				until = v
			}
		}

		blobs, err := bud02.ListAllBlobs(
			ctx.Request.Context(),
			services,
			since,
			until,
		)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				apiError{Message: err.Error()},
			)
			return
		}

		out := make([]adminBlobDescriptor, len(blobs))
		for i := range blobs {
			out[i] = fromDomainAdminBlobDescriptor(blobs[i])
		}

		ctx.JSON(http.StatusOK, out)
	}
}
