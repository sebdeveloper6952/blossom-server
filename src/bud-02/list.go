package bud02

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func ListBlobs(
	ctx context.Context,
	services core.Services,
	pubkey string,
	since int64,
	until int64,
) ([]*core.Blob, error) {
	return services.Blob().GetFromPubkeyPaginated(ctx, pubkey, since, until)
}
