package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebdeveloper6952/blossom-server/src/admin"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
)

type apiMimeType struct {
	Extension string `json:"ext"`
	MimeType  string `json:"mime_type"`
	Allowed   bool   `json:"allowed"`
}

type apiUpdateMimeTypeInput struct {
	MimeType string `json:"mime_type"`
	Allowed  bool   `json:"allowed"`
}

type apiSetting struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type apiUpdateSettingInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func fromCoreMimeType(m *core.MimeType) *apiMimeType {
	return &apiMimeType{
		Extension: m.Extension,
		MimeType:  m.MimeType,
		Allowed:   m.Allowed,
	}
}

func fromSliceCoreMimeType(ms []*core.MimeType) []*apiMimeType {
	apiMimeTypes := make([]*apiMimeType, len(ms))
	for i := range ms {
		apiMimeTypes[i] = fromCoreMimeType(ms[i])
	}

	return apiMimeTypes
}

func fromCoreSetting(m *core.Setting) *apiSetting {
	return &apiSetting{
		Key:   m.Key,
		Value: m.Value,
	}
}

func fromSliceCoreSetting(ms []*core.Setting) []*apiSetting {
	apiMimeTypes := make([]*apiSetting, len(ms))
	for i := range ms {
		apiMimeTypes[i] = fromCoreSetting(ms[i])
	}

	return apiMimeTypes
}

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

func adminGetMimeTypes(
	mimeTypeService core.MimeTypeService,
	log *zap.Logger,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mimeTypes, err := admin.GetMimeTypes(
			ctx.Request.Context(),
			mimeTypeService,
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
	mimeTypeService core.MimeTypeService,
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
			mimeTypeService,
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
	settingService core.SettingService,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		settings, err := settingService.GetAll(ctx.Request.Context())
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
	settingService core.SettingService,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := &apiSetting{}
		if err := ctx.BindJSON(body); err != nil {
			ctx.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}

		setting, err := settingService.Update(
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
