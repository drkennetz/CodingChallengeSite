-- name: CreateQuestionTestCase :one
INSERT INTO question_test_cases (
    question_id,
    inputs,
    outputs
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetSomeQuestionTestCases :many
SELECT * from question_test_cases
where question_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: GetAllQuestionTestCases :many
SELECT * from question_test_cases
where question_id = $1
ORDER BY id;

-- name: UpdateOneQuestionTestCase :one
UPDATE question_test_cases
SET inputs = $3, outputs = $4
where id = $1 and question_id = $2
RETURNING *;

-- name: DeleteOneQuestionTestCase :exec
DELETE from question_test_cases
where id = $1;

-- name: DeleteAllQuestionTestCasesForQuestion :exec
DELETE from question_test_cases
where question_id = $1;
