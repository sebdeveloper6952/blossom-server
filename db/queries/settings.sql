-- name: GetAllSettings :many
SELECT *
FROM settings
ORDER BY key ASC;

-- name: GetSetting :one
SELECT *
FROM settings
WHERE key = ?
LIMIT 1;

-- name: InsertSetting :one
INSERT INTO settings(key, value)
VALUES (?, ?)
RETURNING *;

-- name: UpdateSetting :one
UPDATE settings
SET value = ?
WHERE key = ?
RETURNING *;

-- name: DeleteSetting :exec
DELETE FROM settings
WHERE key = ? AND value = ?;
