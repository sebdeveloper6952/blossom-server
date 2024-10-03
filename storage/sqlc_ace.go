package storage

import (
	"context"

	"go.uber.org/zap"

	"github.com/sebdeveloper6952/blossom-server/db"
	"github.com/sebdeveloper6952/blossom-server/src/core"
)

type sqlcACRStorage struct {
	queries *db.Queries
	log     *zap.Logger
}

func NewSQLCACRStorage(
	queries *db.Queries,
	log *zap.Logger,
) (core.ACRStorage, error) {
	return &sqlcACRStorage{
		queries: queries,
		log:     log,
	}, nil
}

func (s *sqlcACRStorage) Save(
	ctx context.Context,
	action core.ACRAction,
	pubkey string,
	resource core.ACRResource,
	priority int,
) (*core.ACR, error) {
	dbACR, err := s.queries.InsertACR(
		ctx,
		db.InsertACRParams{
			Action:   string(action),
			Pubkey:   pubkey,
			Resource: string(resource),
			Priority: int64(priority),
		},
	)

	return s.dbACRInto(dbACR), err
}

func (s *sqlcACRStorage) GetFromPubkey(
	ctx context.Context,
	pubkey string,
) ([]*core.ACR, error) {
	dbACRList, err := s.queries.GetACRFromPubkey(
		ctx,
		pubkey,
	)
	if err != nil {
		return nil, err
	}

	acrList := make([]*core.ACR, 0, len(dbACRList))
	for i := range dbACRList {
		acrList = append(acrList, s.dbACRInto(dbACRList[i]))
	}

	return acrList, nil
}

func (s *sqlcACRStorage) Delete(
	ctx context.Context,
	action core.ACRAction,
	pubkey string,
	resource core.ACRResource,
) error {
	return s.queries.DeleteACR(
		ctx,
		db.DeleteACRParams{
			Action:   string(action),
			Pubkey:   pubkey,
			Resource: string(resource),
		},
	)
}

func (r *sqlcACRStorage) dbACRInto(acr db.AccessControlRule) *core.ACR {
	return &core.ACR{
		Action:   core.ACRAction(acr.Action),
		Pubkey:   acr.Pubkey,
		Resource: core.ACRResource(acr.Resource),
		Priority: int(acr.Priority),
	}
}
