package bud02

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func ListBlobs(
	ctx context.Context,
	services core.Services,
	pubkey string,
) ([]*core.Blob, error) {
	return services.Blob().GetFromPubkey(ctx, pubkey)
}
