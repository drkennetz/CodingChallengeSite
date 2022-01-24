// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetAUser(ctx context.Context, id int64) (User, error)
	GetAccount(ctx context.Context, id int64) (Account, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateAccountEmail(ctx context.Context, arg UpdateAccountEmailParams) (Account, error)
	UpdateAccountName(ctx context.Context, arg UpdateAccountNameParams) (Account, error)
	UpdateAccountOptedIn(ctx context.Context, arg UpdateAccountOptedInParams) (Account, error)
	UpdateAdminStatus(ctx context.Context, arg UpdateAdminStatusParams) (User, error)
	UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (User, error)
	UpdateUserGrade(ctx context.Context, arg UpdateUserGradeParams) (User, error)
	UpdateUsername(ctx context.Context, arg UpdateUsernameParams) (User, error)
}

var _ Querier = (*Queries)(nil)