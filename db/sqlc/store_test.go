package db

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccountUserTx(t *testing.T) {
	store := NewStore(testDB)
	
	// transactions can be a nightmare if concurrency is not handled correctly
	// best way to test transactions well is to run them with several concurrent
	// goroutines. We will test with 5 concurrent transactions

	n := 5

	errs := make(chan error)
	results := make(chan CreateAccountUserTxResult)

	for i := 0; i < n; i++ {
		arg := CreateAccountUserTxParams {
			FullName: fmt.Sprintf("Test User%v", i),
			Email: fmt.Sprintf("Test%vUser@Test.com", i),
			OptedIn: false,
			AdminUser: false,
			Username: fmt.Sprintf("testman%v", i),
			Password: "password",
			Grade: DevLevelBeginner,
		}
		go func() {
			result, err := store.CreateAccountUserTx(context.Background(), arg)
			errs <- err
			results <- result
		}()
	}

	// check results from outside
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check results
		name := result.Account.FullName
		email := result.Account.Email
		optedIn := result.Account.OptedIn
		accountID := result.Account.ID
		userID := result.User.ID
		adminUser := result.User.AdminUser
		username := result.User.Username
		pass := result.User.Password
		grade := result.User.Grade
		require.True(t, strings.Contains(name, "Test User"))
		require.True(t, strings.Contains(email, "User@Test.com"))
		require.False(t, optedIn)
		require.NotZero(t, accountID)
		require.NotZero(t, userID)
		require.False(t, adminUser)
		require.True(t, strings.Contains(username, "testman"))
		require.True(t, pass == "password")
		require.True(t, grade == DevLevelBeginner)
		require.NotZero(t, result.Account.CreatedAt)
		require.NotZero(t, result.Account.UpdatedAt)
		require.NotZero(t, result.User.CreatedAt)
		require.NotZero(t, result.User.UpdatedAt)
		_, err = store.GetAccount(context.Background(), result.Account.ID)
		require.NoError(t, err)
		_, err = store.GetAUser(context.Background(), result.User.ID)
		require.NoError(t, err)
	}




}