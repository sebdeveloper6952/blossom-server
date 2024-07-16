package client

import (
	"crypto/sha256"
	"fmt"
)

func hash(bytes []byte) (string, error) {
	h := sha256.New()
	_, err := h.Write(bytes)
	hashBytes := h.Sum(nil)

	return fmt.Sprintf("%x", hashBytes), err
}
