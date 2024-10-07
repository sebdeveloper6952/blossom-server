-- +migrate Up
CREATE TABLE IF NOT EXISTS access_control_rules
(
    action   TEXT NOT NULL,
    pubkey   TEXT NOT NULL,
    resource TEXT NOT NULL,

    PRIMARY KEY (action, pubkey, resource)
);

-- +migrate Down
DROP TABLE access_control_rule;
