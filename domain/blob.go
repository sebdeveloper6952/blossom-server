package domain

import (
	"context"
)

type Blob struct {
	Sha256   string
	Contents []byte
}

type BlobRepository interface {
	Save(ctx context.Context, sha256 string, contents []byte) (*Blob, error)
	GetFromHash(ctx context.Context, sha256 string) (*Blob, error)
	DeleteFromHash(ctx context.Context, sha256 string) error
}
