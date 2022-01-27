-- name: CreateQuestionCategory :one
INSERT INTO question_categories (
    question_id,
    category_id
) VALUES ( 
    $1, $2 
) RETURNING *;

