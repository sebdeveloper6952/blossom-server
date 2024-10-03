package gin

import (
	"github.com/sebdeveloper6952/blossom-server/src/core"
)

type blobDescriptor struct {
	Url      string `json:"url"`
	Sha256   string `json:"sha256"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
	Uploaded int64  `json:"uploaded"`
}

type apiError struct {
	Message string `json:"message"`
}

type mirrorInput struct {
	Url string `json:"url"`
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
