package httpapi

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/sebdeveloper6952/blossom-server/v2/internal/core"
	"github.com/sebdeveloper6952/blossom-server/v2/ui"
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
		ExposeHeaders: []string{"Content-Length", HeaderXReason},
	}))

	api := r.Group("/api")
	{
		api.GET(
			"/me",
			nostrAuthMiddleware("list", log),
			getMe(adminPubkey),
		)

		admin := api.Group("/admin")
		admin.Use(
			nostrAuthMiddleware("list", log),
			adminOnlyMiddleware(adminPubkey),
		)
		{
			admin.GET("/blobs", listAllBlobs(services))
		}
	}

	r.GET("/.well-known/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

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
		deleteBlob(services, log),
	)

	// server stats
	r.GET("/stats", getStats(services))

	// admin UI (embedded SvelteKit SPA)
	if uiFS, err := ui.FS(); err == nil {
		r.GET("/admin", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/admin/")
		})
		r.GET("/admin/*filepath", serveUI(uiFS))
	} else {
		log.Warn("failed to load embedded UI: " + err.Error())
	}

	return r
}
