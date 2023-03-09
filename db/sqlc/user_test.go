package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:      "(+65)98765432",
		Password:      "password",
		ContactNumber: "+6598765432",
		RoleID:        "super_administrator",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.ContactNumber, user.ContactNumber)
	require.Equal(t, arg.RoleID, user.RoleID)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}
