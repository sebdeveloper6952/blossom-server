-- +migrate Up
CREATE TABLE IF NOT EXISTS mime_types
(
    extension TEXT NOT NULL,
    mime_type TEXT NOT NULL,
    allowed   INTEGER NOT NULL DEFAULT "TRUE"
);

-- +migrate Down
DROP TABLE mime_types;
