package storage

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MastoCred-Inc/web-app/database"
	"github.com/MastoCred-Inc/web-app/database/postgres_db"
	"github.com/stretchr/testify/suite"
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
