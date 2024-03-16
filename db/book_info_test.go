package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomBookInfo(t *testing.T) *BookInfo {
	bookInfo, err := testQueries.CreateBookInfo(context.Background(), CreateBookInfoParams{
		Name:            "book info name",
		Author:          "book info author",
		PublicationDate: time.Now(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, bookInfo)
	require.NotZero(t, bookInfo.ID)
	require.NotZero(t, bookInfo.Name)
	require.NotZero(t, bookInfo.CreatedAt)
	require.NotZero(t, bookInfo.UpdatedAt)
	require.NotZero(t, bookInfo.IsActive, true)
	return bookInfo
}

func TestQueries_CountBookInfo(t *testing.T) {
	removeAll()
	createRandomBookInfo(t)
	count, err := testQueries.CountBookInfo(context.Background())
	require.NoError(t, err)
	require.Equal(t, count, int64(1))
}

func TestQueries_CreateBookInfo(t *testing.T) {
	removeAll()
	createRandomBookInfo(t)
}

func TestQueries_DeleteOneBookInfoById(t *testing.T) {
	removeAll()
	bookInfo := createRandomBookInfo(t)

	deletedBookInfo, err := testQueries.DeleteOneBookInfoById(context.Background(), bookInfo.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedBookInfo)
	require.NotZero(t, deletedBookInfo.ID)
	require.NotZero(t, deletedBookInfo.Name)
	require.NotZero(t, deletedBookInfo.Author)
	require.NotZero(t, deletedBookInfo.PublicationDate)
	require.NotZero(t, deletedBookInfo.CreatedAt)
	require.NotZero(t, deletedBookInfo.UpdatedAt)
	require.NotZero(t, deletedBookInfo.IsActive, true)
}

func TestQueries_GetOneBookInfoById(t *testing.T) {
	removeAll()
	bookInfo := createRandomBookInfo(t)

	findOneBookInfo, err := testQueries.GetOneBookInfoById(context.Background(), bookInfo.ID)
	require.NoError(t, err)
	require.NotEmpty(t, findOneBookInfo)
	require.NotZero(t, findOneBookInfo.ID)
	require.NotZero(t, findOneBookInfo.Name)
	require.NotZero(t, findOneBookInfo.Author)
	require.NotZero(t, findOneBookInfo.PublicationDate)
	require.NotZero(t, findOneBookInfo.CreatedAt)
	require.NotZero(t, findOneBookInfo.UpdatedAt)
	require.NotZero(t, findOneBookInfo.IsActive, true)
}

func TestQueries_ListBookInfo(t *testing.T) {
	removeAll()
	createRandomBookInfo(t)
	createRandomBookInfo(t)

	categories, err := testQueries.ListBookInfo(context.Background(), ListBookInfoParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	for _, bookInfo := range categories {
		require.NotEmpty(t, bookInfo)
		require.NotZero(t, bookInfo.ID)
		require.NotZero(t, bookInfo.Name)
		require.NotZero(t, bookInfo.Author)
		require.NotZero(t, bookInfo.PublicationDate)
		require.NotZero(t, bookInfo.CreatedAt)
		require.NotZero(t, bookInfo.UpdatedAt)
		require.NotZero(t, bookInfo.IsActive, true)
	}
}

func TestQueries_UpdateOneBookInfoById(t *testing.T) {
	removeAll()
	bookInfo := createRandomBookInfo(t)
	updatedBookInfo, err := testQueries.UpdateOneBookInfoById(context.Background(), UpdateOneBookInfoByIdParams{
		ID:              bookInfo.ID,
		Name:            "Updated Name",
		Author:          "Updated Author",
		PublicationDate: time.Now(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedBookInfo)
	require.NotZero(t, updatedBookInfo.ID)
	require.NotZero(t, updatedBookInfo.Name)
	require.NotZero(t, updatedBookInfo.Author)
	require.NotZero(t, updatedBookInfo.PublicationDate)
	require.Equal(t, updatedBookInfo.Name, "Updated Name")
	require.Equal(t, updatedBookInfo.Author, "Updated Author")
	require.NotZero(t, updatedBookInfo.CreatedAt)
	require.NotZero(t, updatedBookInfo.UpdatedAt)
	require.NotZero(t, updatedBookInfo.IsActive, true)
}
