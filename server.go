package main

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/db"
)

// Server holds the application logic

type Server interface {
	UploadBlob(
		ctx context.Context,
		pubkey string,
		bytes []byte,
	) (*BlobDescriptor, error)
	GetBlob(
		ctx context.Context,
		sha256 string,
	) ([]byte, error)
	HasBlob(
		ctx context.Context,
		sha256 string,
	) (bool, error)
	ListBlobs(
		ctx context.Context,
		pubkey string,
	) ([]BlobDescriptor, error)
	DeleteBlob(
		ctx context.Context,
		sha256 string,
		authSha256 string,
		pubkey string,
	) error
}

type server struct {
	database *db.Queries
	storage  Storage
	hashing  Hashing
}

func NewServer(
	database *db.Queries,
	storage Storage,
	hashing Hashing,
) (Server, error) {
	return &server{
		database,
		storage,
		hashing,
	}, nil
}
