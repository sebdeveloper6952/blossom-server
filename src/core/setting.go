package core

import (
	"context"
	"errors"
)

var (
	ErrFileSizeLimit = errors.New("file size is greater than allowed")
)

type Setting struct {
	Key   string
	Value string
}

type SettingService interface {
	ValidateFileSizeMaxBytes(
		ctx context.Context,
		sizeBytes int,
	) error
}
