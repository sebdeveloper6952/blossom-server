package main

type Storage interface {
	Save(name string, bytes []byte) error
	Read(name string) ([]byte, error)
	Delete(name string) error
}
