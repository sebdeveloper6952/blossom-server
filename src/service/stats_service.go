package service

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/db"
	"github.com/sebdeveloper6952/blossom-server/src/core"
)

type statService struct {
	queries *db.Queries
}

func NewStatService(
	queries *db.Queries,
) (core.StatService, error) {
	return &statService{
		queries,
	}, nil
}

func (s *statService) Get(
	ctx context.Context,
) (*core.Stats, error) {
	stats, err := s.queries.GetStats(ctx)
	if err != nil {
		return nil, err
	}

	return s.dbStatsIntoCore(stats), nil
}

func (s *statService) dbStatsIntoCore(stats db.GetStatsRow) *core.Stats {
	bytesStored := 0
	if stats.BytesStored.Valid {
		bytesStored = int(stats.BytesStored.Float64)
	}

	return &core.Stats{
		BytesStored: bytesStored,
		BlobCount:   int(stats.BlobCount),
		PubkeyCount: int(stats.PubkeyCount),
	}
}
