-- name: InsertSetting :one
INSERT INTO settings(key, value)
VALUES (?, ?)
ON CONFLICT (key, value)
DO NOTHING
RETURNING *;

-- name: UpdateSetting :one
UPDATE settings
SET value = ?
WHERE key = ?
RETURNING *;;

-- name: DeleteSetting :exec
DELETE FROM settings
WHERE key = ? AND value = ?;
