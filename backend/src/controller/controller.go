package controller

import (
	"context"

	"github.com/MastoCred-Inc/web-app/database"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/MastoCred-Inc/web-app/storage"
	"github.com/rs/zerolog"
)

type Operations interface {
	RegisterUser(ctx context.Context, user models.User) (*models.User, error)
}

type Controller struct {
	logger      zerolog.Logger
	userStorage storage.UserStore
}

func New(l zerolog.Logger, s *database.Storage) *Operations {
	user := storage.NewUser(s)

	// build controller struct
	c := &Controller{
		logger:      l,
		userStorage: *user,
	}
	op := Operations(c)
	return &op
}
