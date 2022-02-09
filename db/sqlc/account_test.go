package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/drkennetz/codingchallenge/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account{
	arg := CreateAccountParams {
		FullName: util.RandomFullName(),
		Email: util.RandomEmail(),
		OptedIn: false,
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	// tests that arg = expected insertion into db
	require.Equal(t, arg.FullName, account.FullName)
	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.OptedIn, account.OptedIn)

	// tests that rest of fields should be autopopulated from acc creation
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.UpdatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	// create account first - unit tests should be independent from each other
	// a change in one test should never affect the result of another one
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.FullName, account2.FullName)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.OptedIn, account2.OptedIn)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
	require.WithinDuration(t, account1.UpdatedAt, account2.UpdatedAt, time.Second)
	err = testQueries.DeleteAccount(context.Background(), account2.ID)
	require.NoError(t, err)
}

func TestUpdateAccountEmail(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountEmailParams {
		ID: account1.ID,
		Email: util.RandomEmail(),
	}

	account2, err := testQueries.UpdateAccountEmail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.NotEqual(t, account1.Email, account2.Email)
	err = testQueries.DeleteAccount(context.Background(), account2.ID)
	require.NoError(t, err)
}

func TestUpdateAccountFullName(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountNameParams {
		ID: account1.ID,
		FullName: util.RandomFullName(),
	}

	account2, err := testQueries.UpdateAccountName(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.NotEqual(t, account1.FullName, account2.FullName)
	err = testQueries.DeleteAccount(context.Background(), account2.ID)
	require.NoError(t, err)
}

func TestUpdateAccountOptedIn(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountOptedInParams {
		ID: account1.ID,
		OptedIn: true,
	}

	account2, err := testQueries.UpdateAccountOptedIn(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.True(t, account2.OptedIn)
	err = testQueries.DeleteAccount(context.Background(), account2.ID)
	require.NoError(t, err)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams {
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		err = testQueries.DeleteAccount(context.Background(), account.ID)
		require.NoError(t, err)
	}
}