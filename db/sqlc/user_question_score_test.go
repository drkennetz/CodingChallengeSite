package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createTestUserQuestionScore(t *testing.T) UserQuestionScore{
	arg := CreateUserQuestionScoreParams{
		UserID:       1,
		QuestionID:   1,
		Score:        1000,
		IsMostRecent: sql.NullBool{
			Bool: true,
			Valid: true,
		},	
	}
	uqs, err := testQueries.CreateUserQuestionScore(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, uqs.QuestionID, arg.QuestionID)
	require.Equal(t, uqs.UserID, arg.UserID)
	require.Equal(t, uqs.Score, arg.Score)
	require.True(t, uqs.IsMostRecent.Bool)
	require.WithinDuration(t, uqs.CreatedAt, time.Now(), time.Minute)
	require.WithinDuration(t, uqs.UpdatedAt, time.Now(), time.Minute)
	require.NotZero(t, uqs.ID)
	return uqs
}

func TestDeleteUserQuestionScore(t *testing.T) {
	uqs := createTestUserQuestionScore(t)
	err := testQueries.DeleteUserQuestionScore(context.Background(), uqs.ID)
	require.NoError(t, err)
}

func TestGetLastUserQuestionScore(t *testing.T) {
	uqs := createTestUserQuestionScore(t)
	arg := GetLastUserQuestionScoreParams {
		UserID: uqs.UserID,
		QuestionID: uqs.QuestionID,
	}
	last, err := testQueries.GetLastUserQuestionScore(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, uqs.QuestionID, last.QuestionID)
	require.Equal(t, uqs.UserID, last.UserID)
	require.Equal(t, uqs.Score, last.Score)
	require.True(t, last.IsMostRecent.Bool)
	require.WithinDuration(t, last.CreatedAt, time.Now(), time.Minute)
	require.WithinDuration(t, last.UpdatedAt, time.Now(), time.Minute)
	require.Equal(t, last.ID, uqs.ID)

	err = testQueries.DeleteUserQuestionScore(context.Background(), uqs.ID)
	require.NoError(t, err)
}