package gin

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	accesscontrol "github.com/sebdeveloper6952/blossom-server/src/access-control"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
)

func accessControlMiddleware(
	ac core.ACRStorage,
	resource core.ACRResource,
	log *zap.Logger,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pubkey := ctx.GetString("pk")

		if err := accesscontrol.Validate(
			ctx.Request.Context(),
			ac,
			pubkey,
			resource,
			log,
		); err != nil {
			log.Debug(fmt.Sprintf("[mdwr-acl] reject: %s", err.Error()))

			code := http.StatusBadRequest
			if errors.Is(err, accesscontrol.ErrUnauthorized) {
				code = http.StatusUnauthorized
			} else if errors.Is(err, accesscontrol.ErrMissingRule) {
				code = http.StatusInternalServerError
			}

			ctx.AbortWithStatusJSON(
				code,
				apiError{
					Message: err.Error(),
				},
			)
			return
		}

		ctx.Next()
	}
}
