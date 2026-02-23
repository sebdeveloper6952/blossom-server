package httpapi

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	bud01 "github.com/sebdeveloper6952/blossom-server/internal/bud01"
	"github.com/sebdeveloper6952/blossom-server/internal/core"
)

func getBlob(
	services core.Services,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pathParts := strings.Split(ctx.Param("path"), ".")
		hash := pathParts[0]

		blob, err := bud01.GetBlob(
			ctx.Request.Context(),
			services,
			hash,
		)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusNotFound,
				apiError{
					Message: err.Error(),
				},
			)
			return
		}

		name := hash
		if len(pathParts) > 1 {
			name = hash + "." + pathParts[1]
		}

		ctx.Header("Content-Type", blob.Type)
		http.ServeContent(
			ctx.Writer,
			ctx.Request,
			name,
			time.Unix(blob.Uploaded, 0),
			bytes.NewReader(blob.Blob),
		)
	}
}

func hasBlob(
	services core.Services,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pathParts := strings.Split(ctx.Param("path"), ".")
		_, err := bud01.HasBlob(
			ctx.Request.Context(),
			services,
			pathParts[0],
		)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
