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
	NIP94    *NIP94FileMetadata
}

type NIP94FileMetadata struct {
	Url            string
	MimeType       string
	Sha256         string
	OriginalSha256 string
	Size           *int64
	Dimension      *string
	Magnet         *string
	Infohash       *string
	Blurhash       *string
	ThumbnailUrl   *string
	ImageUrl       *string
	Summary        *string
	Alt            *string
	Fallback       *string
	Service        *string
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
