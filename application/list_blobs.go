package application

import (
	"context"
	"github.com/sebdeveloper6952/blossom-server/domain"
)

func ListBlobs(
	blobDescriptorRepo domain.BlobDescriptorRepo,
) func(
	ctx context.Context,
	pubkey string,
) ([]*domain.BlobDescriptor, error) {
	return func(
		ctx context.Context,
		pubkey string,
	) ([]*domain.BlobDescriptor, error) {
		return blobDescriptorRepo.GetFromPubkey(ctx, pubkey)
	}
}
