package gin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	bud06 "github.com/sebdeveloper6952/blossom-server/src/bud-06"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/hashing"
)

func uploadRequirements(
	services core.Services,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		blobHash := ctx.GetHeader(HeaderXSHA256)
		if err := hashing.IsSHA256(blobHash); err != nil {
			ctx.Header(HeaderXUploadMessage, fmt.Sprintf("invalid SHA-256: %s", err))
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		contentType := ctx.GetHeader(HeaderXContentType)
		contentLength, err := strconv.Atoi(ctx.GetHeader(HeaderXContentLength))
		if err != nil {
			ctx.Header(HeaderXUploadMessage, "couldn't parse Content-Length as an integer")
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := bud06.UploadRequirements(
			ctx,
			services,
			blobHash,
			contentType,
			contentLength,
		); err != nil {
			ctx.Header(HeaderXUploadMessage, err.Error())
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.Status(
			http.StatusOK,
		)
	}
}
