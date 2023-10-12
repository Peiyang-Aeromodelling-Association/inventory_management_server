package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateItemTx(t *testing.T) {
	args := CreateItemParams{
		IdentifierCode: "testidentifiercode",
		Name:           "testname",
		Holder:         1,
		Modifier:       1,
		Description:    sql.NullString{String: "testdescription", Valid: true},
	}

	transaction := NewTransaction(testDB)
	item, err := transaction.CreateItemTx(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, item.ItemID)
	require.Equal(t, args.IdentifierCode, item.IdentifierCode)
	require.Equal(t, args.Name, item.Name)
	require.Equal(t, args.Holder, item.Holder)
	require.Equal(t, args.Modifier, item.Modifier)
	require.Equal(t, args.Description, item.Description)
	require.Equal(t, false, item.Deleted)
}

func TestDeleteItemByIdentifierCodeTx(t *testing.T) {
	// 1. Create a user first
	argCreateUser := CreateUserParams{
		Username:  "testuser",
		Password:  "testpassword",
		Role:      "testrole",
		Activated: true,
	}

	user, err := testQueries.CreateUser(context.Background(), argCreateUser)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	argDeleteItem := DeleteItemsByIdentifierCodeForUpdateParams{
		GetItemsByIdentifierCodeForUpdateParams{
			IdentifierCode: "testidentifiercode",
			Deleted:        false,
		},
		user.Uid,
	}

	transaction := NewTransaction(testDB)
	deleteItem, err := transaction.DeleteItemByIdentifierCodeTx(context.Background(), argDeleteItem)

	require.NoError(t, err)
	require.NotEmpty(t, deleteItem.ItemID)
	require.Equal(t, argDeleteItem.IdentifierCode, deleteItem.IdentifierCode)
	require.Equal(t, argDeleteItem.Modifier, deleteItem.Modifier)
	require.Equal(t, true, deleteItem.Deleted)
}
