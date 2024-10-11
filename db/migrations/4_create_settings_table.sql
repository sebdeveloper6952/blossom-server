-- +migrate Up
CREATE TABLE IF NOT EXISTS settings
(
    key TEXT NOT NULL,
    value TEXT NOT NULL,

    PRIMARY KEY (key, value)
);

-- +migrate Down
DROP TABLE settings;
