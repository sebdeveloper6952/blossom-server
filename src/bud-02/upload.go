package bud02

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/hashing"
)

func UploadBlob(
	ctx context.Context,
	srv core.Services,
	cdnBaseUrl string,
	authHash string,
	pubkey string,
	blobBytes []byte,
) (*core.Blob, error) {
	var (
		blobs    = srv.Blob()
		mimes    = srv.Mime()
		settings = srv.Settings()
	)

	mimeType := mimetype.Detect(blobBytes)
	if err := mimes.IsAllowed(ctx, mimeType.String()); err != nil {
		return nil, fmt.Errorf("mime type %s not allowed", mimeType.String())
	}

	if err := settings.ValidateFileSizeMaxBytes(ctx, len(blobBytes)); err != nil {
		return nil, err
	}

	hash, err := hashing.Hash(blobBytes)
	if err != nil {
		return nil, fmt.Errorf("hash blob: %w", err)
	}

	// calculated hash MUST match hash set in auth event
	if hash != authHash {
		return nil, errors.New("blob hash doesn't match auth event 'x' tag")
	}

	// if blob already exists, return BlobDescriptor from database
	if blob, err := blobs.GetFromHash(ctx, hash); err == nil {
		return blob, nil
	}

	// for now the URL of the file is the URL where the CDN is being hosted
	// plus the file hash
	url := cdnBaseUrl + "/" + hash

	blobDescriptor, err := blobs.Save(
		ctx,
		pubkey,
		hash,
		url,
		int64(len(blobBytes)),
		mimeType.String(),
		blobBytes,
		time.Now().Unix(),
	)
	if err != nil {
		return nil, fmt.Errorf("save blob: %w", err)
	}

	return blobDescriptor, nil
}
