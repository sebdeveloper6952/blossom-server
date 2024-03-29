package main

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"time"
)

func (s *server) UploadBlob(bytes []byte) (*BlobDescriptor, error) {
	mtype := mimetype.Detect(bytes)

	hash, err := s.hashing.Hash(bytes)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://127.0.0.1:8000/%s", hash)

	if err := s.storage.Save(hash, bytes); err != nil {
		return nil, err
	}

	return &BlobDescriptor{
		Url:     url,
		Sha256:  hash,
		Type:    mtype.String(),
		Size:    len(bytes),
		Created: time.Now().Unix(),
	}, nil
}
