package bud01

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func HasBlob(
	ctx context.Context,
	storage core.BlobStorage,
	hash string,
) (bool, error) {
	return storage.Exists(ctx, hash)
}
