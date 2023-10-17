// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: items.sql

package db

import (
	"context"
	"database/sql"
)

const createItem = `-- name: CreateItem :one
INSERT INTO items (identifier_code, name, holder, modification_time, modifier, description, deleted)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, $4, $5, FALSE)
RETURNING item_id, identifier_code, name, holder, modification_time, modifier, description, deleted
`

type CreateItemParams struct {
	IdentifierCode string         `json:"identifier_code"`
	Name           string         `json:"name"`
	Holder         int32          `json:"holder"`
	Modifier       int32          `json:"modifier"`
	Description    sql.NullString `json:"description"`
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem,
		arg.IdentifierCode,
		arg.Name,
		arg.Holder,
		arg.Modifier,
		arg.Description,
	)
	var i Item
	err := row.Scan(
		&i.ItemID,
		&i.IdentifierCode,
		&i.Name,
		&i.Holder,
		&i.ModificationTime,
		&i.Modifier,
		&i.Description,
		&i.Deleted,
	)
	return i, err
}

const deleteItem = `-- name: DeleteItem :one
UPDATE items
SET deleted           = TRUE,
    modification_time = CURRENT_TIMESTAMP,
    modifier          = $2
WHERE item_id = $1
RETURNING item_id, identifier_code, name, holder, modification_time, modifier, description, deleted
`

type DeleteItemParams struct {
	ItemID   int32 `json:"item_id"`
	Modifier int32 `json:"modifier"`
}

func (q *Queries) DeleteItem(ctx context.Context, arg DeleteItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, deleteItem, arg.ItemID, arg.Modifier)
	var i Item
	err := row.Scan(
		&i.ItemID,
		&i.IdentifierCode,
		&i.Name,
		&i.Holder,
		&i.ModificationTime,
		&i.Modifier,
		&i.Description,
		&i.Deleted,
	)
	return i, err
}

const getItemsByHolder = `-- name: GetItemsByHolder :many
SELECT item_id, identifier_code, name, holder, modification_time, modifier, description, deleted
FROM items
WHERE holder = $1
  AND deleted = $2
`

type GetItemsByHolderParams struct {
	Holder  int32 `json:"holder"`
	Deleted bool  `json:"deleted"`
}

func (q *Queries) GetItemsByHolder(ctx context.Context, arg GetItemsByHolderParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getItemsByHolder, arg.Holder, arg.Deleted)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Item{}
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ItemID,
			&i.IdentifierCode,
			&i.Name,
			&i.Holder,
			&i.ModificationTime,
			&i.Modifier,
			&i.Description,
			&i.Deleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemsByIdentifierCode = `-- name: GetItemsByIdentifierCode :one
SELECT item_id, identifier_code, name, holder, modification_time, modifier, description, deleted
FROM items
WHERE identifier_code = $1
  AND deleted = $2
`

type GetItemsByIdentifierCodeParams struct {
	IdentifierCode string `json:"identifier_code"`
	Deleted        bool   `json:"deleted"`
}

func (q *Queries) GetItemsByIdentifierCode(ctx context.Context, arg GetItemsByIdentifierCodeParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItemsByIdentifierCode, arg.IdentifierCode, arg.Deleted)
	var i Item
	err := row.Scan(
		&i.ItemID,
		&i.IdentifierCode,
		&i.Name,
		&i.Holder,
		&i.ModificationTime,
		&i.Modifier,
		&i.Description,
		&i.Deleted,
	)
	return i, err
}

const getItemsByIdentifierCodeForUpdate = `-- name: GetItemsByIdentifierCodeForUpdate :one
SELECT item_id, identifier_code, name, holder, modification_time, modifier, description, deleted
FROM items
WHERE identifier_code = $1
  AND deleted = $2
    FOR NO KEY UPDATE
`

type GetItemsByIdentifierCodeForUpdateParams struct {
	IdentifierCode string `json:"identifier_code"`
	Deleted        bool   `json:"deleted"`
}

func (q *Queries) GetItemsByIdentifierCodeForUpdate(ctx context.Context, arg GetItemsByIdentifierCodeForUpdateParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItemsByIdentifierCodeForUpdate, arg.IdentifierCode, arg.Deleted)
	var i Item
	err := row.Scan(
		&i.ItemID,
		&i.IdentifierCode,
		&i.Name,
		&i.Holder,
		&i.ModificationTime,
		&i.Modifier,
		&i.Description,
		&i.Deleted,
	)
	return i, err
}

const getItemsByItemIdForUpdate = `-- name: GetItemsByItemIdForUpdate :one
SELECT item_id, identifier_code, name, holder, modification_time, modifier, description, deleted
FROM items
WHERE item_id = $1
    FOR NO KEY UPDATE
`

func (q *Queries) GetItemsByItemIdForUpdate(ctx context.Context, itemID int32) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItemsByItemIdForUpdate, itemID)
	var i Item
	err := row.Scan(
		&i.ItemID,
		&i.IdentifierCode,
		&i.Name,
		&i.Holder,
		&i.ModificationTime,
		&i.Modifier,
		&i.Description,
		&i.Deleted,
	)
	return i, err
}

const getItemsByName = `-- name: GetItemsByName :many
SELECT item_id, identifier_code, name, holder, modification_time, modifier, description, deleted
FROM items
WHERE name = $1
  AND deleted = $2
`

type GetItemsByNameParams struct {
	Name    string `json:"name"`
	Deleted bool   `json:"deleted"`
}

func (q *Queries) GetItemsByName(ctx context.Context, arg GetItemsByNameParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getItemsByName, arg.Name, arg.Deleted)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Item{}
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ItemID,
			&i.IdentifierCode,
			&i.Name,
			&i.Holder,
			&i.ModificationTime,
			&i.Modifier,
			&i.Description,
			&i.Deleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listItem = `-- name: ListItem :many
SELECT item_id, identifier_code, name, holder, modification_time, modifier, description, deleted
FROM items
WHERE deleted = FALSE
ORDER BY item_id
LIMIT $1 OFFSET $2
`

type ListItemParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListItem(ctx context.Context, arg ListItemParams) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItem, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Item{}
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ItemID,
			&i.IdentifierCode,
			&i.Name,
			&i.Holder,
			&i.ModificationTime,
			&i.Modifier,
			&i.Description,
			&i.Deleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateItem = `-- name: UpdateItem :one
UPDATE items
SET identifier_code   = $2,
    name              = $3,
    holder            = $4,
    modification_time = CURRENT_TIMESTAMP,
    modifier          = $5,
    description       = $6
WHERE item_id = $1
RETURNING item_id, identifier_code, name, holder, modification_time, modifier, description, deleted
`

type UpdateItemParams struct {
	ItemID         int32          `json:"item_id"`
	IdentifierCode string         `json:"identifier_code"`
	Name           string         `json:"name"`
	Holder         int32          `json:"holder"`
	Modifier       int32          `json:"modifier"`
	Description    sql.NullString `json:"description"`
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, updateItem,
		arg.ItemID,
		arg.IdentifierCode,
		arg.Name,
		arg.Holder,
		arg.Modifier,
		arg.Description,
	)
	var i Item
	err := row.Scan(
		&i.ItemID,
		&i.IdentifierCode,
		&i.Name,
		&i.Holder,
		&i.ModificationTime,
		&i.Modifier,
		&i.Description,
		&i.Deleted,
	)
	return i, err
}
