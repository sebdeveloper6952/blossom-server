package application

import (
	"context"
	"github.com/gabriel-vasile/mimetype"
	"github.com/sebdeveloper6952/blossom-server/domain"
	"github.com/sebdeveloper6952/blossom-server/services"
	"time"
)

func UploadBlob(
	blobRepo domain.BlobRepository,
	blobDescriptorRepo domain.BlobDescriptorRepo,
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
			return nil, err
		}

		// if blob already exists, return BlobDescriptor from database
		if blob, err := blobDescriptorRepo.GetFromHash(ctx, hash); err == nil {
			return blob, nil
		}

		// for now the URL of the file is the URL where the CDN is being hosted
		// plus the file hash
		url := cdnBaseUrl + "/" + hash

		_, err = blobRepo.Save(ctx, hash, bytes)
		if err != nil {
			return nil, err
		}

		blobDescriptor, err := blobDescriptorRepo.Save(
			ctx,
			pubkey,
			hash,
			url,
			int64(len(bytes)),
			mimeType.String(),
			time.Now().Unix(),
		)
		if err != nil {
			return nil, err
		}

		return blobDescriptor, nil
	}
}
