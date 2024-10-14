package bud06

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func UploadRequirements(
	ctx context.Context,
	mimeTypeService core.MimeTypeService,
	settingService core.SettingService,
	blobHash string,
	contentType string,
	contentLength int,
) error {
	if err := mimeTypeService.IsAllowed(ctx, contentType); err != nil {
		return err
	}

	if err := settingService.ValidateFileSizeMaxBytes(ctx, contentLength); err != nil {
		return err
	}

	return nil
}
