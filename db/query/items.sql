-- name: CreateItem :one
INSERT INTO items (identifier_code, name, holder, modification_time, modifier, description, deleted)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, $4, $5, FALSE)
RETURNING *;

-- name: DeleteItem :one
UPDATE items
SET deleted           = TRUE,
    modification_time = CURRENT_TIMESTAMP,
    modifier          = $2
WHERE item_id = $1
RETURNING *;

-- name: GetItemsByIdentifierCode :one
SELECT *
FROM items
WHERE identifier_code = $1
  AND deleted = $2;

-- name: GetItemsByIdentifierCodeForUpdate :one
SELECT *
FROM items
WHERE identifier_code = $1
  AND deleted = $2
    FOR NO KEY UPDATE;

-- name: GetItemsByItemIdForUpdate :one
SELECT *
FROM items
WHERE item_id = $1
    FOR NO KEY UPDATE;

-- name: GetItemsByName :many
SELECT *
FROM items
WHERE name = $1
  AND deleted = $2;

-- name: GetItemsByHolder :many
SELECT *
FROM items
WHERE holder = $1
  AND deleted = $2;

-- name: UpdateItem :one
UPDATE items
SET identifier_code   = $2,
    name              = $3,
    holder            = $4,
    modification_time = CURRENT_TIMESTAMP,
    modifier          = $5,
    description       = $6
WHERE item_id = $1
RETURNING *;

-- name: ListItem :many
SELECT *
FROM items
WHERE deleted = FALSE
ORDER BY item_id
LIMIT $1 OFFSET $2;