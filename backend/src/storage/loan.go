package storage

import (
	"github.com/MastoCred-Inc/web-app/database"
	"github.com/rs/zerolog"
)

//go:generate mockgen -source loan.go -destination ./mock/loan_user.go -package mock LoanStore
type LoanStore interface {
}

// Loan objectx
type Loan struct {
	logger  zerolog.Logger
	storage *database.Storage
}

// NewLoan creates a new reference to the Loan storage entity
func NewLoan(s *database.Storage) *LoanStore {
	l := s.Logger.With().Str("app", "user").Logger()
	loan := &Loan{
		logger:  l,
		storage: s,
	}
	loanDB := LoanStore(loan)
	return &loanDB
}
