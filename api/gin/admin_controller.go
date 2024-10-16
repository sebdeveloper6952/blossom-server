package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebdeveloper6952/blossom-server/src/admin"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
)

func adminGetRules(
	services core.Services,
	_ *zap.Logger,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rules, err := admin.GetRules(ctx.Request.Context(), services)
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
			fromSliceCoreACR(rules),
		)
	}
}

func adminCreateRule(
	services core.Services,
	_ *zap.Logger,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := &createACRInput{}
		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}

		rule, err := admin.CreateRule(
			ctx.Request.Context(),
			services,
			core.ACRAction(body.Action),
			body.Pubkey,
			core.ACRResource(body.Resource),
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
			http.StatusCreated,
			fromCoreACR(rule),
		)
	}
}

func adminDeleteRule(
	services core.Services,
	_ *zap.Logger,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		action := ctx.Query("action")
		pk := ctx.Query("pubkey")
		res := ctx.Query("resource")

		if err := admin.DeleteRule(
			ctx.Request.Context(),
			services,
			core.ACRAction(action),
			pk,
			core.ACRResource(res),
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

func adminGetMimeTypes(
	services core.Services,
	log *zap.Logger,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mimeTypes, err := admin.GetMimeTypes(
			ctx.Request.Context(),
			services,
			log,
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
			fromSliceCoreMimeType(mimeTypes),
		)
	}
}

func adminUpdateMimeType(
	services core.Services,
	log *zap.Logger,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := &apiUpdateMimeTypeInput{}
		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}

		if err := admin.UpdateMimeType(
			ctx.Request.Context(),
			services,
			body.MimeType,
			body.Allowed,
			log,
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

func adminMiddleware(adminPubkey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authedPk := ctx.GetString("pk")

		if authedPk != adminPubkey {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}

func adminGetSettings(
	services core.Services,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		settings, err := services.Settings().GetAll(ctx.Request.Context())
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.JSON(
			http.StatusOK,
			fromSliceCoreSetting(settings),
		)
	}
}

func adminUpdateSetting(
	services core.Services,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := &apiSetting{}
		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}

		setting, err := services.Settings().Update(
			ctx.Request.Context(),
			body.Key,
			body.Value,
		)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.JSON(
			http.StatusOK,
			fromCoreSetting(setting),
		)
	}
}
