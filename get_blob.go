package main

import (
	"context"
)

func (s *server) GetBlob(
	_ context.Context,
	sha256 string,
) ([]byte, error) {
	return s.storage.Read(sha256)
}
