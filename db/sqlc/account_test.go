package db

import (
	"context"
	"testing"

	"github.com/drkennetz/codingchallenge/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
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
}