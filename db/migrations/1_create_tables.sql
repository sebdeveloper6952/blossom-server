-- +migrate Up
CREATE TABLE IF NOT EXISTS blobs
(
    pubkey  TEXT NOT NULL,
    hash    TEXT PRIMARY KEY,
    type    TEXT NOT NULL,
    size    INT NOT NULL,
    blob    BLOB NOT NULL,
    created INT NOT NULL
);

-- +migrate Down
DROP TABLE blobs;
