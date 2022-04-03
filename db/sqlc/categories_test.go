package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestCategory(t *testing.T) Category {
	category, err := testQueries.CreateCategory(context.Background(), "test")
	require.NoError(t, err)
	require.Equal(t, "test", category.Category)
	return category
}

func TestGetACategoryByID(t *testing.T) {
	cat := createTestCategory(t)
	cat2, err := testQueries.GetACategoryByID(context.Background(), cat.ID)
	require.NoError(t, err)
	require.Equal(t, cat.ID, cat2.ID)
	require.Equal(t, cat.Category, cat2.Category)
	err = testQueries.DeleteCategory(context.Background(), cat.ID)
	require.NoError(t, err)
}

func TestGetACategoryIDByName(t *testing.T) {
	cat := createTestCategory(t)
	cat2, err := testQueries.GetACategoryIDByName(context.Background(), cat.Category)
	require.NoError(t, err)
	require.Equal(t, cat.ID, cat2.ID)
	require.Equal(t, cat.Category, cat2.Category)
	err = testQueries.DeleteCategory(context.Background(), cat.ID)
	require.NoError(t, err)
}

func TestListCategories(t *testing.T) {
	arg := ListCategoriesParams {
		Limit: 3,
		Offset: 0,
	}
	cats, err := testQueries.ListCategories(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 3, len(cats))
}

func TestUpdateCategory(t *testing.T) {
	cat := createTestCategory(t)
	arg := UpdateCategoryParams {
		ID: cat.ID,
		Category: "test2",
	}
	newcat, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, "test2", newcat.Category)
	require.Equal(t, cat.ID, newcat.ID)
	err = testQueries.DeleteCategory(context.Background(), cat.ID)
	require.NoError(t, err)
}

