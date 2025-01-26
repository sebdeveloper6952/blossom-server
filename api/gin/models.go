package gin

import (
	"github.com/sebdeveloper6952/blossom-server/src/core"
)

// generic api error
type apiError struct {
	Message string `json:"message"`
}

// blobs
type blobDescriptor struct {
	Url               string             `json:"url"`
	Sha256            string             `json:"sha256"`
	Size              int64              `json:"size"`
	Type              string             `json:"type"`
	Uploaded          int64              `json:"uploaded"`
	NIP94FileMetadata *nip94FileMetadata `json:"nip94,omitempty"`
}

// NIP94 https://github.com/nostr-protocol/nips/blob/master/94.md
type nip94FileMetadata struct {
	Url            string  `json:"url"`
	MimeType       string  `json:"m"`
	Sha256         string  `json:"x"`
	OriginalSha256 string  `json:"ox"`
	Size           *int64  `json:"size,omitempty"`
	Dimension      *string `json:"dim,omitempty"`
	Magnet         *string `json:"magnet,omitempty"`
	Infohash       *string `json:"i,omitempty"`
	Blurhash       *string `json:"blurhash,omitempty"`
	ThumbnailUrl   *string `json:"thumb,omitempty"`
	ImageUrl       *string `json:"image,omitempty"`
	Summary        *string `json:"summary,omitempty"`
	Alt            *string `json:"alt,omitempty"`
	Fallback       *string `json:"fallback,omitempty"`
	Service        *string `json:"service,omitempty"`
}

func fromDomainBlobDescriptor(blob *core.Blob) *blobDescriptor {
	apiBlob := &blobDescriptor{
		Url:      blob.Url,
		Sha256:   blob.Sha256,
		Size:     blob.Size,
		Type:     blob.Type,
		Uploaded: blob.Uploaded,
	}

	if blob.NIP94 != nil {
		apiBlob.NIP94FileMetadata = &nip94FileMetadata{
			Url:            blob.NIP94.Url,
			MimeType:       blob.NIP94.MimeType,
			Sha256:         blob.NIP94.Sha256,
			OriginalSha256: blob.NIP94.OriginalSha256,
		}
	}

	return apiBlob
}

func fromSliceDomainBlobDescriptor(blobs []*core.Blob) []*blobDescriptor {
	apiBlobs := make([]*blobDescriptor, len(blobs))
	for i := range blobs {
		apiBlobs[i] = fromDomainBlobDescriptor(blobs[i])
	}

	return apiBlobs
}

// bud-04 mirror a blob
type mirrorInput struct {
	Url string `json:"url"`
}
