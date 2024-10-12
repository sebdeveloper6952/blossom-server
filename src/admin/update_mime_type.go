package admin

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
)

func UpdateMimeType(
	ctx context.Context,
	mimeTypeService core.MimeTypeService,
	mimeType string,
	allowed bool,
	log *zap.Logger,
) error {
	// TODO: mime type exists?
	return mimeTypeService.UpdateAllowed(ctx, mimeType, allowed)
}
