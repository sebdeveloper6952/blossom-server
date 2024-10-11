// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: settings.sql

package db

import (
	"context"
)

const deleteSetting = `-- name: DeleteSetting :exec
;

DELETE FROM settings
WHERE key = ? AND value = ?
`

type DeleteSettingParams struct {
	Key   string
	Value string
}

func (q *Queries) DeleteSetting(ctx context.Context, arg DeleteSettingParams) error {
	_, err := q.db.ExecContext(ctx, deleteSetting, arg.Key, arg.Value)
	return err
}

const insertSetting = `-- name: InsertSetting :one
INSERT INTO settings(key, value)
VALUES (?, ?)
ON CONFLICT (key, value)
DO NOTHING
RETURNING "key", value
`

type InsertSettingParams struct {
	Key   string
	Value string
}

func (q *Queries) InsertSetting(ctx context.Context, arg InsertSettingParams) (Setting, error) {
	row := q.db.QueryRowContext(ctx, insertSetting, arg.Key, arg.Value)
	var i Setting
	err := row.Scan(&i.Key, &i.Value)
	return i, err
}

const updateSetting = `-- name: UpdateSetting :one
UPDATE settings
SET value = ?
WHERE key = ?
RETURNING "key", value
`

type UpdateSettingParams struct {
	Value string
	Key   string
}

func (q *Queries) UpdateSetting(ctx context.Context, arg UpdateSettingParams) (Setting, error) {
	row := q.db.QueryRowContext(ctx, updateSetting, arg.Value, arg.Key)
	var i Setting
	err := row.Scan(&i.Key, &i.Value)
	return i, err
}
