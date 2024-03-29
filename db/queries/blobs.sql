-- name: GetBlobsFromPubkey :many
select *
from blobs
where pubkey = ?;

-- name: GetBlobFromHash :one
select *
from blobs
where hash = ?
limit 1;

-- name: InsertBlob :one
insert into blobs(
  pubkey,
  hash,
  type,
  size,
  created
) values (?,?,?,?,?)
returning *;

-- name: DeleteBlobFromHash :exec
delete
from blobs
where hash = ?;