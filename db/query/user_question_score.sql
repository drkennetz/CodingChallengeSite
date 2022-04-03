-- name: CreateUserQuestionScore :one
INSERT INTO user_question_score (
    user_id,
    question_id,
    score,
    is_most_recent
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetLastUserQuestionScore :one
SELECT * from user_question_score
where user_id = $1 and question_id = $2 and is_most_recent = true LIMIT 1;

-- name: ListSingleQuestionScoresByUser :many
SELECT * from user_question_score
where user_id = $1 and question_id = $2
ORDER BY score DESC
LIMIT $3
OFFSET $4;

-- name: ListAllQuestionScoresByUser :many
SELECT * from user_question_score
where user_id = $1
ORDER BY score DESC
LIMIT $2
OFFSET $3;

-- name: ListLastScoresByQuestion :many
SELECT * from user_question_score
where question_id = $1 and is_most_recent = true
LIMIT $2
OFFSET $3;

-- name: UpdateUserUserQuestionScore :many
UPDATE user_question_score
SET user_id = $2
where user_id = $1
RETURNING *;

-- name: UpdateLatestUserQuestionScore :one
UPDATE user_question_score
SET is_most_recent = $3
where user_id = $1 and question_id = $2
RETURNING *;

-- name: DeleteUserQuestionScore :exec
DELETE from user_question_score
where id = $1;