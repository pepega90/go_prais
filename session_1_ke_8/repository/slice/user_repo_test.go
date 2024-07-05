package slice

import (
	"go_prais/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserRepo(t *testing.T) {
	repo := NewSliceRepository()

	t.Run("create user", func(t *testing.T) {
		createdUser := model.User{Name: "aji", Email: "aji@handsome.com", Password: "pass"}
		user := repo.CreateUser(createdUser)
		require.Equal(t, 1, user.Id)
		require.Equal(t, "aji", user.Name)
		require.Equal(t, "aji@handsome.com", user.Email)
		require.NotZero(t, user.CreatedAt)
		require.NotZero(t, user.UpdatedAt)
	})

	t.Run("get user", func(t *testing.T) {
		user, err := repo.GetUser(1)

		require.NoError(t, err)
		require.Equal(t, "aji", user.Name)
	})

	t.Run("update user", func(t *testing.T) {
		updated := model.User{Name: "pepeg"}

		u, err := repo.UpdateUser(1, updated)

		require.NoError(t, err)
		require.Equal(t, "pepeg", u.Name)
	})

	t.Run("delete user", func(t *testing.T) {
		deleted := repo.DeleteUser(1)

		require.NoError(t, deleted)
	})
}
