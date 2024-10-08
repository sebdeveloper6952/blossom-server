package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebdeveloper6952/blossom-server/src/admin"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
)

func adminGetRules(
	ac core.ACRStorage,
	_ *zap.Logger,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rules, err := admin.GetRules(ctx.Request.Context(), ac)
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
	ac core.ACRStorage,
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
			ac,
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
	ac core.ACRStorage,
	_ *zap.Logger,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		action := ctx.Query("action")
		pk := ctx.Query("pubkey")
		res := ctx.Query("resource")

		if err := admin.DeleteRule(
			ctx.Request.Context(),
			ac,
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
