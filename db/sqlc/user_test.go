package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vin-oys/api-carpool/util"
)

func randomUserRole() UserRole {
	r := util.Random()
	userRole := []UserRole{UserRoleSuperAdministrator, UserRoleAdministrator, UserRolePassenger}
	k := len(userRole)

	return userRole[r.Intn(k)]
}

func CreateRandomUser(t *testing.T) UserCreateResponse {
	username := util.RandomUsername()
	arg := CreateUserParams{
		Username:      username,
		Password:      util.RandomNumberInString(8),
		ContactNumber: util.GetContactNumberFromUsername(username),
		RoleID:        randomUserRole(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.RoleID, user.RoleID)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.RoleID, user2.RoleID)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	arg := UpdateUserParams{
		Username:      user1.Username,
		ContactNumber: "",
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.RoleID, user2.RoleID)

	require.WithinDuration(t, user1.CreatedAt, HandleNullTime(user2.UpdatedAt), time.Second)

}

func TestUpdateUserPassword(t *testing.T) {
	user1 := CreateRandomUser(t)
	arg := UpdateUserPasswordParams{
		Username: user1.Username,
		Password: util.RandomNumberInString(8),
	}
	user2, err := testQueries.UpdateUserPassword(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.RoleID, user2.RoleID)

	require.WithinDuration(t, user1.CreatedAt, HandleNullTime(user2.UpdatedAt), time.Second)

}

func TestUpdateUserRole(t *testing.T) {
	user1 := CreateRandomUser(t)
	arg := UpdateUserRoleParams{
		Username: user1.Username,
		RoleID:   UserRoleDriver,
	}
	user2, err := testQueries.UpdateUserRole(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, arg.RoleID, user2.RoleID)

	require.WithinDuration(t, user1.CreatedAt, HandleNullTime(user2.UpdatedAt), time.Second)

}

func TestDeleteUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.Username)

	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
		RoleID: "driver",
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

}

func HandleNullTime(nt sql.NullTime) time.Time {
	if !nt.Valid {
		return time.Time{}
	}
	return nt.Time
}
