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

	bud01 "github.com/sebdeveloper6952/blossom-server/src/bud-01"
	bud02 "github.com/sebdeveloper6952/blossom-server/src/bud-02"
	bud04 "github.com/sebdeveloper6952/blossom-server/src/bud-04"
	bud06 "github.com/sebdeveloper6952/blossom-server/src/bud-06"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/hashing"
)

func uploadBlob(
	storage core.BlobStorage,
	mimeTypeService core.MimeTypeService,
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
			storage,
			mimeTypeService,
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

func uploadRequirements(
	mimeTypeService core.MimeTypeService,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		blobHash := ctx.GetHeader(HeaderXSHA256)
		if err := hashing.IsSHA256(blobHash); err != nil {
			ctx.Header(HeaderXUploadMessage, fmt.Sprintf("invalid SHA-256: %s", err))
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		contentType := ctx.GetHeader(HeaderContentType)
		if mimetype.Lookup(contentType) == nil {
			ctx.Header(HeaderXUploadMessage, "invalid Content-Type")
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		contentLength, err := strconv.Atoi(ctx.GetHeader(HeaderXContentLength))
		if err != nil {
			ctx.Header(HeaderXUploadMessage, "couldn't parse Content-Length as an integer")
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := bud06.UploadRequirements(
			ctx,
			mimeTypeService,
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

func mirrorBlob(
	storage core.BlobStorage,
	mimeTypeService core.MimeTypeService,
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
			storage,
			mimeTypeService,
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

func getBlob(
	storage core.BlobStorage,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pathParts := strings.Split(ctx.Param("path"), ".")
		fileBytes, err := bud01.GetBlob(
			ctx.Request.Context(),
			storage,
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
	storage core.BlobStorage,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pathParts := strings.Split(ctx.Param("path"), ".")
		_, err := bud01.HasBlob(
			ctx.Request.Context(),
			storage,
			pathParts[0],
		)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

func listBlobs(
	storage core.BlobStorage,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		blobs, err := bud02.ListBlobs(
			ctx.Request.Context(),
			storage,
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
	storage core.BlobStorage,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := bud02.DeleteBlob(
			ctx.Request.Context(),
			storage,
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
