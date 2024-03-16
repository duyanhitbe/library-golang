package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomBorrower(t *testing.T) *Borrower {
	borrower, err := testQueries.CreateBorrower(context.Background(), CreateBorrowerParams{
		Name:    "name",
		Phone:   "000000000",
		Address: "123.st",
	})
	require.NoError(t, err)
	require.NotEmpty(t, borrower)
	require.NotZero(t, borrower.ID)
	require.NotZero(t, borrower.Name)
	require.NotZero(t, borrower.Phone)
	require.NotZero(t, borrower.Address)
	require.NotZero(t, borrower.CreatedAt)
	require.NotZero(t, borrower.UpdatedAt)
	require.NotZero(t, borrower.IsActive, true)
	return borrower
}

func TestQueries_CountBorrower(t *testing.T) {
	removeAll()
	createRandomBorrower(t)
	count, err := testQueries.CountBorrower(context.Background())
	require.NoError(t, err)
	require.Equal(t, count, int64(1))
}

func TestQueries_CreateBorrower(t *testing.T) {
	removeAll()
	createRandomBorrower(t)
}

func TestQueries_DeleteOneBorrowerById(t *testing.T) {
	removeAll()
	borrower := createRandomBorrower(t)

	deletedBorrower, err := testQueries.DeleteOneBorrowerById(context.Background(), borrower.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedBorrower)
	require.NotZero(t, deletedBorrower.ID)
	require.NotZero(t, deletedBorrower.Name)
	require.NotZero(t, deletedBorrower.Phone)
	require.NotZero(t, deletedBorrower.Address)
	require.NotZero(t, deletedBorrower.CreatedAt)
	require.NotZero(t, deletedBorrower.UpdatedAt)
	require.NotZero(t, deletedBorrower.IsActive, true)
}

func TestQueries_GetOneBorrowerById(t *testing.T) {
	removeAll()
	borrower := createRandomBorrower(t)

	findOneBorrower, err := testQueries.GetOneBorrowerById(context.Background(), borrower.ID)
	require.NoError(t, err)
	require.NotEmpty(t, findOneBorrower)
	require.NotZero(t, findOneBorrower.ID)
	require.NotZero(t, findOneBorrower.Name)
	require.NotZero(t, findOneBorrower.Phone)
	require.NotZero(t, findOneBorrower.Address)
	require.NotZero(t, findOneBorrower.CreatedAt)
	require.NotZero(t, findOneBorrower.UpdatedAt)
	require.NotZero(t, findOneBorrower.IsActive, true)
}

func TestQueries_ListBorrower(t *testing.T) {
	removeAll()
	createRandomBorrower(t)
	createRandomBorrower(t)

	categories, err := testQueries.ListBorrower(context.Background(), ListBorrowerParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	for _, borrower := range categories {
		require.NotEmpty(t, borrower)
		require.NotZero(t, borrower.ID)
		require.NotZero(t, borrower.Name)
		require.NotZero(t, borrower.Phone)
		require.NotZero(t, borrower.Address)
		require.NotZero(t, borrower.CreatedAt)
		require.NotZero(t, borrower.UpdatedAt)
		require.NotZero(t, borrower.IsActive, true)
	}
}

func TestQueries_UpdateOneBorrowerById(t *testing.T) {
	removeAll()
	borrower := createRandomBorrower(t)

	updatedBorrower, err := testQueries.UpdateOneBorrowerById(context.Background(), UpdateOneBorrowerByIdParams{
		ID:      borrower.ID,
		Name:    "Updated name",
		Phone:   "0123123",
		Address: "Updated address",
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedBorrower)
	require.NotZero(t, updatedBorrower.ID)
	require.NotZero(t, updatedBorrower.Name)
	require.NotZero(t, updatedBorrower.Phone)
	require.NotZero(t, updatedBorrower.Address)
	require.Equal(t, updatedBorrower.Name, "Updated name")
	require.Equal(t, updatedBorrower.Phone, "0123123")
	require.Equal(t, updatedBorrower.Address, "Updated address")
	require.NotZero(t, updatedBorrower.CreatedAt)
	require.NotZero(t, updatedBorrower.UpdatedAt)
	require.NotZero(t, updatedBorrower.IsActive, true)
}
