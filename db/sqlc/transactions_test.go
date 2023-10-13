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
		Username:  "testuserfortestdeleteitem",
		Password:  "testpassword",
		Role:      "testrole",
		Activated: true,
	}

	user, err := testQueries.CreateUser(context.Background(), argCreateUser)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	// 2. Create an item
	args := CreateItemParams{
		IdentifierCode: "testidentifiercodefordelete",
		Name:           "testname",
		Holder:         user.Uid,
		Modifier:       user.Uid,
		Description:    sql.NullString{String: "testdescription", Valid: true},
	}

	transactionCreate := NewTransaction(testDB)
	item, err := transactionCreate.CreateItemTx(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, item.ItemID)
	require.Equal(t, args.IdentifierCode, item.IdentifierCode)
	require.Equal(t, args.Name, item.Name)
	require.Equal(t, args.Holder, item.Holder)
	require.Equal(t, args.Modifier, item.Modifier)
	require.Equal(t, args.Description, item.Description)
	require.Equal(t, false, item.Deleted)

	argDeleteItem := DeleteItemsByIdentifierCodeForUpdateParams{
		IdentifierCode: "testidentifiercodefordelete",
		Modifier:       user.Uid,
	}

	transactionDelete := NewTransaction(testDB)
	deleteItem, err := transactionDelete.DeleteItemByIdentifierCodeTx(context.Background(), argDeleteItem)

	require.NoError(t, err)
	require.NotEmpty(t, deleteItem.ItemID)
	require.Equal(t, argDeleteItem.IdentifierCode, deleteItem.IdentifierCode)
	require.Equal(t, argDeleteItem.Modifier, deleteItem.Modifier)
	require.Equal(t, true, deleteItem.Deleted)
}

func TestUpdateItemByIdentifierCodeTx(t *testing.T) {
	// 1. Create a user first
	argCreateUser := CreateUserParams{
		Username:  "testuserfortestupdateitem",
		Password:  "testpassword",
		Role:      "testrole",
		Activated: true,
	}

	user, err := testQueries.CreateUser(context.Background(), argCreateUser)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	// 2. Create an item
	args := CreateItemParams{
		IdentifierCode: "testidentifiercodeforupdate",
		Name:           "testname",
		Holder:         user.Uid,
		Modifier:       user.Uid,
		Description:    sql.NullString{String: "testdescription", Valid: true},
	}

	transactionCreate := NewTransaction(testDB)
	item, err := transactionCreate.CreateItemTx(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, item.ItemID)
	require.Equal(t, args.IdentifierCode, item.IdentifierCode)
	require.Equal(t, args.Name, item.Name)
	require.Equal(t, args.Holder, item.Holder)
	require.Equal(t, args.Modifier, item.Modifier)
	require.Equal(t, args.Description, item.Description)
	require.Equal(t, false, item.Deleted)

	argsUpdateItem := UpdateItemByIdentifierCodeParams{
		UpdateItemParams{
			IdentifierCode: "testidentifiercodeforupdate_updated",
			Name:           "testname_updated",
			Holder:         user.Uid,
			Modifier:       user.Uid,
			Description:    sql.NullString{String: "testdescription_updated", Valid: true},
		},
		item.IdentifierCode,
	}

	transactionUpdate := NewTransaction(testDB)
	updateItem, err := transactionUpdate.UpdateItemByIdentifierCodeTx(context.Background(), argsUpdateItem)

	require.NoError(t, err)
	require.NotEmpty(t, updateItem.ItemID)
	require.Equal(t, argsUpdateItem.IdentifierCode, updateItem.IdentifierCode)
	require.Equal(t, argsUpdateItem.Name, updateItem.Name)
	require.Equal(t, argsUpdateItem.Holder, updateItem.Holder)
	require.Equal(t, argsUpdateItem.Modifier, updateItem.Modifier)
	require.Equal(t, argsUpdateItem.Description, updateItem.Description)
	require.Equal(t, false, updateItem.Deleted)
}
