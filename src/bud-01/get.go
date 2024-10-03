package bud01

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func GetBlob(
	ctx context.Context,
	storage core.BlobStorage,
	hash string,
) ([]byte, error) {
	blob, err := storage.GetFromHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	return blob.Blob, nil
}
