// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"context"
)

type Querier interface {
	CreateHistory(ctx context.Context, arg CreateHistoryParams) (History, error)
	CreateItem(ctx context.Context, arg CreateItemParams) (Items, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	DeleteItem(ctx context.Context, arg DeleteItemParams) (Items, error)
	GetItemsByHolder(ctx context.Context, arg GetItemsByHolderParams) ([]Items, error)
	GetItemsByIdentifierCode(ctx context.Context, arg GetItemsByIdentifierCodeParams) (Items, error)
	GetItemsByIdentifierCodeForUpdate(ctx context.Context, arg GetItemsByIdentifierCodeForUpdateParams) (Items, error)
	GetItemsByItemIdForUpdate(ctx context.Context, itemID int32) (Items, error)
	GetItemsByName(ctx context.Context, arg GetItemsByNameParams) ([]Items, error)
	GetUserByUsername(ctx context.Context, username string) (Users, error)
	GetUserByUsernameForUpdate(ctx context.Context, username string) (Users, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]Users, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error)
}

var _ Querier = (*Queries)(nil)
