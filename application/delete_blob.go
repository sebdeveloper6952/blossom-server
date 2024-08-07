package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/sebdeveloper6952/blossom-server/domain"
)

func DeleteBlob(
	ctx context.Context,
	blobRepo domain.BlobDescriptorRepo,
	pubkey string,
	sha256 string,
	authSha256 string,
) error {
	blobDescriptor, err := blobRepo.GetFromHash(ctx, sha256)
	if err != nil {
		return fmt.Errorf("blob not found: %w", err)
	}

	// only the owner can delete the file
	if blobDescriptor.Pubkey != pubkey {
		return errors.New("unauthorized")
	}

	// verify both hashes are the same
	if sha256 != authSha256 {
		return errors.New("unauthorized")
	}

	if err := blobRepo.DeleteFromHash(ctx, sha256); err != nil {
		return err
	}

	return nil
}
