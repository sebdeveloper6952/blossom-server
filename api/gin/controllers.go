package gin

import (
	"io"
	"net/http"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"

	"github.com/sebdeveloper6952/blossom-server/application"
	"github.com/sebdeveloper6952/blossom-server/domain"
	"github.com/sebdeveloper6952/blossom-server/services"
)

func UploadBlob(
	blobRepo domain.BlobDescriptorRepo,
	hasher services.Hashing,
	cdnBaseUrl string,
) gin.HandlerFunc {
	uploadBlob := application.UploadBlob(
		blobRepo,
		hasher,
		cdnBaseUrl,
	)

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
				gin.H{
					"message": err.Error(),
				},
			)
			return
		}

		blobDescriptor, err := uploadBlob(
			ctx.Request.Context(),
			ctx.GetString("pk"),
			bodyBytes,
		)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": err.Error(),
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

func GetBlob(
	blobRepo domain.BlobDescriptorRepo,
) gin.HandlerFunc {
	getBlob := application.GetBlob(blobRepo)

	return func(ctx *gin.Context) {
		pathParts := strings.Split(ctx.Param("path"), ".")
		fileBytes, err := getBlob(
			ctx.Request.Context(),
			pathParts[0],
		)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": err.Error(),
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

func HasBlob(
	blobRepo domain.BlobDescriptorRepo,
) gin.HandlerFunc {
	hasBlob := application.HasBlob(blobRepo)

	return func(ctx *gin.Context) {
		pathParts := strings.Split(ctx.Param("path"), ".")
		_, err := hasBlob(
			ctx.Request.Context(),
			pathParts[0],
		)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

func ListBlobs(
	blobRepo domain.BlobDescriptorRepo,
) gin.HandlerFunc {
	listBlobs := application.ListBlobs(blobRepo)
	return func(ctx *gin.Context) {
		blobs, err := listBlobs(
			ctx.Request.Context(),
			ctx.Param("pubkey"),
		)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": err.Error(),
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

func DeleteBlob(
	blobRepo domain.BlobDescriptorRepo,
) gin.HandlerFunc {
	deleteBlob := application.DeleteBlob(blobRepo)

	return func(ctx *gin.Context) {
		if err := deleteBlob(
			ctx.Request.Context(),
			ctx.Param("path"),
			ctx.GetString("x"),
			ctx.GetString("pk"),
		); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": err.Error(),
				},
			)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
