package bud04

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/hashing"
)

func MirrorBlob(
	ctx context.Context,
	storage core.BlobStorage,
	mimeTypeService core.MimeTypeService,
	cdnBaseUrl string,
	pubkey string,
	authHash string,
	blobUrl url.URL,
) (*core.Blob, error) {
	// if blob already exists, return BlobDescriptor from database
	if blob, err := storage.GetFromHash(ctx, authHash); err == nil {
		return blob, nil
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest(
		http.MethodGet,
		blobUrl.String(),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("GET blob at: %s: %w", blobUrl.String(), err)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET blob at: %s: %w", blobUrl.String(), err)
	}
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("GET blob at: %s returned HTTP status: %d: %w", blobUrl.String(), res.StatusCode, err)
	}
	defer func() {
		res.Body.Close()
	}()
	blobBytes, err := io.ReadAll(res.Body)
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("GET blob at: %s invalid response body: %w", blobUrl.String(), err)
	}

	mimeType := mimetype.Detect(blobBytes)
	if !mimeTypeService.IsAllowed(ctx, mimeType.String()) {
		return nil, fmt.Errorf("mime type %s not allowed", mimeType.String())
	}

	hash, err := hashing.Hash(blobBytes)
	if err != nil {
		return nil, fmt.Errorf("hash blob: %w", err)
	}

	if hash != authHash {
		return nil, fmt.Errorf("hash from auth doesn't match hash from blob")
	}

	blobDescriptor, err := storage.Save(
		ctx,
		pubkey,
		hash,
		cdnBaseUrl+"/"+hash,
		int64(len(blobBytes)),
		mimeType.String(),
		blobBytes,
		time.Now().Unix(),
	)
	if err != nil {
		return nil, fmt.Errorf("save blob: %w", err)
	}

	return blobDescriptor, nil
}
