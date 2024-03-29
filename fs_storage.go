package main

import "os"

type fsStorage struct {
	basePath string
}

func NewFsStorage(basePath string) (Storage, error) {
	return &fsStorage{
		basePath,
	}, nil
}

func (f *fsStorage) Save(name string, bytes []byte) error {
	_, err := os.Lstat(f.basePath + "/" + name)
	if err != nil && os.IsExist(err) {
		return err
	}

	return os.WriteFile(
		f.basePath+"/"+name,
		bytes,
		0644,
	)
}

func (f *fsStorage) Read(name string) ([]byte, error) {
	return os.ReadFile(f.basePath + "/" + name)
}

func (f *fsStorage) Delete(name string) error {
	_, err := os.Lstat(f.basePath + "/" + name)
	if err != nil && os.IsExist(err) {
		return err
	}

	return os.Remove(f.basePath + "/" + name)
}
