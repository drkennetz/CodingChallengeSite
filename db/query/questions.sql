-- name: CreateQuestion :one
INSERT INTO questions (
    challenge_name,
    description,
    example,
    difficulty,
    complexity,
    completion_time,
    question_type
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetQuestion :one
SELECT * from questions
where id = $1 LIMIT 1;

-- name: UpdateQuestion :one
UPDATE questions
SET challenge_name = $2, description = $3,
example = $4, difficulty = $5, complexity = $6, 
completion_time = $7 where id = $1
RETURNING *;

-- name: ListAllQuestions :many
SELECT * from questions
ORDER BY difficulty ASC
LIMIT $1
OFFSET $2;

-- name: ListAllQuestionsByDifficulty :many
SELECT * from questions
where difficulty = $1
ORDER BY difficulty ASC
LIMIT $2
OFFSET $3;

-- name: DeleteQuestion :exec
DELETE FROM questions
where id = $1;

-- name: ListAllQuestionsByType :many
SELECT * from questions
where question_type = $1
ORDER BY difficulty ASC
LIMIT $2
OFFSET $3;