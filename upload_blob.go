package main

import (
	"context"
	"github.com/sebdeveloper6952/blossom-server/db"
	"time"

	"github.com/gabriel-vasile/mimetype"
)

func (s *server) UploadBlob(
	ctx context.Context,
	pubkey string,
	bytes []byte,
) (*BlobDescriptor, error) {
	mimeType := mimetype.Detect(bytes)

	hash, err := s.hashing.Hash(bytes)
	if err != nil {
		return nil, err
	}

	url := s.cdnUrl + "/" + hash

	if err := s.storage.Save(hash, bytes); err != nil {
		return nil, err
	}

	_, err = s.database.InsertBlob(
		ctx,
		db.InsertBlobParams{
			Pubkey:  pubkey,
			Hash:    hash,
			Type:    mimeType.String(),
			Size:    int64(len(bytes)),
			Created: time.Now().Unix(),
		},
	)
	if err != nil {
		return nil, err
	}

	return &BlobDescriptor{
		Url:     url,
		Sha256:  hash,
		Type:    mimeType.String(),
		Size:    int64(len(bytes)),
		Created: time.Now().Unix(),
	}, nil
}
