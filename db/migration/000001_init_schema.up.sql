CREATE TABLE users
(
    uid       INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username  VARCHAR(255) NOT NULL UNIQUE,
    password  VARCHAR(255) NOT NULL,
    description TEXT DEFAULT '',
    activated BOOLEAN      NOT NULL DEFAULT TRUE -- we don't delete users, we just deactivate them (history f key)
);

CREATE TABLE items
(
    item_id           INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY, -- item id for database
    identifier_code   TEXT UNIQUE  NOT NULL,                            -- item id for barcode or qr code
    name              VARCHAR(255) NOT NULL,                            -- name of the item
    holder            INTEGER      NOT NULL REFERENCES users (uid),     -- user id of the holder
    modification_time TIMESTAMP    NOT NULL,                            -- time of last record change
    modifier          INTEGER      NOT NULL REFERENCES users (uid),     -- user id of the modifier
    description       TEXT                  DEFAULT '',                 -- description of the item
    deleted           BOOLEAN      NOT NULL DEFAULT FALSE               -- valid flag
);

CREATE TABLE history
(
    history_id        INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY, -- history id
    item_id           INTEGER      NOT NULL REFERENCES items (item_id),
    identifier_code   TEXT         NOT NULL,
    name              VARCHAR(255) NOT NULL,
    holder            INTEGER      NOT NULL REFERENCES users (uid),
    modification_time TIMESTAMP    NOT NULL,
    modifier          INTEGER      NOT NULL REFERENCES users (uid),
    description       TEXT                  DEFAULT '',
    deleted           BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_items_identifier_code ON items (identifier_code);
CREATE INDEX idx_items_name ON items (name);
CREATE INDEX idx_items_holder ON items (holder);