package core

import (
	"context"
	"errors"
)

var (
	ErrBlobNotFound = errors.New("blob not found")
)

type Blob struct {
	Pubkey   string
	Url      string
	Sha256   string
	Size     int64
	Type     string
	Blob     []byte
	Uploaded int64
}

type BlobStorage interface {
	Save(
		ctx context.Context,
		pubkey string,
		sha256 string,
		url string,
		size int64,
		mimeType string,
		blob []byte,
		created int64,
	) (*Blob, error)
	Exists(ctx context.Context, sha256 string) (bool, error)
	GetFromHash(ctx context.Context, sha256 string) (*Blob, error)
	GetFromPubkey(ctx context.Context, pubkey string) ([]*Blob, error)
	DeleteFromHash(ctx context.Context, sha256 string) error
}
