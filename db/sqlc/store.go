package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	CreateAccountUserTx(ctx context.Context, arg CreateAccountUserTxParams) (CreateAccountUserTxResult, error)
	DeleteAccountUserTx(ctx context.Context, arg DeleteAccountUserTxParams) (error)
}

// SQLStore provides all functions to execute SQL db queries and transactions
// extending the functionality of queries since one query only handles one operation
type SQLStore struct {
	*Queries // composition - all Queries will be available to Store
	db *sql.DB
}

// NewStore builds a new Store object and returns it
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db: db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	// if we have an error, we need to rollback the transaction.
	// if the rollback errors, we should report both errors.
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		// If the rollback is successful, we should just return the error of the query
		return err
	}

	// if all is well, we should commit the transaction and return its error to the caller.
	return tx.Commit()
}

// account + user creation(?) -> a user is created during account creation
// This is challenging because the input user params must be retrieved from the created account

// CreateAccountUserTxParams will store all the information needed to create an account and a user
type CreateAccountUserTxParams struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	OptedIn bool `json:"opted_in"`
	AdminUser bool `json:"admin_user"`
	Username string `json:"username"`
	Password string `json:"password"`
	Grade DevLevel `json:"grade"`
}

// CreateAccountUserTxResult stores the results of the transaction
type CreateAccountUserTxResult struct {
	Account Account
	User    User
}

// we are creating an empty object of type struct to pass into our context
var txKey = struct{}{}

// CreateAccountUserTx creates an account and a user in a transaction
func (store *SQLStore) CreateAccountUserTx(ctx context.Context, arg CreateAccountUserTxParams) (CreateAccountUserTxResult, error) {
	var result CreateAccountUserTxResult

	// we are accessing the result variable from inside the tx function. similar to the arg variable. This makes the callback fn a closure
	// sin golangs support for generics type, closure is often used when we get the result from a callback function
	// this is because callback function itself doesn't know the exact type of the result
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		txName := ctx.Value(txKey)

		fmt.Println(txName, "create account")
		result.Account, err = q.CreateAccount(ctx, CreateAccountParams{
			FullName: arg.FullName,
			Email: arg.Email,
			OptedIn: arg.OptedIn,
		})
		if err != nil {
			return err
		}

		// to ensure deadlock prevention (IE we are not accessing the user before it is created)
		// we can query the database for the account. This ensures the account exists in the database prior to 
		// creating a user from that account:
		fmt.Println(txName, "make sure account exists in db by retrieving it and locking it until user is created")
		resultAcc, err := q.GetAccount(ctx, result.Account.ID)
		if err != nil {
			return err
		}
		fmt.Println(txName, "create user")
		result.User, err = q.CreateUser(ctx, CreateUserParams{
			AccountID: resultAcc.ID,
			AdminUser: arg.AdminUser,
			Username: arg.Username,
			Password: arg.Password,
			Grade: arg.Grade,
		})
		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}

// DeleteAccountUserTxParams stores all information needed to delete an account
type DeleteAccountUserTxParams struct {
	Email string `json:"email"`
	Username string `json:"username"`
}

// DeleteAccountUserTx updates users user_question_score to default accounts
// and deletes account data given a username and password
func (store *SQLStore) DeleteAccountUserTx(ctx context.Context, arg DeleteAccountUserTxParams) error {
	
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		txName := ctx.Value(txKey)
		log.Println("Beginning transaction to delete account", txName)
		account, err := q.GetAccountByEmail(ctx, arg.Email)
		if err != nil {
			return err
		}
		user, err := q.GetUserByUsername(ctx, arg.Username)
		if err != nil {
			return err
		}
		// This is where we will move data.
		defaultUser, err := q.GetUserByUsername(ctx, string(user.Grade))
		if err != nil {
			return err
		}
		scores, err := q.UpdateUserUserQuestionScore(ctx, UpdateUserUserQuestionScoreParams{
			UserID: user.ID,
			UserID_2: defaultUser.ID,
		})
		if err != sql.ErrNoRows {
			return err
		}
		log.Printf("Total scores updated: %v from %s to %s", len(scores), user.Username, defaultUser.Username)
		err = q.DeleteUser(ctx, user.ID)
		if err != nil {
			return err
		}
		err = q.DeleteAccount(ctx, account.ID)
		if err != nil {
			return err
		}
		log.Println("Completed transaction to delete account", txName)
		return nil
	})
	return err
}
// question creation (requires category selection) -> question_category records

// delete operations related to a question should be rolled into a transaction

// delete question operations should be done in tx

// delete category ""

// delete account ""

// insert new user_question_score can set the old score to false
// and then add the new one
// update question score can set the current question is_most_recent to false
// and then 