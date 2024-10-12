package core

import "context"

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
}
