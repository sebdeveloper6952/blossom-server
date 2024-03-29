package main

import (
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

func Upload(
	server Server,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		defer ctx.Request.Body.Close()
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": err.Error(),
				},
			)
			return
		}

		blobDescriptor, err := server.UploadBlob(
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
			blobDescriptor,
		)
	}
}

func GetBlob(
	server Server,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pathParts := strings.Split(ctx.Param("path"), ".")
		fileBytes, err := server.GetBlob(
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
		}

		mType := mimetype.Detect(fileBytes)
		ctx.Header("Content-Type", mType.String())
		_, _ = ctx.Writer.Write(fileBytes)
		ctx.Status(http.StatusOK)
	}
}

func HasBlob(
	server Server,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pathParts := strings.Split(ctx.Param("path"), ".")
		_, err := server.HasBlob(
			ctx.Request.Context(),
			pathParts[0],
		)
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
		}

		ctx.Status(http.StatusOK)
	}
}

func ListBlobs(
	server Server,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		blobs, err := server.ListBlobs(
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
			blobs,
		)
	}
}

func DeleteBlob(
	server Server,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := server.DeleteBlob(
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
