-- name: CreateUser :one
INSERT INTO users (
    account_id,
    username,
    grade
) VALUES (
  $1, $2, $3  
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

-- name: UpdateUserGrade :one
UPDATE users
SET grade = $2
where id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE from users
where id = $1;