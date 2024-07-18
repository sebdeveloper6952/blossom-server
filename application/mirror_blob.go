package application

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/sebdeveloper6952/blossom-server/domain"
	"github.com/sebdeveloper6952/blossom-server/services"
)

func MirrorBlob(
	ctx context.Context,
	blobRepo domain.BlobDescriptorRepo,
	hasher services.Hashing,
	cdnBaseUrl string,
	pubkey string,
	authSha256 string,
	blobUrl url.URL,
) (*domain.BlobDescriptor, error) {
	// if blob already exists, return BlobDescriptor from database
	if blob, err := blobRepo.GetFromHash(ctx, authSha256); err == nil {
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
	sha256, err := hasher.Hash(blobBytes)
	if err != nil {
		return nil, fmt.Errorf("hash blob: %w", err)
	}

	if sha256 != authSha256 {
		return nil, fmt.Errorf("hash from auth doesn't match hash from blob")
	}

	blobDescriptor, err := blobRepo.Save(
		ctx,
		pubkey,
		sha256,
		cdnBaseUrl+"/"+sha256,
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
