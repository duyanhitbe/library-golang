package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomCategory(t *testing.T) *Category {
	category, err := testQueries.CreateCategory(context.Background(), "category")
	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.NotZero(t, category.ID)
	require.NotZero(t, category.Name)
	require.NotZero(t, category.CreatedAt)
	require.NotZero(t, category.UpdatedAt)
	require.NotZero(t, category.IsActive, true)
	return category
}

func TestQueries_CountCategory(t *testing.T) {
	removeAll()
	createRandomCategory(t)
	count, err := testQueries.CountCategory(context.Background())
	require.NoError(t, err)
	require.Equal(t, count, int64(1))
}

func TestQueries_CreateCategory(t *testing.T) {
	removeAll()
	createRandomCategory(t)
}

func TestQueries_DeleteOneCategoryById(t *testing.T) {
	removeAll()
	category := createRandomCategory(t)

	deletedCategory, err := testQueries.DeleteOneCategoryById(context.Background(), category.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedCategory)
	require.NotZero(t, deletedCategory.ID)
	require.NotZero(t, deletedCategory.Name)
	require.NotZero(t, deletedCategory.CreatedAt)
	require.NotZero(t, deletedCategory.UpdatedAt)
	require.NotZero(t, deletedCategory.IsActive, true)
}

func TestQueries_GetOneCategoryById(t *testing.T) {
	removeAll()
	category := createRandomCategory(t)

	findOneCategory, err := testQueries.GetOneCategoryById(context.Background(), category.ID)
	require.NoError(t, err)
	require.NotEmpty(t, findOneCategory)
	require.NotZero(t, findOneCategory.ID)
	require.NotZero(t, findOneCategory.Name)
	require.NotZero(t, findOneCategory.CreatedAt)
	require.NotZero(t, findOneCategory.UpdatedAt)
	require.NotZero(t, findOneCategory.IsActive, true)
}

func TestQueries_ListCategory(t *testing.T) {
	removeAll()
	createRandomCategory(t)
	createRandomCategory(t)

	categories, err := testQueries.ListCategory(context.Background(), ListCategoryParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	for _, category := range categories {
		require.NotEmpty(t, category)
		require.NotZero(t, category.ID)
		require.NotZero(t, category.Name)
		require.NotZero(t, category.CreatedAt)
		require.NotZero(t, category.UpdatedAt)
		require.NotZero(t, category.IsActive, true)
	}
}

func TestQueries_UpdateOneCategoryById(t *testing.T) {
	removeAll()
	category := createRandomCategory(t)

	updatedCategory, err := testQueries.UpdateOneCategoryById(context.Background(), UpdateOneCategoryByIdParams{
		ID:   category.ID,
		Name: "Updated",
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedCategory)
	require.NotZero(t, updatedCategory.ID)
	require.NotZero(t, updatedCategory.Name)
	require.Equal(t, updatedCategory.Name, "Updated")
	require.NotZero(t, updatedCategory.CreatedAt)
	require.NotZero(t, updatedCategory.UpdatedAt)
	require.NotZero(t, updatedCategory.IsActive, true)
}
