package application

import (
	"context"
	"github.com/sebdeveloper6952/blossom-server/domain"
)

func HasBlob(
	blobDescriptorRepo domain.BlobDescriptorRepo,
) func(
	ctx context.Context,
	sha256 string,
) (bool, error) {
	return func(
		ctx context.Context,
		sha256 string,
	) (bool, error) {
		return blobDescriptorRepo.Exists(ctx, sha256)
	}
}
