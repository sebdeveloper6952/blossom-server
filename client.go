package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sebdeveloper6952/blossom-server/domain"
)

type Client struct {
	urls   []string
	sk     string
	client *http.Client
}

func NewClient(urls []string, sk string) (*Client, error) {
	return &Client{
		urls:   urls,
		sk:     sk,
		client: &http.Client{Timeout: 10 * time.Second},
	}, nil
}

func (c *Client) Upload(blob []byte) (*domain.BlobDescriptor, error) {
	req, err := http.NewRequest(http.MethodPut, c.urls[0], bytes.NewBuffer(blob))
	if err != nil {
		return nil, err
	}

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
