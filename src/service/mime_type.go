package service

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/db"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
)

type mimeTypeService struct {
	queries *db.Queries
	log     *zap.Logger
}

func NewMimeTypeService(
	queries *db.Queries,
	log *zap.Logger,
) (core.MimeTypeService, error) {
	return &mimeTypeService{
		queries,
		log,
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
	dbMimeType, err := s.queries.GetMimeType(ctx, mimeType)
	if err != nil {
		return err
	}

	if !dbBoolToBool(dbMimeType.Allowed) {
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
