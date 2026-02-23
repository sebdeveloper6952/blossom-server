-- name: GetBlobsFromPubkey :many
select *
from blobs
where pubkey = ?;

-- name: GetBlobsFromPubkeyPaginated :many
select *
from blobs
where pubkey = ?
  and created > ?
  and created < ?
order by created asc;

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
  blob,
  created
) values (?,?,?,?,?,?)
returning *;

-- name: DeleteBlobFromHash :exec
delete
from blobs
where hash = ?;
