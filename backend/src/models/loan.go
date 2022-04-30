package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID                  uuid.UUID
	UserID              uuid.UUID
	PayBack             int64
	OtherLoansAmount    float64
	LoanAmount          float64
	MonthlyRepayment    float64
	MonthlyInterestRate float64
	AccountNumber       string
	AccountName         string
	Bank                string
	ProcessingFee       float64
	Status              string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           sql.NullTime
}
