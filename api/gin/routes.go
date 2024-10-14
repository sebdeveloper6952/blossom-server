package gin

import (
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
)

func SetupRoutes(
	blobService core.BlobStorage,
	acrService core.ACRStorage,
	settingService core.SettingService,
	mimeTypeService core.MimeTypeService,
	cdnBaseUrl string,
	apiAddress string,
	adminPubkey string,
	uiEnabled bool,
	log *zap.Logger,
) Api {
	r := gin.New()

	r.Use(ginzap.Ginzap(log, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(log, true))

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "PUT", "HEAD", "DELETE"},
		AllowHeaders: []string{
			HeaderAuthorization,
			HeaderContentType,
			HeaderXSHA256,
			HeaderXContentType,
			HeaderXContentLength,
		},
		ExposeHeaders: []string{"Content-Length"},
	}))

	// serve ui
	if uiEnabled {
		r.StaticFile("/ui", "./ui/build/200.html")
		r.Static("/ui", "./ui/build")
		r.StaticFile("favicon.png", "./ui/build/favicon.png")
	}

	r.HEAD(
		"/upload",
		nostrAuthMiddleware("upload", log),
		accessControlMiddleware(acrService, "UPLOAD", log),
		uploadRequirements(
			mimeTypeService,
			settingService,
		),
	)
	r.PUT(
		"/upload",
		nostrAuthMiddleware("upload", log),
		accessControlMiddleware(acrService, "UPLOAD", log),
		uploadBlob(
			blobService,
			mimeTypeService,
			settingService,
			cdnBaseUrl,
		),
	)

	r.PUT(
		"/mirror",
		nostrAuthMiddleware("upload", log),
		accessControlMiddleware(acrService, "MIRROR", log),
		mirrorBlob(
			blobService,
			mimeTypeService,
			settingService,
			cdnBaseUrl,
		),
	)

	r.GET(
		"/list/:pubkey",
		listBlobs(blobService),
	)

	r.GET(
		"/:path",
		getBlob(blobService),
	)
	r.HEAD(
		"/:path",
		hasBlob(blobService),
	)

	r.DELETE(
		"/:path",
		nostrAuthMiddleware("delete", log),
		accessControlMiddleware(acrService, "DELETE", log),
		deleteBlob(blobService),
	)

	adminGroup := r.Group(
		"/admin",
		nostrAuthMiddleware("admin", log),
		adminMiddleware(adminPubkey),
	)
	adminGroup.GET("/rule", adminGetRules(acrService, log))
	adminGroup.POST("/rule", adminCreateRule(acrService, log))
	adminGroup.DELETE("/rule", adminDeleteRule(acrService, log))
	adminGroup.GET("/mime-type", adminGetMimeTypes(mimeTypeService, log))
	adminGroup.PUT("/mime-type", adminUpdateMimeType(mimeTypeService, log))
	adminGroup.GET("/setting", adminGetSettings(settingService))
	adminGroup.PUT("/setting", adminUpdateSetting(settingService))

	return Api{
		e:       r,
		address: apiAddress,
		log:     log,
	}
}
