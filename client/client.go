package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/nbd-wtf/go-nostr"
	"github.com/sebdeveloper6952/blossom-server/domain"
)

type Client struct {
	urls   []string
	sk     string
	client *http.Client
}

func New(urls []string, sk string) (*Client, error) {
	return &Client{
		urls:   urls,
		sk:     sk,
		client: &http.Client{Timeout: 10 * time.Second},
	}, nil
}

func makeAuthEvent(blobHash string, size string, action string, sk string) (string, error) {
	event := &nostr.Event{
		Kind:    24242,
		Content: "",
		Tags: nostr.Tags{
			nostr.Tag{"t", action},
			nostr.Tag{"x", blobHash},
			nostr.Tag{"expiration", fmt.Sprintf("%d", time.Now().Add(time.Hour).Unix())},
			nostr.Tag{"size", size},
		},
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

func (c *Client) Upload(blob []byte) (*domain.BlobDescriptor, error) {
	hash, err := hash(blob)
	if err != nil {
		return nil, err
	}

	authEventBase64, err := makeAuthEvent(hash, fmt.Sprintf("%d", len(blob)), "upload", c.sk)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, c.urls[0]+"/upload", bytes.NewBuffer(blob))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Nostr "+authEventBase64)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		res.Body.Close()
	}()

	blobDescriptor := &domain.BlobDescriptor{}
	err = json.NewDecoder(res.Body).Decode(blobDescriptor)

	return blobDescriptor, nil
}

func Get(hash string) ([]byte, error) {
	return nil, nil
}
