package core

import "context"

type Services interface {
	Init(context.Context) error
	Blob() BlobStorage
	ACR() ACRStorage
	Mime() MimeTypeService
	Settings() SettingService
	Stats() StatService
}
