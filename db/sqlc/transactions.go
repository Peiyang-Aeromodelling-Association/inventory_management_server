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
func (transaction *Transaction) CreateItemTx(ctx context.Context, arg CreateItemParams) (Items, error) {
	var result Items
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
	GetItemsByIdentifierCodeForUpdateParams
	Modifier int32 `json:"modifier"`
}

// // DeleteItemByIdentifierCodeTx is a transaction function that sets an item as deleted and create a new history entry
// func (transaction *Transaction) DeleteItemByIdentifierCodeTx(ctx context.Context, arg DeleteItemsByIdentifierCodeForUpdateParams) (Items, error) {
// 	var result Items
// 	execErr := transaction.ExecTx(ctx, func(q *Queries) error {
// 		// 1. get item id by identifier code
// 		var getItemErr error
// 		result, getItemErr = q.GetItemsByIdentifierCodeForUpdate(ctx, arg) // TODO: need bug fix

// 		if getItemErr != nil {
// 			return getItemErr
// 		}

// 		// 1.2 if Deleted is true, return nil as finished
// 		if result.Deleted {
// 			return nil
// 		}

// 		// 2. delete item
// 		var deleteErr error
// 		_, deleteErr = q.DeleteItem(ctx, DeleteItemParams{
// 			ItemID:   result.ItemID,
// 			Modifier: arg.Modifier,
// 		})

// 		if deleteErr != nil {
// 			return deleteErr
// 		}

// 		// 3. create history
// 		var historyArg CreateHistoryParams

// 		historyArg.ItemID = result.ItemID
// 		historyArg.IdentifierCode = result.IdentifierCode
// 		historyArg.Name = result.Name
// 		historyArg.Holder = result.Holder
// 		historyArg.ModificationTime = result.ModificationTime
// 		historyArg.Modifier = result.Modifier
// 		historyArg.Description = result.Description
// 		historyArg.Deleted = result.Deleted

// 		_, historyErr := q.CreateHistory(ctx, historyArg)

// 		return historyErr
// 	})

// 	return result, execErr
// }
