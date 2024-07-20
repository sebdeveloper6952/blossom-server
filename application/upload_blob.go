package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/sebdeveloper6952/blossom-server/domain"
	"github.com/sebdeveloper6952/blossom-server/utils"
)

func UploadBlob(
	ctx context.Context,
	blobRepo domain.BlobDescriptorRepo,
	cdnBaseUrl string,
	authSha256 string,
	pubkey string,
	blobBytes []byte,

) (*domain.BlobDescriptor, error) {
	// TODO: here we would check if mimeType is allowed by config
	mimeType := mimetype.Detect(blobBytes)

	hash, err := utils.Hash(blobBytes)
	if err != nil {
		return nil, fmt.Errorf("hash blob: %w", err)
	}

	// calculated hash MUST match hash set in auth event
	if hash != authSha256 {
		return nil, errors.New("blob hash doesn't match auth event 'x' tag")
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
