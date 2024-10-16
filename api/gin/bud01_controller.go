package gin

import (
	"net/http"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	bud01 "github.com/sebdeveloper6952/blossom-server/src/bud-01"
	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func getBlob(
	services core.Services,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pathParts := strings.Split(ctx.Param("path"), ".")
		fileBytes, err := bud01.GetBlob(
			ctx.Request.Context(),
			services,
			pathParts[0],
		)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				apiError{
					Message: err.Error(),
				},
			)
			return
		}

		mType := mimetype.Detect(fileBytes)
		ctx.Header("Content-Type", mType.String())
		_, _ = ctx.Writer.Write(fileBytes)
		ctx.Status(http.StatusOK)
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
