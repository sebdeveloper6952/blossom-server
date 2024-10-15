package admin

import (
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func GetMimeTypes(
	ctx context.Context,
	services core.Services,
	log *zap.Logger,
) ([]*core.MimeType, error) {
	return services.Mime().GetAll(ctx)
}
