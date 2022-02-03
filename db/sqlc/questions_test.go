package db

import (
	"context"
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

	// tests that auto populated field were populated
	require.NotZero(t, question.ID)
	require.NotZero(t, question.CreatedAt)
	require.NotZero(t, question.UpdatedAt)

	return question
}

func TestCreateQuestion(t *testing.T) {
	createRandomQuestion(t, 1)
}