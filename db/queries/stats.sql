-- name: GetStats :one
SELECT SUM(size) AS bytes_stored,
       COUNT(*) AS blob_count,
       COUNT(DISTINCT(pubkey)) AS pubkey_count
FROM blobs;
