package main

type BlobDescriptor struct {
	Url     string `json:"url"`
	Sha256  string `json:"sha256"`
	Size    int64  `json:"size"`
	Type    string `json:"type"`
	Created int64  `json:"created"`
}