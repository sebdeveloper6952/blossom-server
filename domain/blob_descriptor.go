package domain

import (
	"context"
)

type BlobDescriptor struct {
	Pubkey  string
	Url     string
	Sha256  string
	Size    int64
	Type    string
	Blob    []byte
	Created int64
}

type BlobDescriptorRepo interface {
	Save(
		ctx context.Context,
		pubkey string,
		sha256 string,
		url string,
		size int64,
		mimeType string,
		blob []byte,
		created int64,
	) (*BlobDescriptor, error)
	Exists(ctx context.Context, sha256 string) (bool, error)
	GetFromHash(ctx context.Context, sha256 string) (*BlobDescriptor, error)
	GetFromPubkey(ctx context.Context, pubkey string) ([]*BlobDescriptor, error)
	DeleteFromHash(ctx context.Context, sha256 string) error
}
