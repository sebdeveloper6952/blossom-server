package storage

import (
	"context"
	"database/sql"

	"go.uber.org/zap"

	"github.com/sebdeveloper6952/blossom-server/db"
	"github.com/sebdeveloper6952/blossom-server/src/core"
)

type sqlcACRStorage struct {
	db      *sql.DB
	queries *db.Queries
	log     *zap.Logger
}

func NewSQLCACRStorage(
	db *sql.DB,
	queries *db.Queries,
	log *zap.Logger,
) (core.ACRStorage, error) {
	return &sqlcACRStorage{
		db:      db,
		queries: queries,
		log:     log,
	}, nil
}

func (s *sqlcACRStorage) Save(
	ctx context.Context,
	action core.ACRAction,
	pubkey string,
	resource core.ACRResource,
) (*core.ACR, error) {
	dbACR, err := s.queries.InsertACR(
		ctx,
		db.InsertACRParams{
			Action:   string(action),
			Pubkey:   pubkey,
			Resource: string(resource),
		},
	)

	return s.dbACRInto(dbACR), err
}

func (s *sqlcACRStorage) SaveMany(
	ctx context.Context,
	rules []*core.ACR,
) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()
	txQ := s.queries.WithTx(tx)

	for _, rule := range rules {
		_, err := txQ.InsertACR(
			ctx,
			db.InsertACRParams{
				Action:   string(rule.Action),
				Pubkey:   rule.Pubkey,
				Resource: string(rule.Resource),
			},
		)
		if err != nil {
			return err
		}
	}
	tx.Commit()

	return nil
}

func (s *sqlcACRStorage) GetAll(
	ctx context.Context,
) ([]*core.ACR, error) {
	dbACRList, err := s.queries.GetAllACR(
		ctx,
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

func (s *sqlcACRStorage) Get(
	ctx context.Context,
	action core.ACRAction,
	pubkey string,
	resource core.ACRResource,
) (*core.ACR, error) {
	acr, err := s.queries.GetACR(
		ctx,
		db.GetACRParams{
			Action:   string(action),
			Pubkey:   pubkey,
			Resource: string(resource),
		},
	)

	return s.dbACRInto(acr), err
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

func (s *sqlcACRStorage) GetFromPubkeyResource(
	ctx context.Context,
	pubkey string,
	resource core.ACRResource,
) (*core.ACR, error) {
	acr, err := s.queries.GetACRFromPubkeyResource(
		ctx,
		db.GetACRFromPubkeyResourceParams{
			Pubkey:   pubkey,
			Resource: string(resource),
		},
	)

	return s.dbACRInto(acr), err
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
	}
}
