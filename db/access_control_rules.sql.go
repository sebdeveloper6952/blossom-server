// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: access_control_rules.sql

package db

import (
	"context"
)

const deleteACR = `-- name: DeleteACR :exec
DELETE
FROM access_control_rules
WHERE action = ? AND
      pubkey = ? AND
      resource = ?
`

type DeleteACRParams struct {
	Action   string
	Pubkey   string
	Resource string
}

func (q *Queries) DeleteACR(ctx context.Context, arg DeleteACRParams) error {
	_, err := q.db.ExecContext(ctx, deleteACR, arg.Action, arg.Pubkey, arg.Resource)
	return err
}

const getACR = `-- name: GetACR :one
SELECT "action", pubkey, resource, priority
FROM access_control_rules
WHERE action = ? AND
      pubkey = ? AND
      resource = ?
LIMIT 1
`

type GetACRParams struct {
	Action   string
	Pubkey   string
	Resource string
}

func (q *Queries) GetACR(ctx context.Context, arg GetACRParams) (AccessControlRule, error) {
	row := q.db.QueryRowContext(ctx, getACR, arg.Action, arg.Pubkey, arg.Resource)
	var i AccessControlRule
	err := row.Scan(
		&i.Action,
		&i.Pubkey,
		&i.Resource,
		&i.Priority,
	)
	return i, err
}

const getACRFromPubkey = `-- name: GetACRFromPubkey :many
SELECT "action", pubkey, resource, priority
FROM access_control_rules
WHERE pubkey = ?
`

func (q *Queries) GetACRFromPubkey(ctx context.Context, pubkey string) ([]AccessControlRule, error) {
	rows, err := q.db.QueryContext(ctx, getACRFromPubkey, pubkey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AccessControlRule
	for rows.Next() {
		var i AccessControlRule
		if err := rows.Scan(
			&i.Action,
			&i.Pubkey,
			&i.Resource,
			&i.Priority,
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

const getACRFromPubkeyResource = `-- name: GetACRFromPubkeyResource :one
SELECT "action", pubkey, resource, priority
FROM access_control_rules
WHERE pubkey = ? AND
      resource = ?
LIMIT 1
`

type GetACRFromPubkeyResourceParams struct {
	Pubkey   string
	Resource string
}

func (q *Queries) GetACRFromPubkeyResource(ctx context.Context, arg GetACRFromPubkeyResourceParams) (AccessControlRule, error) {
	row := q.db.QueryRowContext(ctx, getACRFromPubkeyResource, arg.Pubkey, arg.Resource)
	var i AccessControlRule
	err := row.Scan(
		&i.Action,
		&i.Pubkey,
		&i.Resource,
		&i.Priority,
	)
	return i, err
}

const insertACR = `-- name: InsertACR :one
INSERT INTO access_control_rules(
    action,
    pubkey,
    resource,
    priority
)
VALUES (?, ?, ?, ?)
RETURNING "action", pubkey, resource, priority
`

type InsertACRParams struct {
	Action   string
	Pubkey   string
	Resource string
	Priority int64
}

func (q *Queries) InsertACR(ctx context.Context, arg InsertACRParams) (AccessControlRule, error) {
	row := q.db.QueryRowContext(ctx, insertACR,
		arg.Action,
		arg.Pubkey,
		arg.Resource,
		arg.Priority,
	)
	var i AccessControlRule
	err := row.Scan(
		&i.Action,
		&i.Pubkey,
		&i.Resource,
		&i.Priority,
	)
	return i, err
}
