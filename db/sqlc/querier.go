// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateCategory(ctx context.Context, category string) (Category, error)
	CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error)
	CreateQuestionCategory(ctx context.Context, arg CreateQuestionCategoryParams) (QuestionCategory, error)
	CreateQuestionTestCase(ctx context.Context, arg CreateQuestionTestCaseParams) (QuestionTestCase, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserQuestionScore(ctx context.Context, arg CreateUserQuestionScoreParams) (UserQuestionScore, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteAllQuestionTestCasesForQuestion(ctx context.Context, questionID int64) error
	DeleteCategory(ctx context.Context, id int64) error
	DeleteOneQuestionTestCase(ctx context.Context, id int64) error
	DeleteQuestion(ctx context.Context, id int64) error
	DeleteQuestionCategory(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	DeleteUserQuestionScore(ctx context.Context, id int64) error
	GetACategoryByID(ctx context.Context, id int64) (Category, error)
	GetACategoryIDByName(ctx context.Context, category string) (Category, error)
	GetAUser(ctx context.Context, id int64) (User, error)
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountByEmail(ctx context.Context, email string) (Account, error)
	GetAllQuestionTestCases(ctx context.Context, questionID int64) ([]QuestionTestCase, error)
	GetLastUserQuestionScore(ctx context.Context, arg GetLastUserQuestionScoreParams) (UserQuestionScore, error)
	GetQuestion(ctx context.Context, id int64) (Question, error)
	GetSomeQuestionTestCases(ctx context.Context, arg GetSomeQuestionTestCasesParams) ([]QuestionTestCase, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListAllQuestionScoresByUser(ctx context.Context, arg ListAllQuestionScoresByUserParams) ([]UserQuestionScore, error)
	ListAllQuestions(ctx context.Context, arg ListAllQuestionsParams) ([]Question, error)
	ListAllQuestionsByCategory(ctx context.Context, categoryID int64) ([]Question, error)
	ListAllQuestionsByDifficulty(ctx context.Context, arg ListAllQuestionsByDifficultyParams) ([]Question, error)
	ListAllQuestionsByType(ctx context.Context, arg ListAllQuestionsByTypeParams) ([]Question, error)
	ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error)
	ListLastScoresByQuestion(ctx context.Context, arg ListLastScoresByQuestionParams) ([]UserQuestionScore, error)
	ListSingleQuestionScoresByUser(ctx context.Context, arg ListSingleQuestionScoresByUserParams) ([]UserQuestionScore, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateAccountEmail(ctx context.Context, arg UpdateAccountEmailParams) (Account, error)
	UpdateAccountName(ctx context.Context, arg UpdateAccountNameParams) (Account, error)
	UpdateAccountOptedIn(ctx context.Context, arg UpdateAccountOptedInParams) (Account, error)
	UpdateAdminStatus(ctx context.Context, arg UpdateAdminStatusParams) (User, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateLatestUserQuestionScore(ctx context.Context, arg UpdateLatestUserQuestionScoreParams) (UserQuestionScore, error)
	UpdateOneQuestionTestCase(ctx context.Context, arg UpdateOneQuestionTestCaseParams) (QuestionTestCase, error)
	UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (User, error)
	UpdateQuestion(ctx context.Context, arg UpdateQuestionParams) (Question, error)
	UpdateUserGrade(ctx context.Context, arg UpdateUserGradeParams) (User, error)
	UpdateUserUserQuestionScore(ctx context.Context, arg UpdateUserUserQuestionScoreParams) ([]UserQuestionScore, error)
	UpdateUsername(ctx context.Context, arg UpdateUsernameParams) (User, error)
}

var _ Querier = (*Queries)(nil)
