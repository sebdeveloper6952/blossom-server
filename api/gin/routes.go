package gin

import (
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
)

func SetupApi(
	blobStorage core.BlobStorage,
	ac core.ACRStorage,
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
		accessControlMiddleware(ac, "UPLOAD", log),
		uploadRequirements(),
	)
	r.PUT(
		"/upload",
		nostrAuthMiddleware("upload", log),
		accessControlMiddleware(ac, "UPLOAD", log),
		uploadBlob(blobStorage, cdnBaseUrl),
	)

	r.PUT(
		"/mirror",
		nostrAuthMiddleware("upload", log),
		accessControlMiddleware(ac, "MIRROR", log),
		mirrorBlob(
			blobStorage,
			cdnBaseUrl,
		),
	)

	r.GET(
		"/list/:pubkey",
		listBlobs(blobStorage),
	)

	r.GET(
		"/:path",
		getBlob(blobStorage),
	)
	r.HEAD(
		"/:path",
		hasBlob(blobStorage),
	)

	r.DELETE(
		"/:path",
		nostrAuthMiddleware("delete", log),
		accessControlMiddleware(ac, "DELETE", log),
		deleteBlob(blobStorage),
	)

	adminGroup := r.Group(
		"/admin",
		nostrAuthMiddleware("admin", log),
		adminMiddleware(adminPubkey),
	)
	adminGroup.GET("/rule", adminGetRules(ac, log))
	adminGroup.POST("/rule", adminCreateRule(ac, log))
	adminGroup.DELETE("/rule", adminDeleteRule(ac, log))

	return Api{
		e:       r,
		address: apiAddress,
		log:     log,
	}
}
