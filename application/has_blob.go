package application

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/domain"
)

func HasBlob(
	ctx context.Context,
	blobRepo domain.BlobDescriptorRepo,
	sha256 string,
) (bool, error) {
	return blobRepo.Exists(ctx, sha256)
}
