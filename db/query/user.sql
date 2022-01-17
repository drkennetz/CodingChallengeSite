-- name: CreateUser :one
INSERT INTO users (
    account_id,
    admin_user,
    username,
    password,
    grade
) VALUES (
  $1, $2, $3, $4, $5  
)  RETURNING *;

-- name: GetAUser :one
SELECT * FROM users
where id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * from users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUsername :one
UPDATE users
SET username = $2
where id = $1
RETURNING *;

-- name: UpdatePassword :one
UPDATE users
SET password = $2
where username = $1
RETURNING *;

-- name: UpdateAdminStatus :one
UPDATE users
SET admin_user = $2
where id = $1
RETURNING *;

-- name: UpdateUserGrade :one
UPDATE users
SET grade = $2
where id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE from users
where id = $1;