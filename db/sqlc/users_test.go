package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	// create a new user
	arg := CreateUserParams{
		Username:  "testuser",
		Password:  "testpassword",
		Role:      "testrole",
		Activated: true,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Role, user.Role)
	require.Equal(t, arg.Activated, user.Activated)

	require.NotZero(t, user.Uid)
}

func TestGetUserByUsername(t *testing.T) {
	// create a new user
	arg := CreateUserParams{
		Username:  "testuserwithusername",
		Password:  "testpassword",
		Role:      "testrole",
		Activated: true,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	// get the user by username
	user2, err := testQueries.GetUserByUsername(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.Uid, user2.Uid)
	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.Password, user2.Password)
	require.Equal(t, user.Role, user2.Role)
	require.Equal(t, user.Activated, user2.Activated)
}

func TestGetUserByUsernameForUpdate(t *testing.T) {
	// create a new user
	arg := CreateUserParams{
		Username:  "testuserwithusernameforupdate",
		Password:  "testpassword",
		Role:      "testrole",
		Activated: true,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	// get the user by username
	user2, err := testQueries.GetUserByUsernameForUpdate(context.Background(), arg.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.Uid, user2.Uid)
	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Password, user2.Password)
	require.Equal(t, arg.Role, user2.Role)
	require.Equal(t, arg.Activated, user2.Activated)
}

func TestListUsers(t *testing.T) {
	users := make([]CreateUserParams, 10)
	for i := 0; i < 10; i++ {
		users[i] = CreateUserParams{
			Username:  "testuser" + fmt.Sprintf("%d", i),
			Password:  "testpassword",
			Role:      "testrole_listusers",
			Activated: true,
		}

		_, err := testQueries.CreateUser(context.Background(), users[i])

		require.NoError(t, err)
	}

	args := ListUsersParams{
		Limit:  100,
		Offset: 0,
	}

	ret_users, err := testQueries.ListUsers(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, ret_users)

	testrole_listusers_cnt := 0
	for _, ret_user := range ret_users {
		require.NotEmpty(t, ret_user)
		if ret_user.Role == "testrole_listusers" {
			testrole_listusers_cnt++
		}
	}

	require.Equal(t, 10, testrole_listusers_cnt)
}

func TestUpdateUser(t *testing.T) {
	// create a new user
	arg := CreateUserParams{
		Username:  "test update user username",
		Password:  "testpassword",
		Role:      "testrole",
		Activated: true,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	// get the user by username
	user2, err := testQueries.GetUserByUsername(context.Background(), user.Username)

	// update the user
	arg2 := UpdateUserParams{
		Uid:       user2.Uid,
		Username:  "test update user username updated",
		Password:  "testpassword updated",
		Role:      "testrole updated",
		Activated: false,
	}

	user3, err := testQueries.UpdateUser(context.Background(), arg2)

	require.NoError(t, err)
	require.NotEmpty(t, user3)

	// query the user again
	user4, err := testQueries.GetUserByUsername(context.Background(), user3.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user4)

	require.Equal(t, user3.Uid, user4.Uid)
	require.Equal(t, arg2.Username, user4.Username)
	require.Equal(t, arg2.Password, user4.Password)
	require.Equal(t, arg2.Role, user4.Role)
	require.Equal(t, arg2.Activated, user4.Activated)
}
