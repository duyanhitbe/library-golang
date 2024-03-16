package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomBook(t *testing.T) *Book {
	category := createRandomCategory(t)
	bookInfo := createRandomBookInfo(t)
	book, err := testQueries.CreateBook(context.Background(), CreateBookParams{
		CategoryID: category.ID,
		BookInfoID: bookInfo.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, book)
	require.NotZero(t, book.ID)
	require.NotZero(t, book.CategoryID)
	require.NotZero(t, book.BookInfoID)
	require.NotZero(t, book.CreatedAt)
	require.NotZero(t, book.UpdatedAt)
	require.NotZero(t, book.IsActive, true)
	return book
}

func TestQueries_CountBook(t *testing.T) {
	removeAll()
	createRandomBook(t)
	count, err := testQueries.CountBook(context.Background())
	require.NoError(t, err)
	require.Equal(t, count, int64(1))
}

func TestQueries_CreateBook(t *testing.T) {
	removeAll()
	createRandomBook(t)
}

func TestQueries_DeleteOneBookById(t *testing.T) {
	removeAll()
	book := createRandomBook(t)

	deletedBook, err := testQueries.DeleteOneBookById(context.Background(), book.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedBook)
	require.NotZero(t, deletedBook.ID)
	require.NotZero(t, deletedBook.CategoryID)
	require.NotZero(t, deletedBook.BookInfoID)
	require.NotZero(t, deletedBook.CreatedAt)
	require.NotZero(t, deletedBook.UpdatedAt)
	require.NotZero(t, deletedBook.IsActive, true)
}

func TestQueries_GetOneBookById(t *testing.T) {
	removeAll()
	book := createRandomBook(t)

	findOneBook, err := testQueries.GetOneBookById(context.Background(), book.ID)
	require.NoError(t, err)
	require.NotEmpty(t, findOneBook)
	require.NotZero(t, findOneBook.ID)
	require.NotZero(t, findOneBook.CategoryID)
	require.NotZero(t, findOneBook.BookInfoID)
	require.NotZero(t, findOneBook.CreatedAt)
	require.NotZero(t, findOneBook.UpdatedAt)
	require.NotZero(t, findOneBook.IsActive, true)
}

func TestQueries_ListBook(t *testing.T) {
	removeAll()
	createRandomBook(t)
	createRandomBook(t)

	categories, err := testQueries.ListBook(context.Background(), ListBookParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	for _, book := range categories {
		require.NotEmpty(t, book)
		require.NotZero(t, book.ID)
		require.NotZero(t, book.CategoryID)
		require.NotZero(t, book.BookInfoID)
		require.NotZero(t, book.CreatedAt)
		require.NotZero(t, book.UpdatedAt)
		require.NotZero(t, book.IsActive, true)
	}
}

func TestQueries_UpdateOneBookById(t *testing.T) {
	removeAll()
	book := createRandomBook(t)
	category := createRandomCategory(t)
	updatedBook, err := testQueries.UpdateOneBookById(context.Background(), UpdateOneBookByIdParams{
		ID:         book.ID,
		CategoryID: category.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedBook)
	require.NotZero(t, updatedBook.ID)
	require.NotZero(t, updatedBook.CategoryID)
	require.NotZero(t, updatedBook.BookInfoID)
	require.Equal(t, updatedBook.CategoryID, category.ID)
	require.NotZero(t, updatedBook.CreatedAt)
	require.NotZero(t, updatedBook.UpdatedAt)
	require.NotZero(t, updatedBook.IsActive, true)
}
