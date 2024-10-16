package service

import (
	"context"
	"fmt"

	"github.com/sebdeveloper6952/blossom-server/db"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/config"
	"go.uber.org/zap"
)

type mimeTypeService struct {
	allowed map[string]struct{}
	queries *db.Queries
	conf    *config.Config
	log     *zap.Logger
}

func NewMimeTypeService(
	ctx context.Context,
	queries *db.Queries,
	conf *config.Config,
	log *zap.Logger,
) (core.MimeTypeService, error) {
	allowed := make(map[string]struct{})
	for _, mime := range conf.AllowedMimeTypes {
		_, err := queries.GetMimeType(ctx, mime)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", mime, core.ErrInvalidMimeType)
		}
		allowed[mime] = struct{}{}
	}

	return &mimeTypeService{
		allowed: allowed,
		queries: queries,
		conf:    conf,
		log:     log,
	}, nil
}

func (s *mimeTypeService) Get(
	ctx context.Context,
	mimeType string,
) (*core.MimeType, error) {
	dbMimeType, err := s.queries.GetMimeType(ctx, mimeType)

	return s.dbMimeTypeIntoCore(dbMimeType), err
}

func (s *mimeTypeService) IsAllowed(
	ctx context.Context,
	mimeType string,
) error {
	_, ok := s.allowed[mimeType]
	if !ok {
		return core.ErrMimeTypeNotAllowed
	}

	return nil
}

func (s *mimeTypeService) dbMimeTypeIntoCore(m db.MimeType) *core.MimeType {
	return &core.MimeType{
		Extension: m.Extension,
		MimeType:  m.MimeType,
		Allowed:   dbBoolToBool(m.Allowed),
	}
}

// TODO: create pkg for sqlite utils
func dbBoolToBool(v int64) bool {
	return v == 1
}

func boolToDbBool(v bool) int64 {
	if v {
		return 1
	}
	return 0
}
