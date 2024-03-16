package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomUser(t *testing.T) *User {
	user, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		Username: "username",
		Password: "password",
		Role:     RoleEnumADMIN,
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.Username)
	require.NotZero(t, user.Password)
	require.NotZero(t, user.Role)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
	require.NotZero(t, user.IsActive, true)
	return user
}

func TestQueries_CountUser(t *testing.T) {
	removeAll()
	createRandomUser(t)
	count, err := testQueries.CountUser(context.Background())
	require.NoError(t, err)
	require.Equal(t, count, int64(1))
}

func TestQueries_CreateUser(t *testing.T) {
	removeAll()
	createRandomUser(t)
}

func TestQueries_DeleteOneUserById(t *testing.T) {
	removeAll()
	user := createRandomUser(t)

	deletedUser, err := testQueries.DeleteOneUserById(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedUser)
	require.NotZero(t, deletedUser.ID)
	require.NotZero(t, deletedUser.Username)
	require.NotZero(t, deletedUser.Password)
	require.NotZero(t, deletedUser.Role)
	require.NotZero(t, deletedUser.CreatedAt)
	require.NotZero(t, deletedUser.UpdatedAt)
	require.NotZero(t, deletedUser.IsActive, true)
}

func TestQueries_GetOneUserById(t *testing.T) {
	removeAll()
	user := createRandomUser(t)

	findOneUser, err := testQueries.GetOneUserById(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, findOneUser)
	require.NotZero(t, findOneUser.ID)
	require.NotZero(t, findOneUser.Username)
	require.NotZero(t, findOneUser.Password)
	require.NotZero(t, findOneUser.Role)
	require.NotZero(t, findOneUser.CreatedAt)
	require.NotZero(t, findOneUser.UpdatedAt)
	require.NotZero(t, findOneUser.IsActive, true)
}

func TestQueries_ListUser(t *testing.T) {
	removeAll()
	createRandomUser(t)
	createRandomUser(t)

	categories, err := testQueries.ListUser(context.Background(), ListUserParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	for _, user := range categories {
		require.NotEmpty(t, user)
		require.NotZero(t, user.ID)
		require.NotZero(t, user.Username)
		require.NotZero(t, user.Password)
		require.NotZero(t, user.Role)
		require.NotZero(t, user.CreatedAt)
		require.NotZero(t, user.UpdatedAt)
		require.NotZero(t, user.IsActive, true)
	}
}

func TestQueries_UpdateOneUserById(t *testing.T) {
	removeAll()
	user := createRandomUser(t)

	updatedUser, err := testQueries.UpdateOneUserById(context.Background(), UpdateOneUserByIdParams{
		ID:       user.ID,
		Username: "Updated",
		Role:     RoleEnumMANAGER,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.NotZero(t, updatedUser.ID)
	require.NotZero(t, updatedUser.Username)
	require.NotZero(t, updatedUser.Password)
	require.NotZero(t, updatedUser.Role)
	require.Equal(t, updatedUser.Username, "Updated")
	require.Equal(t, updatedUser.Role, RoleEnumMANAGER)
	require.NotZero(t, updatedUser.CreatedAt)
	require.NotZero(t, updatedUser.UpdatedAt)
	require.NotZero(t, updatedUser.IsActive, true)
}
