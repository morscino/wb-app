package controller

import (
	"context"

	"github.com/MastoCred-Inc/web-app/database"
	"github.com/MastoCred-Inc/web-app/middleware"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/MastoCred-Inc/web-app/storage"
	"github.com/rs/zerolog"
)

//go:generate mockgen -source controller.go -destination ./mock/mock_controller.go -package mock Operations
type Operations interface {
	Middleware() *middleware.Middleware
	RegisterUser(ctx context.Context, user models.User) (*models.User, error)
}

type Controller struct {
	logger      zerolog.Logger
	userStorage storage.UserStore
	middleware  *middleware.Middleware
}

func New(l zerolog.Logger, s *database.Storage, middleware *middleware.Middleware) *Operations {
	user := storage.NewUser(s)

	// build controller struct
	c := &Controller{
		logger:      l,
		userStorage: *user,
		middleware:  middleware,
	}
	op := Operations(c)
	return &op
}

func (c *Controller) Middleware() *middleware.Middleware {
	return c.middleware
}
