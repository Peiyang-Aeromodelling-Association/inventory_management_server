package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetItemsByHolder(t *testing.T) {
	// create a user to prevent foreign key constraint error
	argUser := CreateUserParams{
		Username:  "testuserwithusernameforgetitembyholder",
		Password:  "testpassword",
		Activated: true,
	}

	user, err := testQueries.CreateUser(context.Background(), argUser)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	// 2. create 10 items
	itemsArgs := make([]CreateItemParams, 10)
	itemResults := make([]Item, 10)

	for i := 0; i < 10; i++ {
		var err error
		itemsArgs[i] = CreateItemParams{
			IdentifierCode: "testIdentifierCodeForTestGetItemsByHolder_" + fmt.Sprintf("%d", i),
			Name:           "testNameForTestGetItemsByHolder_" + fmt.Sprintf("%d", i),
			Holder:         user.Uid,
			Modifier:       user.Uid,
			Description:    sql.NullString{String: "testdescription", Valid: true},
		}

		itemResults[i], err = testQueries.CreateItem(context.Background(), itemsArgs[i])

		require.NoError(t, err)
	}

	// 3. query by holder
	queryResults, queryErr := testQueries.GetItemsByHolder(context.Background(), GetItemsByHolderParams{
		Holder:  user.Uid,
		Deleted: false,
	})

	require.NoError(t, queryErr)

	valid_cnt := 0

	for _, ret_item := range queryResults {
		require.NotEmpty(t, ret_item)
		if ret_item.Holder == user.Uid {
			valid_cnt++
		} else {
			log.Fatal("user id does not match")
		}
	}

	require.Equal(t, valid_cnt, 10)
}
