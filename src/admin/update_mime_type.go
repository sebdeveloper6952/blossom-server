package admin

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
)

func UpdateMimeType(
	ctx context.Context,
	services core.Services,
	mimeType string,
	allowed bool,
	log *zap.Logger,
) error {
	return services.Mime().UpdateAllowed(ctx, mimeType, allowed)
}
