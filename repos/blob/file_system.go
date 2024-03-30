package blob

import (
	"context"
	"github.com/sebdeveloper6952/blossom-server/domain"
	"os"
)

type fsRepo struct {
	basePath string
}

func NewFsRepo(
	basePath string,
) (domain.BlobRepository, error) {
	return &fsRepo{
		basePath,
	}, nil
}

func (f *fsRepo) Save(ctx context.Context, sha256 string, contents []byte) (*domain.Blob, error) {
	_, err := os.Lstat(f.basePath + "/" + sha256)
	if err != nil && os.IsExist(err) {
		return nil, err
	}

	if err := os.WriteFile(
		f.basePath+"/"+sha256,
		contents,
		0644,
	); err != nil {
		return nil, err
	}

	return &domain.Blob{
		Sha256:   sha256,
		Contents: contents,
	}, nil
}

func (f *fsRepo) GetFromHash(ctx context.Context, sha256 string) (*domain.Blob, error) {
	fileBytes, err := os.ReadFile(f.basePath + "/" + sha256)
	if err != nil {
		return nil, err
	}

	return &domain.Blob{
		Sha256:   sha256,
		Contents: fileBytes,
	}, nil
}

func (f *fsRepo) DeleteFromHash(ctx context.Context, sha256 string) error {
	_, err := os.Lstat(f.basePath + "/" + sha256)
	if err != nil && os.IsExist(err) {
		return err
	}

	return os.Remove(f.basePath + "/" + sha256)
}
