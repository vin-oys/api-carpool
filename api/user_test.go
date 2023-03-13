package api

import (
	"database/sql"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	db "github.com/vin-oys/api-carpool/db/sqlc"
	"github.com/vin-oys/api-carpool/util"
	"reflect"
	"testing"
)

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.Password)
	if err != nil {
		return false
	}

	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomNumberInString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username: util.RandomUsername(),
		Password: hashedPassword,
		Firstname: sql.NullString{
			String: util.RandomNumberInString(6),
		},
		Lastname: sql.NullString{
			String: util.RandomNumberInString(6),
		},
		ContactNumber: util.RandomNumberInString(8),
		CreatedAt:     util.RandomTime(),
		UpdatedAt: sql.NullTime{
			Time: util.RandomTime(),
		},
		RoleID: db.UserRoleAdministrator,
	}

	return
}
