CREATE TABLE IF NOT EXISTS access_control_rules
(
    action   TEXT NOT NULL,
    pubkey   TEXT NOT NULL,
    resource TEXT NOT NULL,
    priority INTEGER NOT NULL,

    PRIMARY KEY (action, pubkey, resource)
);
