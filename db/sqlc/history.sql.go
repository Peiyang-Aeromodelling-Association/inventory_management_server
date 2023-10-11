// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: history.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createHistory = `-- name: CreateHistory :one
INSERT INTO history (item_id, identifier_code, name, holder, modification_time, modifier, description, deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING history_id, item_id, identifier_code, name, holder, modification_time, modifier, description, deleted
`

type CreateHistoryParams struct {
	ItemID           int32          `json:"item_id"`
	IdentifierCode   string         `json:"identifier_code"`
	Name             string         `json:"name"`
	Holder           int32          `json:"holder"`
	ModificationTime time.Time      `json:"modification_time"`
	Modifier         int32          `json:"modifier"`
	Description      sql.NullString `json:"description"`
	Deleted          bool           `json:"deleted"`
}

func (q *Queries) CreateHistory(ctx context.Context, arg CreateHistoryParams) (History, error) {
	row := q.db.QueryRowContext(ctx, createHistory,
		arg.ItemID,
		arg.IdentifierCode,
		arg.Name,
		arg.Holder,
		arg.ModificationTime,
		arg.Modifier,
		arg.Description,
		arg.Deleted,
	)
	var i History
	err := row.Scan(
		&i.HistoryID,
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