package postgres

import (
	"fmt"

	"github.com/MastoCred-Inc/web-app/database"
	"github.com/MastoCred-Inc/web-app/utility/environment"
	"github.com/MastoCred-Inc/web-app/utility/helper"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New Storage, however should panic if it can't be pinged. System should be able to connect to the database
func New(z zerolog.Logger, env *environment.Env) *database.Storage {
	l := z.With().Str(helper.LogStrKeyModule, database.PackageName).Logger()
	db, err := gorm.Open(
		postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Africa/Lagos",
			env.Get("PG_HOST"),
			env.Get("PG_PORT"),
			env.Get("PG_USER"),
			env.Get("PG_DATABASE"),
			env.Get("PG_PASSWORD"),
		)),
		&gorm.Config{},
	)
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
