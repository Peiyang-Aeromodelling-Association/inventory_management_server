CREATE TABLE users
(
    uid       SERIAL PRIMARY KEY,
    username  VARCHAR(255) NOT NULL UNIQUE,
    password  VARCHAR(255) NOT NULL,
    role      VARCHAR(255) NOT NULL,
    activated BOOLEAN      NOT NULL DEFAULT TRUE -- we don't delete users, we just deactivate them (history f key)
);

-- Tokens are used for authentication with more security and convenience/
-- It's also a workaround for one user having multiple passwords (e.g. webpage login, card id login)
CREATE TABLE tokens
(
    token       TEXT PRIMARY KEY,                          -- token string for authentication
    uid         INTEGER   NOT NULL REFERENCES users (uid), -- user id of the token owner
    valid_until TIMESTAMP NOT NULL
);


CREATE TABLE items
(
    item_id           SERIAL PRIMARY KEY,                           -- item id for database
    identifier_code   TEXT UNIQUE  NOT NULL,                        -- item id for barcode or qr code
    name              VARCHAR(255) NOT NULL,                        -- name of the item
    holder            INTEGER      NOT NULL REFERENCES users (uid), -- user id of the holder
    modification_time TIMESTAMP    NOT NULL,                        -- time of last record change
    modifier          INTEGER      NOT NULL REFERENCES users (uid), -- user id of the modifier
    description       TEXT                  DEFAULT '',             -- description of the item
    deleted           BOOLEAN      NOT NULL DEFAULT FALSE           -- valid flag
);

CREATE TABLE history
(
    history_id        SERIAL PRIMARY KEY, -- history id
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