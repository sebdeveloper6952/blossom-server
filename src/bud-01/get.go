package bud01

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func GetBlob(
	ctx context.Context,
	services core.Services,
	hash string,
) (*core.Blob, error) {
	blob, err := services.Blob().GetFromHash(ctx, hash)
	if err != nil {
		return nil, core.ErrBlobNotFound
	}

	return blob, nil
}
