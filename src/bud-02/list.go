package bud02

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func ListBlobs(
	ctx context.Context,
	storage core.BlobStorage,
	pubkey string,

) ([]*core.Blob, error) {
	return storage.GetFromPubkey(ctx, pubkey)
}
