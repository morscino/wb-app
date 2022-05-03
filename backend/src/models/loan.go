package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID                uuid.UUID `gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	UserID            uuid.UUID
	RepaymentDuration int64
	OtherLoansAmount  float64
	LoanAmount        float64
	Repayment         float64     `gorm:"-"`
	RepaymentDates    []time.Time `gorm:"-"`
	AccountNumber     string
	AccountName       string
	Bank              string
	//ProcessingFee     float64
	AmountPaid       *float64
	RepaymentStatus  string
	Balance          *float64
	Status           string
	LoanApprovalDate sql.NullTime
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        sql.NullTime
}

const (
	LoanStatusApproved string = "approved"
	LoanStatusDeclined string = "declined"
	LoanStatusPending  string = "pending"

	RepaymentStatusPaid    string = "paid"
	RepaymentStatusNotPaid string = "not paid"
	RepaymentStatusOngoing string = "ongoing"

	RepaymentDurationTwoWeeks    int64 = 2
	RepaymentDurationOneMonth    int64 = 4
	RepaymentDurationTwoMonths   int64 = 8
	RepaymentDurationThreeMonths int64 = 12
	RepaymentDurationFourMonths  int64 = 16
	RepaymentDurationFiveMonths  int64 = 20
	RepaymentDurationSixMonths   int64 = 24
)

func (l *Loan) GetTotalLoan(interestRate, processingFee float64) float64 {
	// get loan total with interest and processing fee
	totalLoan := l.LoanAmount + (((interestRate / 100) * l.LoanAmount) * float64(l.RepaymentDuration)) + processingFee

	return totalLoan
}

func (l *Loan) GetWeeklyRepayment(interestRate, processingFee float64) float64 {
	repayment := l.GetTotalLoan(interestRate, processingFee) / float64(l.RepaymentDuration)
	return repayment
}

func (l *Loan) GetMonthlyRepayment(interestRate, processingFee float64) float64 {
	repayment := l.GetWeeklyRepayment(interestRate, processingFee) * 4
	return repayment
}

func (l *Loan) GetRepayment(interestRate, processingFee float64) float64 {
	var repayment float64
	if l.RepaymentDuration == RepaymentDurationTwoWeeks || l.RepaymentDuration == RepaymentDurationOneMonth {
		repayment = l.GetWeeklyRepayment(interestRate, processingFee) * float64(l.RepaymentDuration)
	} else if l.RepaymentDuration == RepaymentDurationTwoMonths {
		repayment = l.GetMonthlyRepayment(interestRate, processingFee) * 2
	} else if l.RepaymentDuration == RepaymentDurationThreeMonths {
		repayment = l.GetMonthlyRepayment(interestRate, processingFee) * 3
	} else if l.RepaymentDuration == RepaymentDurationFourMonths {
		repayment = l.GetMonthlyRepayment(interestRate, processingFee) * 4
	} else if l.RepaymentDuration == RepaymentDurationFiveMonths {
		repayment = l.GetMonthlyRepayment(interestRate, processingFee) * 5
	} else if l.RepaymentDuration == RepaymentDurationThreeMonths {
		repayment = l.GetMonthlyRepayment(interestRate, processingFee) * 6
	}
	return repayment
}

func (l *Loan) GetRepaymentDate() []time.Time {
	var t []time.Time

	if l.RepaymentDuration < RepaymentDurationTwoWeeks {
		return t
	}

	if l.RepaymentDuration == RepaymentDurationTwoWeeks || l.RepaymentDuration == RepaymentDurationOneMonth {
		t = append(t, l.LoanApprovalDate.Time.AddDate(0, 0, int(l.RepaymentDuration)*7))
	} else {
		m := l.RepaymentDuration / 4
		for i := 1; i <= int(m); i++ {
			t = append(t, l.LoanApprovalDate.Time.AddDate(0, i, 0))
		}
	}
	return t
}
