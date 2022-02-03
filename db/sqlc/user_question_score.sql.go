// Code generated by sqlc. DO NOT EDIT.
// source: user_question_score.sql

package db

import (
	"context"
	"database/sql"
)

const createUserQuestionScore = `-- name: CreateUserQuestionScore :one
INSERT INTO user_question_score (
    user_id,
    question_id,
    score
) VALUES (
    $1, $2, $3
) RETURNING id, user_id, question_id, score, is_most_recent, created_at, updated_at
`

type CreateUserQuestionScoreParams struct {
	UserID     int64 `json:"user_id"`
	QuestionID int64 `json:"question_id"`
	Score      int32 `json:"score"`
}

func (q *Queries) CreateUserQuestionScore(ctx context.Context, arg CreateUserQuestionScoreParams) (UserQuestionScore, error) {
	row := q.db.QueryRowContext(ctx, createUserQuestionScore, arg.UserID, arg.QuestionID, arg.Score)
	var i UserQuestionScore
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.QuestionID,
		&i.Score,
		&i.IsMostRecent,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUserQuestionScore = `-- name: DeleteUserQuestionScore :exec
DELETE from user_question_score
where id = $1
`

func (q *Queries) DeleteUserQuestionScore(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUserQuestionScore, id)
	return err
}

const getLastUserQuestionScore = `-- name: GetLastUserQuestionScore :one
SELECT id, user_id, question_id, score, is_most_recent, created_at, updated_at from user_question_score
where user_id = $1 and question_id = $2 and is_most_recent = true
`

type GetLastUserQuestionScoreParams struct {
	UserID     int64 `json:"user_id"`
	QuestionID int64 `json:"question_id"`
}

func (q *Queries) GetLastUserQuestionScore(ctx context.Context, arg GetLastUserQuestionScoreParams) (UserQuestionScore, error) {
	row := q.db.QueryRowContext(ctx, getLastUserQuestionScore, arg.UserID, arg.QuestionID)
	var i UserQuestionScore
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.QuestionID,
		&i.Score,
		&i.IsMostRecent,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listLastScoresByQuestion = `-- name: ListLastScoresByQuestion :many
SELECT id, user_id, question_id, score, is_most_recent, created_at, updated_at from user_question_score
where question_id = $1 and is_most_recent = true
LIMIT $2
OFFSET $3
`

type ListLastScoresByQuestionParams struct {
	QuestionID int64 `json:"question_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) ListLastScoresByQuestion(ctx context.Context, arg ListLastScoresByQuestionParams) ([]UserQuestionScore, error) {
	rows, err := q.db.QueryContext(ctx, listLastScoresByQuestion, arg.QuestionID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserQuestionScore{}
	for rows.Next() {
		var i UserQuestionScore
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.QuestionID,
			&i.Score,
			&i.IsMostRecent,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listQuestionScoresByUser = `-- name: ListQuestionScoresByUser :many
SELECT id, user_id, question_id, score, is_most_recent, created_at, updated_at from user_question_score
where user_id = $1 and question_id = $2
ORDER BY DESCENDING(id)
LIMIT $3
OFFSET $4
`

type ListQuestionScoresByUserParams struct {
	UserID     int64 `json:"user_id"`
	QuestionID int64 `json:"question_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) ListQuestionScoresByUser(ctx context.Context, arg ListQuestionScoresByUserParams) ([]UserQuestionScore, error) {
	rows, err := q.db.QueryContext(ctx, listQuestionScoresByUser,
		arg.UserID,
		arg.QuestionID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserQuestionScore{}
	for rows.Next() {
		var i UserQuestionScore
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.QuestionID,
			&i.Score,
			&i.IsMostRecent,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserQuestionScore = `-- name: UpdateUserQuestionScore :one
UPDATE user_question_score
SET is_most_recent = $3
where user_id = $1 and question_id = $2
RETURNING id, user_id, question_id, score, is_most_recent, created_at, updated_at
`

type UpdateUserQuestionScoreParams struct {
	UserID       int64        `json:"user_id"`
	QuestionID   int64        `json:"question_id"`
	IsMostRecent sql.NullBool `json:"is_most_recent"`
}

func (q *Queries) UpdateUserQuestionScore(ctx context.Context, arg UpdateUserQuestionScoreParams) (UserQuestionScore, error) {
	row := q.db.QueryRowContext(ctx, updateUserQuestionScore, arg.UserID, arg.QuestionID, arg.IsMostRecent)
	var i UserQuestionScore
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.QuestionID,
		&i.Score,
		&i.IsMostRecent,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}