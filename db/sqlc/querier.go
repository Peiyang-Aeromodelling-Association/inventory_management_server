// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"context"
)

type Querier interface {
	CreateHistory(ctx context.Context, arg CreateHistoryParams) (History, error)
	CreateItem(ctx context.Context, arg CreateItemParams) (Item, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteItem(ctx context.Context, arg DeleteItemParams) (Item, error)
	GetItemsByHolder(ctx context.Context, arg GetItemsByHolderParams) ([]Item, error)
	GetItemsByIdentifierCode(ctx context.Context, arg GetItemsByIdentifierCodeParams) (Item, error)
	GetItemsByIdentifierCodeForUpdate(ctx context.Context, arg GetItemsByIdentifierCodeForUpdateParams) (Item, error)
	GetItemsByItemIdForUpdate(ctx context.Context, itemID int32) (Item, error)
	GetItemsByName(ctx context.Context, arg GetItemsByNameParams) ([]Item, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUserByUsernameForUpdate(ctx context.Context, username string) (User, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
