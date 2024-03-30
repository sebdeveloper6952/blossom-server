package blob

import (
	"github.com/sebdeveloper6952/blossom-server/domain"
)

type sqliteRepo struct{}

func (s sqliteRepo) Save(sha256 string, contents []byte) (*domain.Blob, error) {
	//TODO implement me
	panic("implement me")
}

func (s sqliteRepo) GetFromHash(sha256 string) (*domain.Blob, error) {
	//TODO implement me
	panic("implement me")
}
