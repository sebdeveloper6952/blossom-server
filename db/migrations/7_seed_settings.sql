-- +migrate Up
INSERT INTO settings(key, value)
VALUES ('UPLOAD_MAX_SIZE_BYTES', '999999999999999999');

-- +migrate Down
DELETE * FROM settings;
