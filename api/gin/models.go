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

// access control rules
type createACRInput struct {
	Action   string `json:"action"`
	Pubkey   string `json:"pubkey"`
	Resource string `json:"resource"`
}

type acr struct {
	Action   string `json:"action"`
	Pubkey   string `json:"pubkey"`
	Resource string `json:"resource"`
}

func fromCoreACR(rule *core.ACR) *acr {
	return &acr{
		Action:   string(rule.Action),
		Pubkey:   rule.Pubkey,
		Resource: string(rule.Resource),
	}
}

func fromSliceCoreACR(rules []*core.ACR) []*acr {
	apiRules := make([]*acr, len(rules))
	for i := range rules {
		apiRules[i] = fromCoreACR(rules[i])
	}

	return apiRules
}

// mime types
type apiMimeType struct {
	Extension string `json:"ext"`
	MimeType  string `json:"mime_type"`
	Allowed   bool   `json:"allowed"`
}

type apiUpdateMimeTypeInput struct {
	MimeType string `json:"mime_type"`
	Allowed  bool   `json:"allowed"`
}

func fromCoreMimeType(m *core.MimeType) *apiMimeType {
	return &apiMimeType{
		Extension: m.Extension,
		MimeType:  m.MimeType,
		Allowed:   m.Allowed,
	}
}

func fromSliceCoreMimeType(ms []*core.MimeType) []*apiMimeType {
	apiMimeTypes := make([]*apiMimeType, len(ms))
	for i := range ms {
		apiMimeTypes[i] = fromCoreMimeType(ms[i])
	}

	return apiMimeTypes
}

// generic api settings
type apiSetting struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type apiUpdateSettingInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func fromCoreSetting(m *core.Setting) *apiSetting {
	return &apiSetting{
		Key:   m.Key,
		Value: m.Value,
	}
}

func fromSliceCoreSetting(ms []*core.Setting) []*apiSetting {
	apiMimeTypes := make([]*apiSetting, len(ms))
	for i := range ms {
		apiMimeTypes[i] = fromCoreSetting(ms[i])
	}

	return apiMimeTypes
}
