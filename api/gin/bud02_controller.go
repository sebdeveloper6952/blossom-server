package gin

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	bud02 "github.com/sebdeveloper6952/blossom-server/src/bud-02"
	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func uploadBlob(
	services core.Services,
	cdnBaseUrl string,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		defer func(body io.ReadCloser) {
			err := body.Close()
			if err != nil {

			}
		}(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				apiError{
					Message: fmt.Sprintf("failed to read request body: %s", err.Error()),
				},
			)
			return
		}

		blobDescriptor, err := bud02.UploadBlob(
			ctx.Request.Context(),
			services,
			cdnBaseUrl,
			ctx.GetString("x"),
			ctx.GetString("pk"),
			bodyBytes,
		)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				apiError{
					Message: fmt.Sprintf("%s", err.Error()),
				},
			)
			return
		}

		ctx.JSON(
			http.StatusOK,
			fromDomainBlobDescriptor(blobDescriptor),
		)
	}
}

func listBlobs(
	services core.Services,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		blobs, err := bud02.ListBlobs(
			ctx.Request.Context(),
			services,
			ctx.Param("pubkey"),
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

		ctx.JSON(
			http.StatusOK,
			fromSliceDomainBlobDescriptor(blobs),
		)
	}
}

func deleteBlob(
	services core.Services,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := bud02.DeleteBlob(
			ctx.Request.Context(),
			services,
			ctx.GetString("pk"),
			ctx.Param("path"),
			ctx.GetString("x"),
		); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				apiError{
					Message: err.Error(),
				},
			)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
