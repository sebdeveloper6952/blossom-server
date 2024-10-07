-- name: GetACRFromPubkey :many
SELECT *
FROM access_control_rules
WHERE pubkey = ?;

-- name: GetACRFromPubkeyResource :one
SELECT *
FROM access_control_rules
WHERE pubkey = ? AND
      resource = ?
LIMIT 1;

-- name: GetACR :one
SELECT *
FROM access_control_rules
WHERE action = ? AND
      pubkey = ? AND
      resource = ?
LIMIT 1;

-- name: InsertACR :one
INSERT INTO access_control_rules(
    action,
    pubkey,
    resource
)
VALUES (?, ?, ?)
ON CONFLICT (
    action, 
    pubkey, 
    resource
) DO NOTHING
RETURNING *;

-- name: DeleteACR :exec
DELETE
FROM access_control_rules
WHERE action = ? AND
      pubkey = ? AND
      resource = ?;
