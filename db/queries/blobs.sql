-- name: GetBlobsFromPubkey :many
select *
from blobs
where pubkey = ?;

-- name: InsertBlob :one
insert into blobs(
  pubkey,
                  hash,
                  type,
                  size,
                  created
) values (?,?,?,?,?)
returning *;