package core

import (
	"context"
	"errors"
	"strconv"
)

var (
	ErrInvalidCastInt = errors.New("can't get value as int")
)

type Setting struct {
	Key   string
	Value string
}

func (s *Setting) ValueAsInt() (int, error) {
	v, err := strconv.Atoi(s.Value)
	if err != nil {
		return 0, ErrInvalidCastInt
	}
	return v, nil
}

type SettingService interface {
	AddAllowedMIMEType(
		ctx context.Context,
		mimeType string,
	) error
	DeleteAllowedMIMEType(
		ctx context.Context,
		mimeType string,
	) error
	UpdateUploadMaxSizeBytes(
		ctx context.Context,
		sizeBytes int,
	) error
}
