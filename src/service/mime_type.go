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

func (s *mimeTypeService) GetAll(
	ctx context.Context,
) ([]*core.MimeType, error) {
	dbMimeTypes, err := s.queries.GetAllMimeTypes(ctx)
	if err != nil {
		return nil, err
	}

	mimeTypes := make([]*core.MimeType, len(dbMimeTypes))
	for i := range dbMimeTypes {
		mimeTypes[i] = s.dbMimeTypeIntoCore(dbMimeTypes[i])
	}

	return mimeTypes, nil
}

func (s *mimeTypeService) UpdateAllowed(
	ctx context.Context,
	mimeType string,
	allowed bool,
) error {
	_, err := s.queries.UpdateMimeType(
		ctx,
		db.UpdateMimeTypeParams{
			MimeType: mimeType,
			Allowed:  boolToDbBool(allowed),
		},
	)

	return err
}

func (s *mimeTypeService) IsAllowed(
	ctx context.Context,
	mimeType string,
) bool {
	dbMimeType, err := s.queries.GetMimeType(ctx, mimeType)
	if err != nil {
		return false
	}

	return dbBoolToBool(dbMimeType.Allowed)
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
