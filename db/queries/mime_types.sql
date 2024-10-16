-- name: GetMimeType :one
SELECT *
FROM mime_types
WHERE mime_type = ?
LIMIT 1;

-- name: GetAllMimeTypes :many
SELECT *
FROM mime_types;
