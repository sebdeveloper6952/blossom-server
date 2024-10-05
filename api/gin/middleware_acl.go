package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
)

func middlewareAccessControl(
	_ core.ACRStorage,
	log *zap.Logger,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Debug(fmt.Sprintf("[mdwr-acl] %s", ctx.Request.RequestURI))
		ctx.Next()
	}
}
