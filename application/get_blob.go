package application

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/domain"
)

func GetBlob(
	blobRepo domain.BlobDescriptorRepo,
) func(
	ctx context.Context,
	sha256 string,
) ([]byte, error) {
	return func(ctx context.Context, sha256 string) ([]byte, error) {
		blob, err := blobRepo.GetFromHash(ctx, sha256)
		if err != nil {
			return nil, err
		}

		return blob.Blob, nil
	}
}
