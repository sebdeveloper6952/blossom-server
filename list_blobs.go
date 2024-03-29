package main

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/db"
)

func (s *server) ListBlobs(
	ctx context.Context,
	pubkey string,
) ([]BlobDescriptor, error) {
	dbBlobs, err := s.database.GetBlobsFromPubkey(ctx, pubkey)
	if err != nil {
		return nil, err
	}

	blobs := make([]BlobDescriptor, 0, len(dbBlobs))
	for i := range dbBlobs {
		blobs = append(blobs, dbBlobIntoBlobDescriptor(dbBlobs[i]))
	}

	return blobs, nil
}

func dbBlobIntoBlobDescriptor(b db.Blob) BlobDescriptor {
	return BlobDescriptor{
		Url:     "http://127.0.0.1:8000/" + b.Hash,
		Sha256:  b.Hash,
		Size:    b.Size,
		Type:    b.Type,
		Created: b.Created,
	}
}
