package service

import (
	"context"
	"database/sql"

	"github.com/sebdeveloper6952/blossom-server/db"
	accesscontrol "github.com/sebdeveloper6952/blossom-server/src/access-control"
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
		database,
		queries,
		log,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	settingsService, err := NewSettingService(
		database,
		queries,
		log,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	mimeTypeService, err := NewMimeTypeService(
		queries,
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
	if err := accesscontrol.EnsureAdminHasAccess(
		ctx,
		s.acrs,
		s.conf.AdminPubkey,
	); err != nil {
		return err
	}

	return nil
}
