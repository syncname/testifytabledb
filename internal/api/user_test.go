package api

import (
	"database/sql"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/syncname/testifyexample/internal/models"
	"github.com/syncname/testifyexample/internal/util"
	"testing"
	"time"
)

const (
	UniqueViolationErr = pq.ErrorCode("23505")
)

func IsErrorCode(err error, errcode pq.ErrorCode) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == errcode
	}
	return false
}

func newRandomUser() *models.User {
	return &models.User{
		Email: util.RandomMail(),
		Name:  util.RandomName(),
	}
}

func CreateRandomUser(t *testing.T) *models.User {
	newUser := newRandomUser()
	u, err := testServer.CreateUser(newUser)
	assert.NoError(t, err)
	assert.Equal(t, newUser.Name, u.Name)
	assert.Equal(t, newUser.Email, u.Email)
	assert.NotZero(t, u.CreatedAt)
	return u
}

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"create random user", func() {
			CreateRandomUser(t)
		}},
		{"unique name error", func() {

			u := CreateRandomUser(t)

			u.Email = util.RandomMail()
			violationUser, err := testServer.CreateUser(u)
			assert.Equal(t, true, IsErrorCode(err, UniqueViolationErr))
			//не нужно делать так, т.к. код проверки на nil
			///несколько сложнее
			//https://stackoverflow.com/questions/55900390/how-to-detect-whether-a-struct-pointer-is-nil-in-golang
			//assert.Equal(t, violationUser, nil)
			assert.Nil(t, violationUser, nil)

		}},
		{"unique email error", func() {

			u := CreateRandomUser(t)

			u.Name = util.RandomName()
			violationUser, err := testServer.CreateUser(u)
			assert.Equal(t, true, IsErrorCode(err, UniqueViolationErr))
			assert.Nil(t, violationUser)

		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}

	//u, err := testServer.CreateUser(models.User{
	//	Email: "new@mail",
	//	Name:  "user",
	//})

	//https://www.postgresql.org/docs/current/errcodes-appendix.html
	//Больше примеров
	//https://github.com/go-gorm/gorm/issues/4135#issuecomment-1336553267

	//assert.Equal(t, true, IsErrorCode(err, UniqueViolationErr))
	//
	//fmt.Println(u, err)

}

func TestUpdateUser(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"update user email", func() {

			user := CreateRandomUser(t)
			user.Email = util.RandomMail()
			updUser, err := testServer.UpdateUserEmail(user)

			assert.NoError(t, err)
			assert.Equal(t, user.Name, updUser.Name)
			assert.Equal(t, user.Email, updUser.Email)
			assert.WithinDuration(t, user.CreatedAt, updUser.CreatedAt, time.Second)

		}},
		{"update unknown user", func() {
			user := CreateRandomUser(t)
			user.Email = util.RandomMail()
			user.Name = util.RandomName()
			updUser, err := testServer.UpdateUserEmail(user)
			assert.ErrorIs(t, err, sql.ErrNoRows)
			assert.Nil(t, updUser)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestGetUser(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"get user", func() {
			user := CreateRandomUser(t)
			newUser, err := testServer.GetUser(user.Name)
			assert.NoError(t, err)
			assert.Equal(t, user.Name, newUser.Name)
			assert.Equal(t, user.Email, newUser.Email)
			assert.WithinDuration(t, user.CreatedAt, newUser.CreatedAt, time.Second)

		}},
		{"get unknown user", func() {
			newUser, err := testServer.GetUser(util.RandomName())
			assert.ErrorIs(t, err, sql.ErrNoRows)
			assert.Nil(t, newUser)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestDeleteUser(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"delete user", func() {
			user := CreateRandomUser(t)
			newUser, err := testServer.DeleteUser(user.Name)
			assert.NoError(t, err)
			assert.Equal(t, user.Name, newUser.Name)
			assert.Equal(t, user.Email, newUser.Email)
			assert.WithinDuration(t, user.CreatedAt, newUser.CreatedAt, time.Second)
		}},
		{"delete unknown user", func() {
			newUser, err := testServer.DeleteUser(util.RandomName())
			assert.ErrorIs(t, err, sql.ErrNoRows)
			assert.Nil(t, newUser)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}
