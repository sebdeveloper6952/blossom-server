package bud02

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/sebdeveloper6952/blossom-server/internal/core"
)

func DeleteBlob(
	ctx context.Context,
	services core.Services,
	pubkey string,
	hash string,
	authHash string,
	log *zap.Logger,
) error {
	var (
		blobs = services.Blob()
	)

	log.Debug("delete blob request", zap.String("pubkey", pubkey), zap.String("hash", hash))

	blobDescriptor, err := blobs.GetFromHash(ctx, hash)
	if err != nil {
		log.Debug("blob not found", zap.String("hash", hash), zap.Error(err))
		return core.ErrBlobNotFound
	}

	// only the owner can delete the file
	if blobDescriptor.Pubkey != pubkey {
		log.Debug("delete unauthorized: pubkey mismatch",
			zap.String("owner", blobDescriptor.Pubkey),
			zap.String("requester", pubkey),
		)
		return errors.New("unauthorized")
	}

	// verify both hashes are the same
	if hash != authHash {
		log.Debug("delete unauthorized: hash mismatch",
			zap.String("hash", hash),
			zap.String("authHash", authHash),
		)
		return errors.New("unauthorized")
	}

	if err := blobs.DeleteFromHash(ctx, hash); err != nil {
		log.Error("failed to delete blob", zap.String("hash", hash), zap.Error(err))
		return err
	}

	log.Debug("blob deleted", zap.String("hash", hash))
	return nil
}
