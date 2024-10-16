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
	services core.Services,
	cdnBaseUrl string,
	adminPubkey string,
	log *zap.Logger,
) *gin.Engine {
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

	r.HEAD(
		"/upload",
		nostrAuthMiddleware("upload", log),
		uploadRequirements(services),
	)
	r.PUT(
		"/upload",
		nostrAuthMiddleware("upload", log),
		uploadBlob(
			services,
			cdnBaseUrl,
		),
	)

	r.PUT(
		"/mirror",
		nostrAuthMiddleware("upload", log),
		mirrorBlob(
			services,
			cdnBaseUrl,
		),
	)

	r.GET(
		"/list/:pubkey",
		listBlobs(services),
	)

	r.GET(
		"/:path",
		getBlob(services),
	)
	r.HEAD(
		"/:path",
		hasBlob(services),
	)

	r.DELETE(
		"/:path",
		nostrAuthMiddleware("delete", log),
		deleteBlob(services),
	)

	// admin routes
	adminGroup := r.Group(
		"/admin",
		nostrAuthMiddleware("admin", log),
		adminMiddleware(adminPubkey),
	)
	adminGroup.GET("/rule", adminGetRules(services, log))
	adminGroup.POST("/rule", adminCreateRule(services, log))
	adminGroup.DELETE("/rule", adminDeleteRule(services, log))
	adminGroup.GET("/mime-type", adminGetMimeTypes(services, log))
	adminGroup.PUT("/mime-type", adminUpdateMimeType(services, log))
	adminGroup.GET("/setting", adminGetSettings(services))
	adminGroup.PUT("/setting", adminUpdateSetting(services))

	// server stats
	r.GET("/stats", getStats(services))

	return r
}
