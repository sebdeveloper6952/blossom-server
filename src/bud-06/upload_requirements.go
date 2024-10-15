package bud06

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func UploadRequirements(
	ctx context.Context,
	services core.Services,
	blobHash string,
	contentType string,
	contentLength int,
) error {
	if err := services.Mime().IsAllowed(ctx, contentType); err != nil {
		return err
	}

	if err := services.Settings().ValidateFileSizeMaxBytes(ctx, contentLength); err != nil {
		return err
	}

	return nil
}
