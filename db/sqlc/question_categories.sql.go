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

const listAllQuestionsByCategory = `-- name: ListAllQuestionsByCategory :many
SELECT id, challenge_name, description, example, difficulty, complexity, completion_time, question_type, created_at, updated_at from questions
where id in (
    SELECT question_id from question_categories where category_id = $1
) ORDER BY ASCENDING(difficulty)
`

func (q *Queries) ListAllQuestionsByCategory(ctx context.Context, categoryID int64) ([]Question, error) {
	rows, err := q.db.QueryContext(ctx, listAllQuestionsByCategory, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Question{}
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.ID,
			&i.ChallengeName,
			&i.Description,
			&i.Example,
			&i.Difficulty,
			&i.Complexity,
			&i.CompletionTime,
			&i.QuestionType,
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