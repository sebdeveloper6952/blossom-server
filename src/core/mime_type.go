package core

import (
	"context"
	"errors"
)

var (
	ErrMimeTypeNotAllowed = errors.New("mime/content type not allowed")
)

type MimeType struct {
	Extension string
	MimeType  string
	Allowed   bool
}

type MimeTypeService interface {
	Get(
		ctx context.Context,
		mimeType string,
	) (*MimeType, error)
	IsAllowed(
		ctx context.Context,
		mimeType string,
	) error
}
