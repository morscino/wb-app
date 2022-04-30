package graph

import (
	"github.com/MastoCred-Inc/web-app/controller"
	"github.com/MastoCred-Inc/web-app/utility/environment"
	"github.com/rs/zerolog"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	logger     zerolog.Logger
	env        *environment.Env
	controller controller.Operations
}

// New created a new instance of Resolver
func New(l zerolog.Logger, c controller.Operations, env *environment.Env) *Resolver {
	return &Resolver{
		logger:     l,
		controller: c,
		env:        env,
	}
}
