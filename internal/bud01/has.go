package bud01

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/internal/core"
)

func HasBlob(
	ctx context.Context,
	services core.Services,
	hash string,
) (bool, error) {
	return services.Blob().Exists(ctx, hash)
}
