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
	storage core.BlobStorage,
	mimeTypeService core.MimeTypeService,
	cdnBaseUrl string,
	authHash string,
	pubkey string,
	blobBytes []byte,

) (*core.Blob, error) {
	mimeType := mimetype.Detect(blobBytes)
	if !mimeTypeService.IsAllowed(ctx, mimeType.String()) {
		return nil, fmt.Errorf("mime type %s not allowed", mimeType.String())
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
	if blob, err := storage.GetFromHash(ctx, hash); err == nil {
		return blob, nil
	}

	// for now the URL of the file is the URL where the CDN is being hosted
	// plus the file hash
	url := cdnBaseUrl + "/" + hash

	blobDescriptor, err := storage.Save(
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
