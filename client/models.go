package client

type BlobDescriptor struct {
	Pubkey  string `json:"pubkey"`
	Url     string `json:"url"`
	Sha256  string `json:"sha256"`
	Size    int64  `json:"size"`
	Type    string `json:"type"`
	Blob    []byte `json:"-"`
	Created int64  `json:"created"`
}
