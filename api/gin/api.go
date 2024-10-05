package gin

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
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
	blobStorage core.BlobStorage,
	acrStorage core.ACRStorage,
	cdnBaseUrl string,
	apiAddress string,
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

	r.Use(middlewareAccessControl(acrStorage, log))

	r.LoadHTMLFiles("index.html")

	r.GET("", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.PUT(
		"/upload",
		nostrAuthMiddleware("upload", log),
		UploadBlob(blobStorage, cdnBaseUrl),
	)

	// bud-06
	r.HEAD(
		"/upload",
		UploadRequirements(),
	)

	r.PUT(
		"/mirror",
		nostrAuthMiddleware("upload", log),
		MirrorBlob(
			blobStorage,
			cdnBaseUrl,
		),
	)

	r.GET(
		"/list/:pubkey",
		ListBlobs(blobStorage),
	)

	r.GET(
		"/:path",
		GetBlob(blobStorage),
	)

	r.HEAD(
		"/:path",
		HasBlob(blobStorage),
	)

	r.DELETE(
		"/:path",
		nostrAuthMiddleware("delete", log),
		DeleteBlob(blobStorage),
	)

	return Api{
		e:       r,
		address: apiAddress,
		log:     log,
	}
}
