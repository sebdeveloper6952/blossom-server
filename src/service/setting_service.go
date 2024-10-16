package service

import (
	"context"
	"database/sql"
	"strconv"

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

func (s *settingService) Get(
	ctx context.Context,
	key string,
) (*core.Setting, error) {
	dbSetting, err := s.queries.GetSetting(ctx, key)
	if err != nil {
		return nil, err
	}

	return s.dbSettingIntoCore(dbSetting), err
}

func (s *settingService) GetAll(
	ctx context.Context,
) ([]*core.Setting, error) {
	dbSettings, err := s.queries.GetAllSettings(ctx)
	if err != nil {
		return nil, err
	}

	settings := make([]*core.Setting, len(dbSettings))
	for i := range dbSettings {
		settings[i] = s.dbSettingIntoCore(dbSettings[i])
	}

	return settings, err
}

func (s *settingService) Update(
	ctx context.Context,
	key string,
	value string,
) (*core.Setting, error) {
	dbSetting, err := s.queries.UpdateSetting(
		ctx,
		db.UpdateSettingParams{
			Key:   keyUploadMaxSizeBytes,
			Value: value,
		},
	)

	return s.dbSettingIntoCore(dbSetting), err
}

func (s *settingService) ValidateFileSizeMaxBytes(
	ctx context.Context,
	sizeBytes int,
) error {
	setting, err := s.queries.GetSetting(ctx, keyUploadMaxSizeBytes)
	if err != nil {
		return err
	}

	settingValueInt, err := strconv.Atoi(setting.Value)
	if err != nil {
		return core.ErrInvalidCastInt
	}

	if sizeBytes > settingValueInt {
		return core.ErrFileSizeLimit
	}

	return nil
}

func (s *settingService) dbSettingIntoCore(m db.Setting) *core.Setting {
	return &core.Setting{
		Key:   m.Key,
		Value: m.Value,
	}
}
