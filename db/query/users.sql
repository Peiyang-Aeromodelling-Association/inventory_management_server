-- name: CreateUser :one
INSERT INTO users (username, password, description, activated)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET username  = $2,
    password  = $3,
    description = $4,
    activated = $5
WHERE uid = $1
RETURNING *;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1;

-- name: GetUserByUsernameForUpdate :one
SELECT *
FROM users
WHERE username = $1
    FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT *
FROM users
WHERE activated = TRUE
ORDER BY uid
LIMIT $1 OFFSET $2;
