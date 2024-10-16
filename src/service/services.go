package service

import (
	"context"
	"database/sql"

	"github.com/sebdeveloper6952/blossom-server/db"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/config"
	"go.uber.org/zap"
)

type services struct {
	blobs    core.BlobStorage
	acrs     core.ACRStorage
	mimes    core.MimeTypeService
	settings core.SettingService
	stats    core.StatService
	conf     *config.Config
}

func New(
	ctx context.Context,
	database *sql.DB,
	queries *db.Queries,
	conf *config.Config,
	log *zap.Logger,
) core.Services {
	blobService, err := NewBlobService(
		database,
		queries,
		conf.CdnUrl,
		log,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	acrService, err := NewACRService(
		conf,
		log,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	settingsService, err := NewSettingService(
		conf.MaxUploadSizeBytes,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	mimeTypeService, err := NewMimeTypeService(
		ctx,
		queries,
		conf,
		log,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	statService, err := NewStatService(queries)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &services{
		blobs:    blobService,
		acrs:     acrService,
		mimes:    mimeTypeService,
		settings: settingsService,
		stats:    statService,
		conf:     conf,
	}
}

func (s *services) Blob() core.BlobStorage {
	return s.blobs
}

func (s *services) ACR() core.ACRStorage {
	return s.acrs
}

func (s *services) Mime() core.MimeTypeService {
	return s.mimes
}

func (s *services) Settings() core.SettingService {
	return s.settings
}

func (s *services) Stats() core.StatService {
	return s.stats
}

func (s *services) Init(ctx context.Context) error {
	return nil
}
