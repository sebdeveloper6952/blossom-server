package services

import (
	"crypto/sha256"
	"fmt"
)

type Hashing interface {
	Hash(bytes []byte) (string, error)
}

type hasher struct{}

func NewSha256() (Hashing, error) {
	return &hasher{}, nil
}

func (s *hasher) Hash(bytes []byte) (string, error) {
	h := sha256.New()
	_, err := h.Write(bytes)
	hashBytes := h.Sum(nil)

	return fmt.Sprintf("%x", hashBytes), err
}
