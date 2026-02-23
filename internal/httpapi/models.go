package httpapi

import (
	"fmt"

	"github.com/sebdeveloper6952/blossom-server/internal/core"
)

// generic api error
type apiError struct {
	Message string `json:"message"`
}

// blobs
type blobDescriptor struct {
	Url     string     `json:"url"`
	Sha256  string     `json:"sha256"`
	Size    int64      `json:"size"`
	Type    string     `json:"type"`
	Uploaded int64     `json:"uploaded"`
	NIP94   [][]string `json:"nip94,omitempty"`
}

func fromDomainBlobDescriptor(blob *core.Blob) *blobDescriptor {
	apiBlob := &blobDescriptor{
		Url:     blob.Url,
		Sha256:  blob.Sha256,
		Size:    blob.Size,
		Type:    blob.Type,
		Uploaded: blob.Uploaded,
	}

	if blob.NIP94 != nil {
		nip94 := [][]string{
			{"url", blob.NIP94.Url},
			{"m", blob.NIP94.MimeType},
			{"x", blob.NIP94.Sha256},
			{"ox", blob.NIP94.OriginalSha256},
			{"size", fmt.Sprintf("%d", blob.Size)},
		}
		apiBlob.NIP94 = nip94
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
