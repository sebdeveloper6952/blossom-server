package main

import (
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
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
	address string,
	server Server,
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

	r.PUT(
		"/upload",
		nostrAuthMiddleware("upload", log),
		whitelistPkMiddleware(whitelistedPks, log),
		Upload(server),
	)
	r.GET("/list/:pubkey", ListBlobs(server))
	r.GET("/:path", GetBlob(server))
	r.HEAD("/:path", HasBlob(server))
	r.DELETE("/:path", nostrAuthMiddleware("delete", log), DeleteBlob(server))

	return Api{
		e:       r,
		address: address,
		log:     log,
	}
}
