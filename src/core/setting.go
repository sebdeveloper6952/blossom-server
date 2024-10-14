package core

import (
	"context"
	"errors"
	"strconv"
)

var (
	ErrInvalidCastInt = errors.New("can't get value as int")
	ErrFileSizeLimit  = errors.New("file size is greater than allowed")
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
	Get(
		ctx context.Context,
		key string,
	) (*Setting, error)
	GetAll(
		ctx context.Context,
	) ([]*Setting, error)
	Update(
		ctx context.Context,
		key string,
		value string,
	) (*Setting, error)
	ValidateFileSizeMaxBytes(
		ctx context.Context,
		sizeBytes int,
	) error
}
