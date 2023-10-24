package db

import (
	"context"
	"fmt"

	"database/sql"
)

type Transaction struct {
	*Queries
	db *sql.DB
}

func NewTransaction(db *sql.DB) *Transaction {
	return &Transaction{
		Queries: New(db),
		db:      db,
	}
}

func (transaction *Transaction) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := transaction.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v", err, rollbackErr)
		}

		return err
	}

	return tx.Commit()
}

// below are specific transaction functions

// CreateItemTx is a transaction function that creates an item and its history
func (transaction *Transaction) CreateItemTx(ctx context.Context, arg CreateItemParams) (Item, error) {
	var result Item
	execErr := transaction.ExecTx(ctx, func(q *Queries) error {
		// 1. create item
		var createErr error
		result, createErr = q.CreateItem(ctx, arg)
		if createErr != nil {
			return createErr
		}

		// 2. create history
		var historyArg CreateHistoryParams

		historyArg.ItemID = result.ItemID
		historyArg.IdentifierCode = result.IdentifierCode
		historyArg.Name = result.Name
		historyArg.Holder = result.Holder
		historyArg.ModificationTime = result.ModificationTime
		historyArg.Modifier = result.Modifier
		historyArg.Description = result.Description
		historyArg.Deleted = result.Deleted

		_, historyErr := q.CreateHistory(ctx, historyArg)

		return historyErr
	})

	return result, execErr
}

type DeleteItemsByIdentifierCodeForUpdateParams struct {
	IdentifierCode string `json:"identifier_code"`
	Modifier       int32  `json:"modifier"`
}

// DeleteItemByIdentifierCodeTx is a transaction function that sets an item as deleted and create a new history entry
func (transaction *Transaction) DeleteItemByIdentifierCodeTx(ctx context.Context, arg DeleteItemsByIdentifierCodeForUpdateParams) (Item, error) {
	var result Item
	execErr := transaction.ExecTx(ctx, func(q *Queries) error {
		// 1. get item id by identifier code
		var getItemErr error

		// make sure Deleted is false
		argGetItem := GetItemsByIdentifierCodeForUpdateParams{
			IdentifierCode: arg.IdentifierCode,
			Deleted:        false,
		}

		result, getItemErr = q.GetItemsByIdentifierCodeForUpdate(ctx, argGetItem) // TODO: need bug fix

		if getItemErr != nil {
			return getItemErr
		}

		// 1.2 if Deleted is true, return nil as finished
		if result.Deleted {
			return nil
		}

		// 2. delete item
		var deleteErr error
		result, deleteErr = q.DeleteItem(ctx, DeleteItemParams{
			ItemID:   result.ItemID,
			Modifier: arg.Modifier,
		})

		if deleteErr != nil {
			return deleteErr
		}

		// 3. create history
		var historyArg CreateHistoryParams

		historyArg.ItemID = result.ItemID
		historyArg.IdentifierCode = result.IdentifierCode
		historyArg.Name = result.Name
		historyArg.Holder = result.Holder
		historyArg.ModificationTime = result.ModificationTime
		historyArg.Modifier = result.Modifier
		historyArg.Description = result.Description
		historyArg.Deleted = result.Deleted

		_, historyErr := q.CreateHistory(ctx, historyArg)

		return historyErr
	})

	return result, execErr
}

type UpdateItemByIdentifierCodeParams struct {
	UpdateItemParams
	QueryIdentifierCode string
}

// UpdateItemByIdentifierCodeTx is a transaction function that updates an item and create a new history entry
func (transaction *Transaction) UpdateItemByIdentifierCodeTx(ctx context.Context, arg UpdateItemByIdentifierCodeParams) (Item, error) {
	var result Item
	execErr := transaction.ExecTx(ctx, func(q *Queries) error {
		// 1. get item id by identifier code
		var getItemErr error

		// make sure Deleted is false
		argGetItem := GetItemsByIdentifierCodeForUpdateParams{
			IdentifierCode: arg.QueryIdentifierCode,
			Deleted:        false,
		}

		result, getItemErr = q.GetItemsByIdentifierCodeForUpdate(ctx, argGetItem) // TODO: need bug fix

		if getItemErr != nil {
			return getItemErr
		}

		// 1.2 if Deleted is true, return nil as finished
		if result.Deleted {
			return nil
		}

		// 2. update item
		var updateErr error

		updateArgs := UpdateItemParams{
			ItemID:         result.ItemID, // ItemID used for querying should be results from previous query step
			IdentifierCode: arg.IdentifierCode,
			Name:           arg.Name,
			Holder:         arg.Holder,
			Modifier:       arg.Modifier,
			Description:    arg.Description,
		}

		result, updateErr = q.UpdateItem(ctx, updateArgs)

		if updateErr != nil {
			return updateErr
		}

		// 3. create history
		var historyArg CreateHistoryParams

		historyArg.ItemID = result.ItemID
		historyArg.IdentifierCode = result.IdentifierCode
		historyArg.Name = result.Name
		historyArg.Holder = result.Holder
		historyArg.ModificationTime = result.ModificationTime
		historyArg.Modifier = result.Modifier
		historyArg.Description = result.Description
		historyArg.Deleted = result.Deleted

		_, historyErr := q.CreateHistory(ctx, historyArg)

		return historyErr
	})

	return result, execErr
}

func (transaction *Transaction) CreateUserTx(ctx context.Context, arg CreateUserParams) (User, error) {
	var result User
	execErr := transaction.ExecTx(ctx, func(q *Queries) error {
		var createErr error
		result, createErr = q.CreateUser(ctx, arg)
		return createErr
	})

	return result, execErr
}

func (transaction *Transaction) ListUsersTx(ctx context.Context, arg ListUsersParams) ([]User, error) {
	var result []User
	execErr := transaction.ExecTx(ctx, func(q *Queries) error {
		var listErr error
		result, listErr = q.ListUsers(ctx, arg)
		return listErr
	})

	return result, execErr
}

func (transaction *Transaction) ListItemTx(ctx context.Context, arg ListItemParams) ([]Item, error) {
	var result []Item
	execErr := transaction.ExecTx(ctx, func(q *Queries) error {
		var listErr error
		result, listErr = q.ListItem(ctx, arg)
		return listErr
	})

	return result, execErr
}

type UpdateUserByUsernameParams struct {
	UpdateUserParams
	QueryUsername string
}

func (transaction *Transaction) UpdateUserByUsernameTx(ctx context.Context, arg UpdateUserByUsernameParams) (User, error) {
	var result User
	execErr := transaction.ExecTx(ctx, func(q *Queries) error {
		// 1. get user by username
		user, queryErr := q.GetUserByUsernameForUpdate(ctx, arg.QueryUsername)
		if queryErr != nil {
			return queryErr
		}
		// 2. update user
		updateArgs := UpdateUserParams{
			Uid:         user.Uid,
			Username:    arg.Username,
			Password:    arg.Password,
			Description: arg.Description,
			Activated:   arg.Activated,
		}
		var updateErr error
		result, updateErr = q.UpdateUser(ctx, updateArgs)
		if updateErr != nil {
			return updateErr
		}
		return nil
	})

	return result, execErr
}
