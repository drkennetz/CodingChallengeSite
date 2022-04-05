package db

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestQuestionTestCase(t *testing.T) (Question, QuestionTestCase){
	q := createRandomQuestion(t, 1000000)
	
	input := []byte(`{"test": 1}`)
	output := []byte(`{"test": 2}`)
	arg := CreateQuestionTestCaseParams {
		QuestionID: q.ID,
		Inputs: json.RawMessage(input),
		Outputs: json.RawMessage(output),
	}
	qtc, err := testQueries.CreateQuestionTestCase(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, qtc.QuestionID, arg.QuestionID)
	return q, qtc
}

func TestGetAllQuestionTestCases(t *testing.T) {
	q, qtc := createTestQuestionTestCase(t)
	require.NotEmpty(t, qtc)
	qs, err := testQueries.GetAllQuestionTestCases(context.Background(), q.ID)
	require.NoError(t, err)
	require.Equal(t, qs[0].QuestionID, q.ID)
	err = testQueries.DeleteAllQuestionTestCasesForQuestion(context.Background(), q.ID)
	require.NoError(t, err)
	err = testQueries.DeleteQuestion(context.Background(), q.ID)
	require.NoError(t, err)
}

func TestDeleteAllQuestionTestCasesForQuestion(t *testing.T) {
	q, _ := createTestQuestionTestCase(t)
	err := testQueries.DeleteAllQuestionTestCasesForQuestion(context.Background(), q.ID)
	require.NoError(t, err)
	err = testQueries.DeleteQuestion(context.Background(), q.ID)
	require.NoError(t, err)
}

func TestDeleteOneQuestionTestCase(t *testing.T) {
	q, qtc := createTestQuestionTestCase(t)
	err := testQueries.DeleteOneQuestionTestCase(context.Background(), qtc.ID)
	require.NoError(t, err)
	err = testQueries.DeleteQuestion(context.Background(), q.ID)
	require.NoError(t, err)
}

func TestGetSomeQuestionTestCases(t *testing.T) {
	q, _ := createTestQuestionTestCase(t)
	arg := GetSomeQuestionTestCasesParams {
		q.ID,
		1,
		0,
	}
	qtcs, err := testQueries.GetSomeQuestionTestCases(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 1, len(qtcs))
	err = testQueries.DeleteAllQuestionTestCasesForQuestion(context.Background(), q.ID)
	require.NoError(t, err)
	err = testQueries.DeleteQuestion(context.Background(), q.ID)
	require.NoError(t, err)
}

func TestUpdateOneQuestionTestCase(t *testing.T) {
	q, qtc := createTestQuestionTestCase(t)
	input := []byte(`{"test1": 2}`)
	output := []byte(`{"test1": 2}`)
	arg := UpdateOneQuestionTestCaseParams {
		ID: qtc.ID,
		QuestionID: q.ID,
		Inputs: json.RawMessage(input),
		Outputs: json.RawMessage(output),
	}
	uqtc, err := testQueries.UpdateOneQuestionTestCase(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Inputs, uqtc.Inputs)
	require.Equal(t, arg.QuestionID, uqtc.QuestionID)
	require.Equal(t, arg.Outputs, uqtc.Outputs)
	require.Equal(t, arg.ID, uqtc.ID)
	err = testQueries.DeleteAllQuestionTestCasesForQuestion(context.Background(), q.ID)
	require.NoError(t, err)
	err = testQueries.DeleteQuestion(context.Background(), q.ID)
	require.NoError(t, err)
}