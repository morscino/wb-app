package storage

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gitlab.com/mastocred/web-app/database"
	"gitlab.com/mastocred/web-app/database/postgres_db"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB           *gorm.DB
	userDatabase UserStore
	mock         sqlmock.Sqlmock
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))

}

func (s *Suite) SetupSuite() {
	var store *database.Storage
	s.mock, store = postgres_db.GetStorage(s.T())

	s.userDatabase = *NewUser(store)
}
