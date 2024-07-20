package application

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/domain"
)

func ListBlobs(
	ctx context.Context,
	blobDescriptorRepo domain.BlobDescriptorRepo,
	pubkey string,

) ([]*domain.BlobDescriptor, error) {
	return blobDescriptorRepo.GetFromPubkey(ctx, pubkey)
}
