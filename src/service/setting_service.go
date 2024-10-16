package service

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

const (
	keyAllowedMIMEType    = "ALLOWED_MIME_TYPE"
	keyUploadMaxSizeBytes = "UPLOAD_MAX_SIZE_BYTES"
)

type settingService struct {
	maxUploadSizeBytes int
}

func NewSettingService(
	maxUploadSizeBytes int,
) (core.SettingService, error) {
	return &settingService{
		maxUploadSizeBytes,
	}, nil
}

func (s *settingService) ValidateFileSizeMaxBytes(
	ctx context.Context,
	sizeBytes int,
) error {
	if sizeBytes > s.maxUploadSizeBytes {
		return core.ErrFileSizeLimit
	}

	return nil
}
