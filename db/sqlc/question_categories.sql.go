// Code generated by sqlc. DO NOT EDIT.
// source: question_categories.sql

package db

import (
	"context"
)

const createQuestionCategory = `-- name: CreateQuestionCategory :one
INSERT INTO question_categories (
    question_id,
    category_id
) VALUES ( 
    $1, $2 
) RETURNING id, question_id, category_id
`

type CreateQuestionCategoryParams struct {
	QuestionID int64 `json:"question_id"`
	CategoryID int64 `json:"category_id"`
}

func (q *Queries) CreateQuestionCategory(ctx context.Context, arg CreateQuestionCategoryParams) (QuestionCategory, error) {
	row := q.db.QueryRowContext(ctx, createQuestionCategory, arg.QuestionID, arg.CategoryID)
	var i QuestionCategory
	err := row.Scan(&i.ID, &i.QuestionID, &i.CategoryID)
	return i, err
}
