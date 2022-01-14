package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/drkennetz/codingchallenge/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateAccountParams {
		FullName: util.RandomFullName(),
		Email: util.RandomEmail(),
		OptedIn: false,
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	

	arg2 := CreateUserParams {
		AccountID: account.ID,
		Username: util.RandomUsername(),
		Grade: DevLevel(util.RandomGrade()), 
	}

	user, err := testQueries.CreateUser(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg2.AccountID, user.AccountID)
	require.Equal(t, arg2.Username, user.Username)
	require.Equal(t, arg2.Grade, user.Grade)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetAUser(context.Background(), user1.ID)
	require.NotEmpty(t, user2)
	require.NoError(t, err)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.AccountID, user2.AccountID)
	require.Equal(t, user1.Grade, user2.Grade)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	err := testQueries.DeleteAccount(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetAccount(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams {
		Limit: 5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestUpdateUserGrade(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdateUserGradeParams {
		ID: user1.ID,
		Grade: DevLevelMidlevel,
	}
	user2, err := testQueries.UpdateUserGrade(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.True(t, user2.Grade == DevLevelMidlevel)
}

func TestUpdateUsername(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdateUsernameParams {
		ID: user1.ID,
		Username: "Testing This",
	}
	user2, err := testQueries.UpdateUsername(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.True(t, user2.Username == "Testing This")
}