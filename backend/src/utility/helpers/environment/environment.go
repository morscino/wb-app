// Package environment defines helpers accessing environment values
package environment

import (
	"github.com/joho/godotenv"
)

// Env represents environmental variable instance
type Env struct{}

// New creates a new instance of Env and returns an error if any occurs
func New() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Env{}, nil
}
