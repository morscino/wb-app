package graph

import (
	"github.com/MastoCred-Inc/web-app/controller"
	"github.com/rs/zerolog"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	logger     zerolog.Logger
	controller controller.Operations
}

// New created a new instance of Resolver
func New(l zerolog.Logger, c controller.Operations) *Resolver {
	return &Resolver{
		logger:     l,
		controller: c,
	}
}
