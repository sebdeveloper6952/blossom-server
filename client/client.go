package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

func (c *Client) Upload(blob []byte) (*BlobDescriptor, error) {
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

	blobDescriptor := &BlobDescriptor{}
	err = json.NewDecoder(res.Body).Decode(blobDescriptor)

	return blobDescriptor, nil
}

func (c *Client) Mirror(blobUrl string) (*BlobDescriptor, error) {
	return &BlobDescriptor{}, nil
}

func (c *Client) Has(blobHash string) (bool, error) {
	req, err := http.NewRequest(http.MethodHead, c.serverUrl+"/"+blobHash, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res.StatusCode == http.StatusOK, nil
}

func (c *Client) List(pubkeyHex string) ([]BlobDescriptor, error) {
	var blobDescriptors []BlobDescriptor
	return blobDescriptors, nil
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
