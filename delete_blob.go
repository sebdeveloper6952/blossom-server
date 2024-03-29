package main

import (
	"context"
	"errors"
)

func (s *server) DeleteBlob(
	ctx context.Context,
	sha256 string,
	authSha256 string,
	pubkey string,
) error {
	blob, err := s.database.GetBlobFromHash(ctx, sha256)
	if err != nil {
		return err
	}

	if blob.Pubkey != pubkey {
		return errors.New("unauthorized")
	}

	if sha256 != authSha256 {
		return errors.New("unauthorized")
	}

	if err := s.storage.Delete(sha256); err != nil {
		return err
	}

	return s.database.DeleteBlobFromHash(
		ctx,
		sha256,
	)
}
