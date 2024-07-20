package blob_descriptor

import (
	"context"

	"go.uber.org/zap"

	"github.com/sebdeveloper6952/blossom-server/db"
	"github.com/sebdeveloper6952/blossom-server/domain"
)

type sqlcRepo struct {
	queries    *db.Queries
	cdnBaseUrl string
	l          *zap.Logger
}

func NewSqlcRepo(
	queries *db.Queries,
	cdnBaseUrl string,
	log *zap.Logger,
) (domain.BlobDescriptorRepo, error) {
	return &sqlcRepo{
		queries:    queries,
		cdnBaseUrl: cdnBaseUrl,
		l:          log,
	}, nil
}

func (r *sqlcRepo) Save(
	ctx context.Context,
	pubkey string,
	sha256 string,
	url string,
	size int64,
	mimeType string,
	blob []byte,
	created int64,
) (*domain.BlobDescriptor, error) {
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

	return &domain.BlobDescriptor{
		Url:      url,
		Sha256:   sha256,
		Size:     size,
		Type:     mimeType,
		Uploaded: created,
	}, nil
}

func (r *sqlcRepo) Exists(ctx context.Context, sha256 string) (bool, error) {
	_, err := r.queries.GetBlobFromHash(ctx, sha256)

	return err == nil, err
}

func (r *sqlcRepo) GetFromHash(ctx context.Context, sha256 string) (*domain.BlobDescriptor, error) {
	blob, err := r.queries.GetBlobFromHash(ctx, sha256)

	return r.dbBlobIntoBlobDescriptor(blob), err
}

func (r *sqlcRepo) GetFromPubkey(ctx context.Context, pubkey string) ([]*domain.BlobDescriptor, error) {
	dbBlobs, err := r.queries.GetBlobsFromPubkey(ctx, pubkey)
	if err != nil {
		return nil, err
	}

	blobs := make([]*domain.BlobDescriptor, len(dbBlobs))
	for i := range dbBlobs {
		blobs[i] = r.dbBlobIntoBlobDescriptor(dbBlobs[i])
	}

	return blobs, nil
}

func (r *sqlcRepo) DeleteFromHash(ctx context.Context, sha256 string) error {
	return r.queries.DeleteBlobFromHash(ctx, sha256)
}

func (r *sqlcRepo) dbBlobIntoBlobDescriptor(blob db.Blob) *domain.BlobDescriptor {
	return &domain.BlobDescriptor{
		Url:      r.cdnBaseUrl + "/" + blob.Hash,
		Sha256:   blob.Hash,
		Size:     blob.Size,
		Type:     blob.Type,
		Blob:     blob.Blob,
		Uploaded: blob.Created,
	}
}
