package db

import (
	"context"
	"database/sql"
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

	n := 2

	errs := make(chan error)
	results := make(chan CreateAccountUserTxResult)

	for i := 0; i < n; i++ {
		// deadlock debugging (I do not have a deadlock, but I want to add this safety)
		txName := fmt.Sprintf("tx %d", i+1)
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
			// This will allow the context to hold the transaction name
			ctx := context.WithValue(context.Background(), txKey, txName)
			result, err := store.CreateAccountUserTx(ctx, arg)
			errs <- err
			results <- result
			err = store.DeleteAccountUserTx(ctx, DeleteAccountUserTxParams{
				Email: arg.Email,
				Username: arg.Username,
			})
			errs <- err
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

func TestDeleteAccountUserTx(t *testing.T) {
	store := NewStore(testDB)
	arg := CreateAccountUserTxParams {
		FullName: "Birdman Junior",
		Email: "birdman@aol.com",
		OptedIn: true,
		AdminUser: false,
		Username: "birdman1",
		Password: "password",
		Grade: DevLevelJunior,
	}

	txName := "tx 1"
	ctx := context.WithValue(context.Background(), txKey, txName)
	result, err := store.CreateAccountUserTx(ctx, arg)
	require.NoError(t, err)
	require.Equal(t, result.Account.Email, arg.Email)
	require.Equal(t, result.Account.FullName, arg.FullName)
	require.Equal(t, result.User.Username, arg.Username)
	require.Equal(t, result.User.AdminUser, arg.AdminUser)
	require.Equal(t, result.User.Grade, arg.Grade)

	deleteArg := DeleteAccountUserTxParams {
		Email: arg.Email,
		Username: arg.Username,
	}
	err = store.DeleteAccountUserTx(ctx, deleteArg)
	require.NoError(t, err)
	retrieved, err := store.GetAUser(context.Background(), result.User.ID)
	require.Equal(t, err, sql.ErrNoRows)
	require.Empty(t, retrieved)
}

func TestCreateQuestionCategoryTx(t *testing.T) {
	store := NewStore(testDB)
	arg := CreateQuestionCategoryTxParams {
		ChallengeName: "Test Challenge",
		Description: "Test Description",
		Example: "Test Example",
		Difficulty: 1,
		Complexity: "O(n) | O(1)",
		CompletionTime: 10,
		QuestionType: QuestionTypePractice,
		Category: "math",
	}
	txName := "Successful Question Category Tx"
	ctx := context.WithValue(context.Background(), txKey, txName)
	result, err := store.CreateQuestionCategoryTx(ctx, arg)
	require.NoError(t, err)
	require.Equal(t, result.Question.ChallengeName, arg.ChallengeName)
	require.Equal(t, result.Question.Description, arg.Description)
	require.Equal(t, result.Question.Example, arg.Example)
	require.Equal(t, result.Question.Difficulty, arg.Difficulty)
	require.Equal(t, result.Question.Complexity, arg.Complexity)
	require.Equal(t, result.Question.CompletionTime, arg.CompletionTime)
	require.Equal(t, result.Question.QuestionType, arg.QuestionType)
	require.Equal(t, result.QuestionCategory.QuestionID, result.Question.ID)
	require.NotZero(t, result.QuestionCategory.ID)

	// delete question category entry
	// delete question entry (category must already exist)
}

// test bad question
// test bad category
// test bad question category(?)

