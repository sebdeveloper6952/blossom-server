package main

import (
	"context"
)

func (s *server) HasBlob(
	ctx context.Context,
	sha256 string,
) (bool, error) {
	_, err := s.storage.Read(sha256)

	return err == nil, err
}
