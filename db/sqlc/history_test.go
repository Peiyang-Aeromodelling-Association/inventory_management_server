package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateHistory(t *testing.T) {
	// create a user to prevent foreign key constraint error
	argUser := CreateUserParams{
		Username:    "testuserwithusernameforhistory",
		Password:    "testpassword",
		Description: sql.NullString{String: "test description", Valid: true},
		Activated:   true,
	}

	user, err := testQueries.CreateUser(context.Background(), argUser)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	// create a item to prevent foreign key constraint error
	argItem := CreateItemParams{
		IdentifierCode: "123testcode",
		Name:           "test",
		Holder:         user.Uid,
		Modifier:       user.Uid,
		Description:    sql.NullString{String: "test description", Valid: true},
	}

	item, err := testQueries.CreateItem(context.Background(), argItem)

	require.NoError(t, err)
	require.NotEmpty(t, item)

	// create a new history entry
	arg := CreateHistoryParams{
		ItemID:           item.ItemID,
		IdentifierCode:   item.IdentifierCode,
		Name:             item.Name,
		Holder:           item.Holder,
		ModificationTime: item.ModificationTime,
		Modifier:         item.Modifier,
		Description:      item.Description,
		Deleted:          false,
	}

	history, err := testQueries.CreateHistory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, history)

	// check if the history entry is created correctly
	require.Equal(t, arg.ItemID, history.ItemID)
	require.Equal(t, arg.IdentifierCode, history.IdentifierCode)
	require.Equal(t, arg.Name, history.Name)
	require.Equal(t, arg.Holder, history.Holder)
	require.Equal(t, arg.ModificationTime, history.ModificationTime)
	require.Equal(t, arg.Modifier, history.Modifier)
	require.Equal(t, arg.Description, history.Description)
	require.Equal(t, arg.Deleted, history.Deleted)
}
