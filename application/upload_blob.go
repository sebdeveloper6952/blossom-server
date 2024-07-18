package application

import (
	"context"
	"fmt"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/sebdeveloper6952/blossom-server/domain"
	"github.com/sebdeveloper6952/blossom-server/services"
)

func UploadBlob(
	blobRepo domain.BlobDescriptorRepo,
	hasher services.Hashing,
	cdnBaseUrl string,
) func(ctx context.Context,
	pubkey string,
	bytes []byte,
) (*domain.BlobDescriptor, error) {
	return func(ctx context.Context, pubkey string, bytes []byte) (*domain.BlobDescriptor, error) {
		mimeType := mimetype.Detect(bytes)

		hash, err := hasher.Hash(bytes)
		if err != nil {
			return nil, fmt.Errorf("hash blob: %w", err)
		}

		// if blob already exists, return BlobDescriptor from database
		if blob, err := blobRepo.GetFromHash(ctx, hash); err == nil {
			return blob, nil
		}

		// for now the URL of the file is the URL where the CDN is being hosted
		// plus the file hash
		url := cdnBaseUrl + "/" + hash

		blobDescriptor, err := blobRepo.Save(
			ctx,
			pubkey,
			hash,
			url,
			int64(len(bytes)),
			mimeType.String(),
			bytes,
			time.Now().Unix(),
		)
		if err != nil {
			return nil, fmt.Errorf("write to database: %w", err)
		}

		return blobDescriptor, nil
	}
}
