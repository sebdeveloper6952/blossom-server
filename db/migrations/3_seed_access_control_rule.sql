-- +migrate Up
INSERT INTO access_control_rules(
    action,
    pubkey,
    resource
) VALUES
("DENY", "ALL", "UPLOAD"),  -- (HEAD|PUT) /upload
("ALLOW", "ALL", "GET"),    -- (GET|HEAD) /<hash>
("DENY", "ALL", "DELETE"), -- (DELETE) /<hash>
("ALLOW", "ALL", "LIST"),   -- (GET) /list<pubkey>
("DENY", "ALL", "MIRROR")  -- (PUT) /mirror
ON CONFLICT(action, pubkey, resource) DO NOTHING;

-- +migrate Down
DELETE FROM access_control_rules;
