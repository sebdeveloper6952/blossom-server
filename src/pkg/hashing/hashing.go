package hashing

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

func Hash(bytes []byte) (string, error) {
	h := sha256.New()
	_, err := h.Write(bytes)
	hashBytes := h.Sum(nil)

	return fmt.Sprintf("%x", hashBytes), err
}

func IsSHA256(hash string) error {
	bytes, err := hex.DecodeString(hash)
	if err != nil {
		return err
	}

	if len(bytes) != 32 {
		return errors.New("length != 32")
	}

	return nil
}
