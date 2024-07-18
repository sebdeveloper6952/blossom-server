// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: blobs.sql

package db

import (
	"context"
)

const deleteBlobFromHash = `-- name: DeleteBlobFromHash :exec
delete
from blobs
where hash = ?
`

func (q *Queries) DeleteBlobFromHash(ctx context.Context, hash string) error {
	_, err := q.db.ExecContext(ctx, deleteBlobFromHash, hash)
	return err
}

const getBlobFromHash = `-- name: GetBlobFromHash :one
select pubkey, hash, type, size, blob, created
from blobs
where hash = ?
limit 1
`

func (q *Queries) GetBlobFromHash(ctx context.Context, hash string) (Blob, error) {
	row := q.db.QueryRowContext(ctx, getBlobFromHash, hash)
	var i Blob
	err := row.Scan(
		&i.Pubkey,
		&i.Hash,
		&i.Type,
		&i.Size,
		&i.Blob,
		&i.Created,
	)
	return i, err
}

const getBlobsFromPubkey = `-- name: GetBlobsFromPubkey :many
select pubkey, hash, type, size, blob, created
from blobs
where pubkey = ?
`

func (q *Queries) GetBlobsFromPubkey(ctx context.Context, pubkey string) ([]Blob, error) {
	rows, err := q.db.QueryContext(ctx, getBlobsFromPubkey, pubkey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Blob
	for rows.Next() {
		var i Blob
		if err := rows.Scan(
			&i.Pubkey,
			&i.Hash,
			&i.Type,
			&i.Size,
			&i.Blob,
			&i.Created,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertBlob = `-- name: InsertBlob :one
insert into blobs(
  pubkey,
  hash,
  type,
  size,
  blob,
  created
) values (?,?,?,?,?,?)
returning pubkey, hash, type, size, blob, created
`

type InsertBlobParams struct {
	Pubkey  string
	Hash    string
	Type    string
	Size    int64
	Blob    []byte
	Created int64
}

func (q *Queries) InsertBlob(ctx context.Context, arg InsertBlobParams) (Blob, error) {
	row := q.db.QueryRowContext(ctx, insertBlob,
		arg.Pubkey,
		arg.Hash,
		arg.Type,
		arg.Size,
		arg.Blob,
		arg.Created,
	)
	var i Blob
	err := row.Scan(
		&i.Pubkey,
		&i.Hash,
		&i.Type,
		&i.Size,
		&i.Blob,
		&i.Created,
	)
	return i, err
}
