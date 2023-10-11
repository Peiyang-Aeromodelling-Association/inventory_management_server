// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, password, role, activated)
VALUES ($1, $2, $3, $4)
RETURNING uid, username, password, role, activated
`

type CreateUserParams struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Activated bool   `json:"activated"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Password,
		arg.Role,
		arg.Activated,
	)
	var i Users
	err := row.Scan(
		&i.Uid,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.Activated,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT uid, username, password, role, activated
FROM users
WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i Users
	err := row.Scan(
		&i.Uid,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.Activated,
	)
	return i, err
}

const getUserByUsernameForUpdate = `-- name: GetUserByUsernameForUpdate :one
SELECT uid, username, password, role, activated
FROM users
WHERE username = $1
    FOR NO KEY UPDATE
`

func (q *Queries) GetUserByUsernameForUpdate(ctx context.Context, username string) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsernameForUpdate, username)
	var i Users
	err := row.Scan(
		&i.Uid,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.Activated,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT uid, username, password, role, activated
FROM users
ORDER BY uid
LIMIT $1 OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]Users, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Users{}
	for rows.Next() {
		var i Users
		if err := rows.Scan(
			&i.Uid,
			&i.Username,
			&i.Password,
			&i.Role,
			&i.Activated,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET username  = $2,
    password  = $3,
    role      = $4,
    activated = $5
WHERE uid = $1
RETURNING uid, username, password, role, activated
`

type UpdateUserParams struct {
	Uid       int32  `json:"uid"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Activated bool   `json:"activated"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Uid,
		arg.Username,
		arg.Password,
		arg.Role,
		arg.Activated,
	)
	var i Users
	err := row.Scan(
		&i.Uid,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.Activated,
	)
	return i, err
}