package gin

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	bud04 "github.com/sebdeveloper6952/blossom-server/src/bud-04"
	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func mirrorBlob(
	services core.Services,
	cdnBaseUrl string,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pubkey := ctx.GetString("pk")
		authSha256 := ctx.GetString("x")

		if pubkey == "" {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				apiError{
					Message: "pubkey missing from context",
				},
			)
		}

		if authSha256 == "" {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				apiError{
					Message: "blob hash missing from context",
				},
			)
		}

		mirrorInput := &mirrorInput{}
		if err := ctx.ShouldBindJSON(mirrorInput); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				apiError{
					Message: "invalid request body",
				},
			)
		}

		blobUrl, err := url.Parse(mirrorInput.Url)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				apiError{
					Message: "invalid blob URL",
				},
			)
		}

		blobDescriptor, err := bud04.MirrorBlob(
			ctx,
			services,
			cdnBaseUrl,
			pubkey,
			authSha256,
			*blobUrl,
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
			fromDomainBlobDescriptor(blobDescriptor),
		)
	}
}
