package bud06

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func UploadRequirements(
	ctx context.Context,
	mimeTypeService core.MimeTypeService,
	blobHash string,
	contentType string,
	contentLength int,
) error {
	if !mimeTypeService.IsAllowed(ctx, contentType) {
		return core.ErrMimeTypeNotAllowed
	}

	return nil
}
