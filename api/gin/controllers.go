package gin

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"

	"github.com/sebdeveloper6952/blossom-server/application"
	"github.com/sebdeveloper6952/blossom-server/domain"
	"github.com/sebdeveloper6952/blossom-server/utils"
)

func UploadBlob(
	blobRepo domain.BlobDescriptorRepo,
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

		blobDescriptor, err := application.UploadBlob(
			ctx.Request.Context(),
			blobRepo,
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

const (
	HeaderSha256        = "X-SHA-256"
	HeaderContentType   = "X-Content-Type"
	HeaderContentLength = "X-Content-Length"
	HeaderUploadMessage = "X-Upload-Message"
)

func UploadRequirements() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		blobHash := ctx.GetHeader(HeaderSha256)
		if err := utils.IsSHA256(blobHash); err != nil {
			ctx.Header(HeaderUploadMessage, fmt.Sprintf("invalid SHA-256: %s", err))
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		contentType := ctx.GetHeader(HeaderContentType)
		if contentType == "" || mimetype.Lookup(contentType) == nil {
			ctx.Header(HeaderUploadMessage, "invalid Content-Type")
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		contentLength, err := strconv.Atoi(ctx.GetHeader(HeaderContentLength))
		if err != nil {
			ctx.Header(HeaderUploadMessage, "couldn't parse Content-Length as an integer")
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := application.UploadRequirements(
			ctx,
			blobHash,
			contentType,
			contentLength,
		); err != nil {
			ctx.Header(HeaderUploadMessage, err.Error())
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.Status(
			http.StatusOK,
		)
	}
}

func MirrorBlob(
	blobRepo domain.BlobDescriptorRepo,
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

		blobDescriptor, err := application.MirrorBlob(
			ctx,
			blobRepo,
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
	return func(ctx *gin.Context) {
		pathParts := strings.Split(ctx.Param("path"), ".")
		fileBytes, err := application.GetBlob(
			ctx.Request.Context(),
			blobRepo,
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

func HasBlob(
	blobRepo domain.BlobDescriptorRepo,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pathParts := strings.Split(ctx.Param("path"), ".")
		_, err := application.HasBlob(
			ctx.Request.Context(),
			blobRepo,
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
	return func(ctx *gin.Context) {
		blobs, err := application.ListBlobs(
			ctx.Request.Context(),
			blobRepo,
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

func DeleteBlob(
	blobRepo domain.BlobDescriptorRepo,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := application.DeleteBlob(
			ctx.Request.Context(),
			blobRepo,
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
