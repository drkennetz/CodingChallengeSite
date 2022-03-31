package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	//"log"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	CreateAccountUserTx(ctx context.Context, arg CreateAccountUserTxParams) (CreateAccountUserTxResult, error)
	DeleteAccountUserTx(ctx context.Context, arg DeleteAccountUserTxParams) (error)
	CreateQuestionCategoryTx(ctx context.Context, arg CreateQuestionCategoryTxParams) (CreateQuestionCategoryTxResult, error)
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

		log.Println(txName, "create account")
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
		log.Println(txName, "make sure account exists in db by retrieving it and locking it until user is created")
		resultAcc, err := q.GetAccount(ctx, result.Account.ID)
		if err != nil {
			return err
		}
		log.Println(txName, "create user")
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
		fmt.Println("Beginning transaction to delete account", txName)
		account, err := q.GetAccountByEmail(ctx, arg.Email)
		if err != nil {
			return err
		}
		fmt.Println(account.Email, " Retrieved.")
		user, err := q.GetUserByUsername(ctx, arg.Username)
		if err != nil {
			return err
		}
		fmt.Println(user.Username, "Retrieved.")
		// This is where we will move data.
		defaultUser, err := q.GetUserByUsername(ctx, string(user.Grade))
		if err != nil {
			return err
		}
		fmt.Println(defaultUser.Username, "Retrieved")
		// need to see if user has any questions with listalluserquestionscores
		// should paginate only a few results
		hasScores, err := q.ListAllQuestionScoresByUser(ctx, ListAllQuestionScoresByUserParams{
			UserID: user.ID,
		})
		fmt.Println("Recovered scores")
		if err == sql.ErrNoRows {
			fmt.Printf("user %s returned %v scores... continuing", user.Username, len(hasScores))
		} else if err != nil {
			return err
		} else {
			scores, err := q.UpdateUserUserQuestionScore(ctx, UpdateUserUserQuestionScoreParams{
				UserID: user.ID,
				UserID_2: defaultUser.ID,
			})
			if err != nil {
				return err
			}
			fmt.Printf("Total scores updated: %v from %s to %s", len(scores), user.Username, defaultUser.Username)
		} 
		
		fmt.Println("successfully updated or skipped scores, deleting user and account...")
		err = q.DeleteUser(ctx, user.ID)
		if err != nil {
			return err
		}
		fmt.Println(user.Username, "deleted.")
		err = q.DeleteAccount(ctx, account.ID)
		if err != nil {
			return err
		}
		fmt.Println(account.Email, "deleted.")
		fmt.Println("Completed transaction to delete account", txName)
		return nil
	})
	return err
}
// question creation (requires category selection) -> question_category records

// CreateQuestionCategoryTxParams stores all the information needed to insert a question, and its relevant category
type CreateQuestionCategoryTxParams struct {
	ChallengeName string `json:"challenge_name"`
	Description string `json:"description"`
	Example string `json:"example"`
	Difficulty int32 `json:"difficulty"`
	Complexity string `json:"complexity"`
	CompletionTime int32 `json:"completion_time"`
	QuestionType QuestionType `json:"question_type"`
	Category string `json:"category"`
}

// CreateQuestionCategoryTxResult stores the results of the transaction
type CreateQuestionCategoryTxResult struct {
	Question Question
	QuestionCategory QuestionCategory
}

// CreateQuestionCategoryTx creates an Question and a QuestionCategory record in a transaction, since a question should always be assigned a category
func (store *SQLStore) CreateQuestionCategoryTx(ctx context.Context, arg CreateQuestionCategoryTxParams) (CreateQuestionCategoryTxResult, error) {
	var result CreateQuestionCategoryTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		txName := ctx.Value(txKey)

		// we should first see if a category exists
		log.Println(txName, "Get Category ID for Question Category Creation")
		category, err := q.GetACategoryIDByName(ctx, arg.Category)
		if err != nil {
			return err
		}
		log.Println(txName, "create Question and QuestionCategory")
		result.Question, err = q.CreateQuestion(ctx, CreateQuestionParams{
			ChallengeName: arg.ChallengeName,
			Description: arg.Description,
			Example: arg.Example,
			Difficulty: arg.Difficulty,
			Complexity: arg.Complexity,
			CompletionTime: arg.CompletionTime,
			QuestionType: arg.QuestionType,
		})
		if err != nil {
			return err
		}
	
		log.Println(txName, "Create QuestionCategory")
		result.QuestionCategory, err = q.CreateQuestionCategory(ctx, CreateQuestionCategoryParams{
			QuestionID: result.Question.ID,
			CategoryID: category.ID,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}

// insert many question test cases in single transaction

// insert new user_question_score can set the old score to false and then add the new one
