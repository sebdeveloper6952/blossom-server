// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

type Blob struct {
	Pubkey  string
	Hash    string
	Type    string
	Size    int64
	Blob    []byte
	Created int64
}

type MimeType struct {
	Extension string
	MimeType  string
}
