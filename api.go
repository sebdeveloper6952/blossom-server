package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Api struct {
	e       *gin.Engine
	address string
}

func (a *Api) Run() error {
	return a.e.Run(a.address)
}

func SetupApi(
	address string,
	server Server,
	whitelistedPks map[string]struct{},
) Api {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "PUT", "HEAD", "DELETE"},
		AllowHeaders:    []string{"Authorization"},
		ExposeHeaders:   []string{"Content-Length"},
	}))

	r.PUT(
		"/upload",
		nostrAuthMiddleware("upload"),
		whitelistPkMiddleware(whitelistedPks),
		Upload(server),
	)
	r.GET("/list/:pubkey", ListBlobs(server))
	r.GET("/:path", GetBlob(server))
	r.HEAD("/:path", HasBlob(server))
	r.DELETE("/:path", nostrAuthMiddleware("delete"), DeleteBlob(server))

	return Api{
		e:       r,
		address: address,
	}
}
