package gin

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/sebdeveloper6952/blossom-server/domain"
	"github.com/sebdeveloper6952/blossom-server/services"
)

type Api struct {
	e       *gin.Engine
	address string
	log     *zap.Logger
}

func (a *Api) Run() error {
	return a.e.Run(a.address)
}

func SetupApi(
	blobDescriptorRepo domain.BlobDescriptorRepo,
	hasher services.Hashing,
	cdnBaseUrl string,
	apiAddress string,
	whitelistedPks map[string]struct{},
	log *zap.Logger,
) Api {
	r := gin.New()

	r.Use(ginzap.Ginzap(log, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(log, true))

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "PUT", "HEAD", "DELETE"},
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
		},
		ExposeHeaders: []string{"Content-Length"},
	}))

	r.LoadHTMLFiles("index.html")

	r.GET("", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.PUT(
		"/upload",
		nostrAuthMiddleware("upload", log),
		whitelistPkMiddleware(whitelistedPks, log),
		UploadBlob(blobDescriptorRepo, hasher, cdnBaseUrl),
	)

	r.PUT(
		"/mirror",
		nostrAuthMiddleware("upload", log),
		whitelistPkMiddleware(whitelistedPks, log),
		MirrorBlob(
			blobDescriptorRepo,
			hasher,
			cdnBaseUrl,
		),
	)

	r.GET(
		"/list/:pubkey",
		ListBlobs(blobDescriptorRepo),
	)

	r.GET(
		"/:path",
		GetBlob(blobDescriptorRepo),
	)

	r.HEAD(
		"/:path",
		HasBlob(blobDescriptorRepo),
	)

	r.DELETE(
		"/:path",
		nostrAuthMiddleware("delete", log),
		DeleteBlob(blobDescriptorRepo),
	)

	return Api{
		e:       r,
		address: apiAddress,
		log:     log,
	}
}
