package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
)

const (
	HeaderAuthorization  = "Authorization"
	HeaderContentType    = "Content-Type"
	HeaderXSHA256        = "X-SHA-256"
	HeaderXContentType   = "X-Content-Type"
	HeaderXContentLength = "X-Content-Length"
	HeaderXUploadMessage = "X-Upload-Message"
)

type Api struct {
	e        *gin.Engine
	address  string
	services core.Services
	log      *zap.Logger
}

func (a *Api) Run() error {
	return a.e.Run(a.address)
}
