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
	GetAll(
		ctx context.Context,
	) ([]*MimeType, error)
	UpdateAllowed(
		ctx context.Context,
		mimeType string,
		allowed bool,
	) error
	IsAllowed(
		ctx context.Context,
		mimeType string,
	) bool
}
