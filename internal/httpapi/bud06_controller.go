package httpapi

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	bud06 "github.com/sebdeveloper6952/blossom-server/internal/bud06"
	"github.com/sebdeveloper6952/blossom-server/internal/core"
	"github.com/sebdeveloper6952/blossom-server/internal/pkg/hashing"
	"github.com/sebdeveloper6952/blossom-server/internal/service"
)

func uploadRequirements(
	services core.Services,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		blobHash := ctx.GetHeader(HeaderXSHA256)
		if err := hashing.IsSHA256(blobHash); err != nil {
			ctx.Header(HeaderXReason, fmt.Sprintf("invalid SHA-256: %s", err))
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		contentType := ctx.GetHeader(HeaderXContentType)

		contentLengthStr := ctx.GetHeader(HeaderXContentLength)
		if contentLengthStr == "" {
			ctx.Header(HeaderXReason, "missing X-Content-Length header")
			ctx.AbortWithStatus(http.StatusLengthRequired)
			return
		}

		contentLength, err := strconv.Atoi(contentLengthStr)
		if err != nil {
			ctx.Header(HeaderXReason, "couldn't parse X-Content-Length as an integer")
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		authedPubkey := ctx.GetString("pk")
		if err := bud06.UploadRequirements(
			ctx,
			services,
			authedPubkey,
			blobHash,
			contentType,
			contentLength,
		); err != nil {
			ctx.Header(HeaderXReason, err.Error())
			ctx.AbortWithStatus(mapBud06Error(err))
			return
		}

		ctx.Status(http.StatusOK)
	}
}

func mapBud06Error(err error) int {
	switch {
	case errors.Is(err, service.ErrUnauthorized):
		return http.StatusForbidden
	case errors.Is(err, core.ErrMimeTypeNotAllowed):
		return http.StatusUnsupportedMediaType
	case errors.Is(err, core.ErrFileSizeLimit):
		return http.StatusRequestEntityTooLarge
	default:
		return http.StatusBadRequest
	}
}
