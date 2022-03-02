package storage

import (
	"context"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MastoCred-Inc/web-app/utility/test_util"
	"github.com/stretchr/testify/require"
)

var userTableColumns = []string{"last_name", "first_name", "email", "password", "salt", "username", "created_at", "updated_at", "id"}

func (s *Suite) TestRegisterUser() {
	testUser := test_util.NewTestUser()

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("last_name","first_name","email","password","salt","username","created_at","updated_at","id")
	 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`)).
		WithArgs(testUser.LastName, testUser.FirstName, testUser.Email, testUser.Password, testUser.Salt, testUser.Username, testUser.CreatedAt, testUser.UpdatedAt, testUser.ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(testUser.ID),
		)
	s.mock.ExpectCommit()
	retUser, err := s.userDatabase.RegisterUser(context.Background(), testUser)

	require.NoError(s.T(), err)
	require.Equal(s.T(), retUser, testUser)
}

func (s *Suite) TestGetUserByEmail() {
	testUser := test_util.NewTestUser()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email = $1`)).
		WithArgs(testUser.Email).
		WillReturnRows(sqlmock.NewRows(userTableColumns).
			AddRow(testUser.LastName, testUser.FirstName, testUser.Email, testUser.Password, testUser.Salt, testUser.Username, testUser.CreatedAt, testUser.UpdatedAt, testUser.ID))

	retUser, err := s.userDatabase.GetUserByEmail(context.Background(), testUser.Email)

	require.NoError(s.T(), err)
	require.Equal(s.T(), retUser, testUser)

}
