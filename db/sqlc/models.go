// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"database/sql"
	"time"
)

type History struct {
	HistoryID        int32          `json:"history_id"`
	ItemID           int32          `json:"item_id"`
	IdentifierCode   string         `json:"identifier_code"`
	Name             string         `json:"name"`
	Holder           int32          `json:"holder"`
	ModificationTime time.Time      `json:"modification_time"`
	Modifier         int32          `json:"modifier"`
	Description      sql.NullString `json:"description"`
	Deleted          bool           `json:"deleted"`
}

type Item struct {
	ItemID           int32          `json:"item_id"`
	IdentifierCode   string         `json:"identifier_code"`
	Name             string         `json:"name"`
	Holder           int32          `json:"holder"`
	ModificationTime time.Time      `json:"modification_time"`
	Modifier         int32          `json:"modifier"`
	Description      sql.NullString `json:"description"`
	Deleted          bool           `json:"deleted"`
}

type User struct {
	Uid         int32          `json:"uid"`
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	Description sql.NullString `json:"description"`
	Activated   bool           `json:"activated"`
}
