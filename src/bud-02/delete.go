package bud02

import (
	"context"
	"errors"
	"fmt"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func DeleteBlob(
	ctx context.Context,
	services core.Services,
	pubkey string,
	hash string,
	authHash string,
) error {
	var (
		blobs = services.Blob()
	)
	blobDescriptor, err := blobs.GetFromHash(ctx, hash)
	if err != nil {
		return fmt.Errorf("blob not found: %w", err)
	}

	// only the owner can delete the file
	if blobDescriptor.Pubkey != pubkey {
		return errors.New("unauthorized")
	}

	// verify both hashes are the same
	if hash != authHash {
		return errors.New("unauthorized")
	}

	if err := blobs.DeleteFromHash(ctx, hash); err != nil {
		return err
	}

	return nil
}
