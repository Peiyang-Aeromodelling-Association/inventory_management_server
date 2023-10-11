-- name: CreateHistory :one
INSERT INTO history (item_id, identifier_code, name, holder, modification_time, modifier, description, deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;
