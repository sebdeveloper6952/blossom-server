-- name: GetACRFromPubkey :many
SELECT *
FROM access_control_rules
WHERE pubkey = ?;

-- name: InsertACR :one
INSERT INTO access_control_rules(
    action,
    pubkey,
    resource,
    priority
)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: DeleteACR :exec
DELETE
FROM access_control_rules
WHERE action = ? AND
      pubkey = ? AND
      resource = ?;
