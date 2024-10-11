package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sebdeveloper6952/blossom-server/db"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
)

const (
	keyAllowedMIMEType    = "ALLOWED_MIME_TYPE"
	keyUploadMaxSizeBytes = "UPLOAD_MAX_SIZE_BYTES"
)

type settingService struct {
	db      *sql.DB
	queries *db.Queries
	log     *zap.Logger
}

func NewSettingService(
	db *sql.DB,
	queries *db.Queries,
	log *zap.Logger,
) (core.SettingService, error) {
	return &settingService{
		db,
		queries,
		log,
	}, nil
}

func (s *settingService) AddAllowedMIMEType(
	ctx context.Context,
	mimeType string,
) error {
	_, err := s.queries.InsertSetting(
		ctx,
		db.InsertSettingParams{
			Key:   keyAllowedMIMEType,
			Value: mimeType,
		},
	)

	return err
}

func (s *settingService) DeleteAllowedMIMEType(
	ctx context.Context,
	mimeType string,
) error {
	err := s.queries.DeleteSetting(
		ctx,
		db.DeleteSettingParams{
			Key:   keyAllowedMIMEType,
			Value: mimeType,
		},
	)
	return err
}

func (s *settingService) UpdateUploadMaxSizeBytes(
	ctx context.Context,
	sizeBytes int,
) error {
	_, err := s.queries.UpdateSetting(
		ctx,
		db.UpdateSettingParams{
			Key:   keyUploadMaxSizeBytes,
			Value: fmt.Sprintf("%d", sizeBytes),
		},
	)

	return err
}
