package client

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nbd-wtf/go-nostr"
)

func makeAuthEvent(blobHash string, size string, action string, sk string) (string, error) {
	event := &nostr.Event{
		Kind:    24242,
		Content: "",
		Tags: nostr.Tags{
			nostr.Tag{"t", action},
			nostr.Tag{"x", blobHash},
			nostr.Tag{"expiration", fmt.Sprintf("%d", time.Now().Add(time.Hour).Unix())},
		},
	}

	if size != "" {
		event.Tags = append(event.Tags, nostr.Tag{"size", size})
	}

	if err := event.Sign(sk); err != nil {
		return "", err
	}

	eventJson, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(eventJson), nil
}

func hash(bytes []byte) (string, error) {
	h := sha256.New()
	_, err := h.Write(bytes)
	hashBytes := h.Sum(nil)

	return fmt.Sprintf("%x", hashBytes), err
}
