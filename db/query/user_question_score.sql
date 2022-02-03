-- name: CreateUserQuestionScore :one
INSERT INTO user_question_score (
    user_id,
    question_id,
    score
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetLastUserQuestionScore :one
SELECT * from user_question_score
where user_id = $1 and question_id = $2 and is_most_recent = true;

-- name: ListQuestionScoresByUser :many
SELECT * from user_question_score
where user_id = $1 and question_id = $2
ORDER BY DESCENDING(id)
LIMIT $3
OFFSET $4;

-- name: ListLastScoresByQuestion :many
SELECT * from user_question_score
where question_id = $1 and is_most_recent = true
LIMIT $2
OFFSET $3;

-- name: UpdateUserQuestionScore :one
UPDATE user_question_score
SET is_most_recent = $3
where user_id = $1 and question_id = $2
RETURNING *;

-- name: DeleteUserQuestionScore :exec
DELETE from user_question_score
where id = $1;