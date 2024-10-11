package service

import (
	"context"
	"database/sql"

	"go.uber.org/zap"

	"github.com/sebdeveloper6952/blossom-server/db"
	"github.com/sebdeveloper6952/blossom-server/src/core"
)

type blobService struct {
	db         *sql.DB
	queries    *db.Queries
	cdnBaseUrl string
	log        *zap.Logger
}

func NewBlobService(
	db *sql.DB,
	queries *db.Queries,
	cdnBaseUrl string,
	log *zap.Logger,
) (core.BlobStorage, error) {
	return &blobService{
		db:         db,
		queries:    queries,
		cdnBaseUrl: cdnBaseUrl,
		log:        log,
	}, nil
}

func (r *blobService) Save(
	ctx context.Context,
	pubkey string,
	sha256 string,
	url string,
	size int64,
	mimeType string,
	blob []byte,
	created int64,
) (*core.Blob, error) {
	_, err := r.queries.InsertBlob(
		ctx,
		db.InsertBlobParams{
			Pubkey:  pubkey,
			Hash:    sha256,
			Type:    mimeType,
			Size:    size,
			Blob:    blob,
			Created: created,
		},
	)
	if err != nil {
		return nil, err
	}

	return &core.Blob{
		Url:      url,
		Sha256:   sha256,
		Size:     size,
		Type:     mimeType,
		Uploaded: created,
	}, nil
}

func (r *blobService) Exists(ctx context.Context, sha256 string) (bool, error) {
	_, err := r.queries.GetBlobFromHash(ctx, sha256)

	return err == nil, err
}

func (r *blobService) GetFromHash(ctx context.Context, sha256 string) (*core.Blob, error) {
	blob, err := r.queries.GetBlobFromHash(ctx, sha256)

	return r.dbBlobIntoBlobDescriptor(blob), err
}

func (r *blobService) GetFromPubkey(ctx context.Context, pubkey string) ([]*core.Blob, error) {
	dbBlobs, err := r.queries.GetBlobsFromPubkey(ctx, pubkey)
	if err != nil {
		return nil, err
	}

	blobs := make([]*core.Blob, len(dbBlobs))
	for i := range dbBlobs {
		blobs[i] = r.dbBlobIntoBlobDescriptor(dbBlobs[i])
	}

	return blobs, nil
}

func (r *blobService) DeleteFromHash(ctx context.Context, sha256 string) error {
	return r.queries.DeleteBlobFromHash(ctx, sha256)
}

func (r *blobService) dbBlobIntoBlobDescriptor(blob db.Blob) *core.Blob {
	return &core.Blob{
		Url:      r.cdnBaseUrl + "/" + blob.Hash,
		Sha256:   blob.Hash,
		Size:     blob.Size,
		Type:     blob.Type,
		Blob:     blob.Blob,
		Uploaded: blob.Created,
	}
}
