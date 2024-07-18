package application

import (
	"context"
	"fmt"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/sebdeveloper6952/blossom-server/domain"
	"github.com/sebdeveloper6952/blossom-server/utils"
)

func UploadBlob(
	blobRepo domain.BlobDescriptorRepo,
	cdnBaseUrl string,
) func(ctx context.Context,
	pubkey string,
	blobBytes []byte,
) (*domain.BlobDescriptor, error) {
	return func(ctx context.Context, pubkey string, blobBytes []byte) (*domain.BlobDescriptor, error) {
		mimeType := mimetype.Detect(blobBytes)

		hash, err := utils.Hash(blobBytes)
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
}
