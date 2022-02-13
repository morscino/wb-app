package graph

import "github.com/rs/zerolog"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	logger zerolog.Logger
}

// New created a new instance of Resolver
func New(l zerolog.Logger) *Resolver {
	return &Resolver{
		logger: l,
	}
}
