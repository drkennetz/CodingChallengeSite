package db

import (
	"context"
	"testing"
	
	"github.com/stretchr/testify/require"
)

func TestListAllQuestionsByCategory(t *testing.T) {
	qs, err := testQueries.ListAllQuestionsByCategory(context.Background(), 1)
	require.NoError(t, err)
	require.True(t, len(qs) > 0)
}

