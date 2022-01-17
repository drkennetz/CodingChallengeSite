package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
// extending the functionality of queries since one query only handles one operation
type Store struct {
	*Queries // composition - all Queries will be available to Store
	db *sql.DB
}

// NewStore builds a new Store object and returns it
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
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

func (store *Store) CreateAccountUserTx(ctx context.Context, arg CreateAccountUserTxParams) (CreateAccountUserTxResult, error) {
	var result CreateAccountUserTxResult

	// we are accessing the result variable from inside the tx function. similar to the arg variable. This makes the callback fn a closure
	// sin golangs support for generics type, closure is often used when we get the result from a callback function
	// this is because callback function itself doesn't know the exact type of the result
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Account, err = q.CreateAccount(ctx, CreateAccountParams{
			FullName: arg.FullName,
			Email: arg.Email,
			OptedIn: arg.OptedIn,
		})
		if err != nil {
			return err
		}

		// might have to try to prevent deadlock here since using a value generated
		// by db in step above. If DB takes longer to generate the account ID than
		// it does for code to get here (which is possible), we could run into
		// race conditions
		result.User, err = q.CreateUser(ctx, CreateUserParams{
			AccountID: result.Account.ID,
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
// question creation (requires category selection) -> question_category records

// delete operations related to a question should be rolled into a transaction

// delete question operations should be done in tx

// delete category ""

// delete account ""
