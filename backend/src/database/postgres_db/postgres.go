package postgres_db

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"gitlab.com/mastocred/web-app/database"
	"gitlab.com/mastocred/web-app/utility/environment"
	"gitlab.com/mastocred/web-app/utility/gorm_sqlmock"
	"gitlab.com/mastocred/web-app/utility/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New Storage, however should panic if it can't be pinged. System should be able to connect to the database
func New(z zerolog.Logger, env *environment.Env) *database.Storage {
	l := z.With().Str(helper.LogStrKeyModule, database.PackageName).Logger()
	dialect := postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Africa/Lagos",
		env.Get("PG_HOST"),
		env.Get("PG_PORT"),
		env.Get("PG_USER"),
		env.Get("PG_DATABASE"),
		env.Get("PG_PASSWORD"),
	))
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		l.Fatal().Err(err)
		panic(err)
	}

	return &database.Storage{
		Logger: l,
		Env:    env,
		DB:     db,
	}
}

// GetStorage helper for tests/mock
// I expect our storage tests to use this helper going forward.
func GetStorage(t *testing.T) (sqlmock.Sqlmock, *database.Storage) {
	var (
		mock sqlmock.Sqlmock
		db   *gorm.DB
		err  error
	)

	db, mock, err = gorm_sqlmock.New(gorm_sqlmock.Config{
		Config:     &gorm.Config{},
		DriverName: "postgres",
		DSN:        "mock",
	})

	require.NoError(t, err)

	return mock, NewFromDB(db)
}

// NewFromDB created a new storage with just the database reference passed in
func NewFromDB(db *gorm.DB) *database.Storage {
	return &database.Storage{
		DB: db,
	}
}
