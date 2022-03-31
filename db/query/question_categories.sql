-- name: CreateQuestionCategory :one
INSERT INTO question_categories (
    question_id,
    category_id
) VALUES ( 
    $1, $2 
) RETURNING *;

-- name: ListAllQuestionsByCategory :many
SELECT * from questions
where id in (
    SELECT question_id from question_categories where category_id = $1
) ORDER BY difficulty ASC;

-- name: DeleteQuestionCategory :exec
DELETE from question_categories
where id = $1;