package core

import "context"

type Stats struct {
	BytesStored int
	BlobCount   int
	PubkeyCount int
}

type StatService interface {
	Get(ctx context.Context) (*Stats, error)
}
