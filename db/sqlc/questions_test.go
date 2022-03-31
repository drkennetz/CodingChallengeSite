package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomQuestion(t *testing.T, i int) Question {
	arg := CreateQuestionParams {
		ChallengeName: fmt.Sprintf("TestQuestion%v", i),
		Description: "This is a test question.",
		Example: "Here is what you do.",
		Difficulty: 1,
		Complexity: "O(1) | O(1)",
		CompletionTime: 10,
		QuestionType: QuestionTypePractice,
	}

	question, err := testQueries.CreateQuestion(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, question)

	// tests that arg = expected insertion into db
	require.Equal(t, arg.ChallengeName, question.ChallengeName)
	require.Equal(t, arg.Description, question.Description)
	require.Equal(t, arg.Example, question.Example)
	require.Equal(t, arg.Difficulty, question.Difficulty)
	require.Equal(t, arg.Complexity, question.Complexity)
	require.Equal(t, arg.CompletionTime, question.CompletionTime)
	require.Equal(t, arg.QuestionType, question.QuestionType)

	// tests that auto populated field were populated
	require.NotZero(t, question.ID)
	require.NotZero(t, question.CreatedAt)
	require.NotZero(t, question.UpdatedAt)

	return question
}

func TestCreateQuestion(t *testing.T) {
	q := createRandomQuestion(t, 1)
	err := testQueries.DeleteQuestion(context.Background(), q.ID)
	require.NoError(t, err)
}

func TestGetQuestion(t *testing.T) {
	q, err := testQueries.GetQuestion(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, q)
}

func TestGetQuestionError(t *testing.T) {
	q, err := testQueries.GetQuestion(context.Background(), 1000000)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, q)
}

func TestListAllQuestions(t *testing.T) {
	arg := ListAllQuestionsParams {
		Limit: 3,
		Offset: 0,
	}
	qs, err := testQueries.ListAllQuestions(context.Background(), arg)
	require.Equal(t, 3, len(qs))
	require.NoError(t, err)
}

func TestListAllQuestionsByCategory(t *testing.T) {
	arg := ListAllQuestionsByDifficultyParams {
		Difficulty: 1,
		Limit: 2,
		Offset: 0,
	}
	qs, err := testQueries.ListAllQuestionsByDifficulty(context.Background(), arg)
	require.Equal(t, 2, len(qs))
	require.NoError(t, err)
}

func TestListAllQuestionsByType(t *testing.T) {
	arg := ListAllQuestionsByTypeParams {
		QuestionType: QuestionTypePractice,
		Limit: 3,
		Offset: 0,
	}
	qs, err := testQueries.ListAllQuestionsByType(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(qs), 3)
}

func TestUpdateQuestion(t *testing.T) {
	q := createRandomQuestion(t, 1)
	update := UpdateQuestionParams {
		ID: q.ID,
		ChallengeName: "test",
		Description: q.Description,
		Example: q.Example,
		Difficulty: q.Difficulty,
		Complexity: q.Complexity,
		CompletionTime: q.CompletionTime,
	} 
	
	qs, err := testQueries.UpdateQuestion(context.Background(), update)
	require.NoError(t, err)
	require.Equal(t, "test", qs.ChallengeName)
	err = testQueries.DeleteQuestion(context.Background(), qs.ID)
	require.NoError(t, err)
}