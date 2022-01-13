package db

import (
	"context"
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