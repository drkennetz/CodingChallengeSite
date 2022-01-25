-- name: CreateCategory :one
INSERT INTO categories (
    category
) VALUES (
    $1
) RETURNING *;

-- name: GetACategoryByID :one
SELECT * FROM categories
where id = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM categories
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCategory :one
UPDATE categories
SET category = $2
where id = $1
RETURNING *;

-- name: DeleteCategory :one
DELETE from categories
where id = $1;
