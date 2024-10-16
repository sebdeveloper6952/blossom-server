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
	Url      string `json:"url"`
	Sha256   string `json:"sha256"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
	Uploaded int64  `json:"uploaded"`
}

func fromDomainBlobDescriptor(blob *core.Blob) *blobDescriptor {
	return &blobDescriptor{
		Url:      blob.Url,
		Sha256:   blob.Sha256,
		Size:     blob.Size,
		Type:     blob.Type,
		Uploaded: blob.Uploaded,
	}
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
