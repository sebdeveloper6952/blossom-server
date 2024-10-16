package service

import (
	"context"
	"database/sql"
	"errors"

	"go.uber.org/zap"

	"github.com/sebdeveloper6952/blossom-server/db"
	"github.com/sebdeveloper6952/blossom-server/src/core"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrMissingRule  = errors.New("internal server error: missing rule")
)

type acrService struct {
	db      *sql.DB
	queries *db.Queries
	log     *zap.Logger
}

func NewACRService(
	db *sql.DB,
	queries *db.Queries,
	log *zap.Logger,
) (core.ACRStorage, error) {
	return &acrService{
		db:      db,
		queries: queries,
		log:     log,
	}, nil
}

func (s *acrService) Save(
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

func (s *acrService) SaveMany(
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

func (s *acrService) GetAll(
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

func (s *acrService) Get(
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

func (s *acrService) GetFromPubkey(
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

func (s *acrService) GetFromPubkeyResource(
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

func (s *acrService) Delete(
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

func (r *acrService) Validate(
	ctx context.Context,
	pubkey string,
	resource core.ACRResource,
) error {
	allAcr, err := r.GetFromPubkeyResource(ctx, "ALL", resource)
	if err != nil {
		// critical error: by core logic, every resource needs to have
		// an "ALL" rule
		return ErrMissingRule
	}

	pubkeyAcr, _ := r.GetFromPubkeyResource(
		ctx,
		pubkey,
		resource,
	)

	return validate(allAcr, pubkeyAcr)
}

func validate(
	allAcr *core.ACR,
	pubkeyAcr *core.ACR,
) error {
	allow := false

	if allAcr != nil {
		allow = allAcr.Action == core.ACRActionAllow
	}

	if pubkeyAcr != nil {
		allow = pubkeyAcr.Action == core.ACRActionAllow
	}

	if allow {
		return nil
	}

	return ErrUnauthorized
}

func (r *acrService) dbACRInto(acr db.AccessControlRule) *core.ACR {
	return &core.ACR{
		Action:   core.ACRAction(acr.Action),
		Pubkey:   acr.Pubkey,
		Resource: core.ACRResource(acr.Resource),
	}
}
