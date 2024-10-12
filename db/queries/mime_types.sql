-- name: GetAllMimeTypes :many
SELECT *
FROM mime_types
ORDER BY mime_type ASC;

-- name: UpdateMimeType :one
UPDATE mime_types
SET allowed = ?
WHERE mime_type = ?
RETURNING *;
