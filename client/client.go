package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sebdeveloper6952/blossom-server/domain"
)

type Client struct {
	serverUrl string
	sk        string
	client    *http.Client
}

func New(serverUrl string, sk string) (*Client, error) {
	return &Client{
		serverUrl: serverUrl,
		sk:        sk,
		client:    &http.Client{Timeout: 10 * time.Second},
	}, nil
}

func (c *Client) Upload(blob []byte) (*domain.BlobDescriptor, error) {
	blobHash, err := hash(blob)
	if err != nil {
		return nil, err
	}

	authEventBase64, err := makeAuthEvent(blobHash, fmt.Sprintf("%d", len(blob)), "upload", c.sk)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, c.serverUrl+"/upload", bytes.NewBuffer(blob))
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

func (c *Client) Mirror(blobUrl string) (*domain.BlobDescriptor, error) {
	return &domain.BlobDescriptor{}, nil
}

func (c *Client) Has(blobHash string) (bool, error) {
	return true, nil
}

func (c *Client) List(pubkeyHex string) ([]domain.BlobDescriptor, error) {
	return nil, nil
}

func (c *Client) Get(blobHash string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.serverUrl+"/"+blobHash, http.NoBody)
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

	return io.ReadAll(res.Body)
}

func (c *Client) Delete(blobHash string) error {
	return nil
}
