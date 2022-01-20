// Code generated by sqlc. DO NOT EDIT.
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
  full_name,
  email,
  opted_in
) VALUES (
  $1, $2, $3
) RETURNING id, full_name, email, opted_in, created_at, updated_at
`

type CreateAccountParams struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	OptedIn  bool   `json:"opted_in"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.FullName, arg.Email, arg.OptedIn)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.OptedIn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
where id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, full_name, email, opted_in, created_at, updated_at FROM accounts
where id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.OptedIn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, full_name, email, opted_in, created_at, updated_at FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.Email,
			&i.OptedIn,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccountEmail = `-- name: UpdateAccountEmail :one
UPDATE accounts
SET email = $2
where id = $1
RETURNING id, full_name, email, opted_in, created_at, updated_at
`

type UpdateAccountEmailParams struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func (q *Queries) UpdateAccountEmail(ctx context.Context, arg UpdateAccountEmailParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountEmail, arg.ID, arg.Email)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.OptedIn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateAccountName = `-- name: UpdateAccountName :one
UPDATE accounts
SET full_name = $2
where id = $1
RETURNING id, full_name, email, opted_in, created_at, updated_at
`

type UpdateAccountNameParams struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
}

func (q *Queries) UpdateAccountName(ctx context.Context, arg UpdateAccountNameParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountName, arg.ID, arg.FullName)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.OptedIn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateAccountOptedIn = `-- name: UpdateAccountOptedIn :one
UPDATE accounts
SET opted_in = $2
where id = $1
RETURNING id, full_name, email, opted_in, created_at, updated_at
`

type UpdateAccountOptedInParams struct {
	ID      int64 `json:"id"`
	OptedIn bool  `json:"opted_in"`
}

func (q *Queries) UpdateAccountOptedIn(ctx context.Context, arg UpdateAccountOptedInParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountOptedIn, arg.ID, arg.OptedIn)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.OptedIn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
