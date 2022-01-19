-- name: CreateAccount :one
INSERT INTO accounts (
  full_name,
  email,
  opted_in
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
where id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
where id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccountName :one
UPDATE accounts
SET full_name = $2
where id = $1
RETURNING *;

-- name: UpdateAccountEmail :one
UPDATE accounts
SET email = $2
where id = $1
RETURNING *;

-- name: UpdateAccountOptedIn :one
UPDATE accounts
SET opted_in = $2
where id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
where id = $1;