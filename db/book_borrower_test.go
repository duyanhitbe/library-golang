package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomBookBorrower(t *testing.T) *BookBorrower {
	book := createRandomBook(t)
	borrower := createRandomBorrower(t)
	bookBorrower, err := testQueries.CreateBookBorrower(context.Background(), CreateBookBorrowerParams{
		BorrowerID: borrower.ID,
		BookID:     book.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, bookBorrower)
	require.NotZero(t, bookBorrower.BookID)
	require.NotZero(t, bookBorrower.BorrowerID)
	return bookBorrower
}

func TestQueries_CreateBookBorrower(t *testing.T) {
	removeAll()
	createRandomBookBorrower(t)
}

func TestQueries_GetAllBookBorrowerByBookId(t *testing.T) {
	removeAll()
	bookBorrower := createRandomBookBorrower(t)

	bookBorrowers, err := testQueries.GetAllBookBorrowerByBookId(context.Background(), bookBorrower.BookID)
	require.NoError(t, err)
	require.NotEmpty(t, bookBorrowers[0])
	require.NotZero(t, bookBorrowers[0].BorrowerID)
	require.NotZero(t, bookBorrowers[0].BookID)
	require.Equal(t, bookBorrowers[0].BookID, bookBorrower.BookID)
}

func TestQueries_GetAllBookBorrowerByBorrowerId(t *testing.T) {
	removeAll()
	bookBorrower := createRandomBookBorrower(t)

	bookBorrowers, err := testQueries.GetAllBookBorrowerByBorrowerId(context.Background(), bookBorrower.BorrowerID)
	require.NoError(t, err)
	require.NotEmpty(t, bookBorrowers[0])
	require.NotZero(t, bookBorrowers[0].BorrowerID)
	require.NotZero(t, bookBorrowers[0].BookID)
	require.Equal(t, bookBorrowers[0].BorrowerID, bookBorrower.BorrowerID)
}

func TestQueries_GetOneBookBorrower(t *testing.T) {
	removeAll()
	bookBorrower := createRandomBookBorrower(t)

	findOneBookBorrower, err := testQueries.GetOneBookBorrower(context.Background(), GetOneBookBorrowerParams{
		BorrowerID: bookBorrower.BorrowerID,
		BookID:     bookBorrower.BookID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, findOneBookBorrower)
	require.NotZero(t, findOneBookBorrower.BorrowerID)
	require.NotZero(t, findOneBookBorrower.BookID)
}
