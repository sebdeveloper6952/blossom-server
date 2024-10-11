CREATE TABLE IF NOT EXISTS settings
(
    key TEXT NOT NULL,
    value TEXT NOT NULL,

    UNIQUE(key, value)
);
