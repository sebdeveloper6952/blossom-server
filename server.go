package main

type Server interface {
	UploadBlob(bytes []byte) (*BlobDescriptor, error)
	GetBlob(sha256 string) ([]byte, error)
	HasBlob(sha256 string) (*BlobDescriptor, error)
}

type server struct {
	storage Storage
	hashing Hashing
}

func NewServer(
	storage Storage,
	hashing Hashing,
) (Server, error) {
	return &server{
		storage,
		hashing,
	}, nil
}
