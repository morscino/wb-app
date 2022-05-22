package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/MastoCred-Inc/web-app/h/graph/generated"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/MastoCred-Inc/web-app/utility/helper"
)

func (r *loanResolver) ID(ctx context.Context, obj *models.Loan) (string, error) {
	return obj.ID.String(), nil
}

func (r *loanResolver) RepaymentDuration(ctx context.Context, obj *models.Loan) (*string, error) {
	var d string
	switch obj.RepaymentDuration {
	case models.RepaymentDurationTwoWeeks:
		d = "2 weeks"
	default:
		var month string
		m := obj.RepaymentDuration / 4
		if m <= 1 {
			month = "month"
		} else {
			month = "months"
		}
		t := helper.ConverInt64ToString(m)
		d = fmt.Sprintf("%v %v", t, month)

	}
	return &d, nil
}

func (r *loanResolver) RepaymentAmount(ctx context.Context, obj *models.Loan) (*float64, error) {
	interestRate, _ := helper.StringToFloat64(r.env.Get("INTEREST_RATE"))
	processingFee, _ := helper.StringToFloat64(r.env.Get("PROCESSING_FEE"))
	amount := obj.GetRepayment(interestRate, processingFee)

	return &amount, nil
}

func (r *loanResolver) LoanRepaymentDates(ctx context.Context, obj *models.Loan) ([]*string, error) {
	var dates []*string
	repayments := obj.GetRepaymentDates()
	for i := 0; i < len(repayments); i++ {
		s := repayments[i].String()
		dates = append(dates, &s)
	}
	return dates, nil
}

func (r *loanResolver) LoanApprovalDate(ctx context.Context, obj *models.Loan) (*string, error) {
	l := obj.LoanApprovalDate.Time.String()
	return &l, nil
}

func (r *loanInstalmentResolver) ID(ctx context.Context, obj *models.LoanInstalment) (string, error) {
	return obj.ID.String(), nil
}

func (r *loanInstalmentResolver) UserID(ctx context.Context, obj *models.LoanInstalment) (*string, error) {
	u := obj.UserID.String()
	return &u, nil
}

func (r *loanInstalmentResolver) LoanID(ctx context.Context, obj *models.LoanInstalment) (*string, error) {
	l := obj.LoanID.String()
	return &l, nil
}

func (r *loanInstalmentResolver) LoanRepaymentAmount(ctx context.Context, obj *models.LoanInstalment) (*float64, error) {
	l := obj.RepaymentAmount
	return &l, nil
}

func (r *loanInstalmentResolver) LoanRepaymentDate(ctx context.Context, obj *models.LoanInstalment) (*string, error) {
	u := obj.RepaymentDate.String()
	return &u, nil
}

func (r *loanInstalmentResolver) RepaymentDuration(ctx context.Context, obj *models.LoanInstalment) (*string, error) {
	var d string
	switch obj.RepaymentDuration {
	case models.RepaymentDurationTwoWeeks:
		d = "2 weeks"
	default:
		var month string
		m := obj.RepaymentDuration / 4
		if m <= 1 {
			month = "month"
		} else {
			month = "months"
		}
		t := helper.ConverInt64ToString(m)
		d = fmt.Sprintf("%v %v", t, month)

	}

	return &d, nil
}

// Loan returns generated.LoanResolver implementation.
func (r *Resolver) Loan() generated.LoanResolver { return &loanResolver{r} }

// LoanInstalment returns generated.LoanInstalmentResolver implementation.
func (r *Resolver) LoanInstalment() generated.LoanInstalmentResolver {
	return &loanInstalmentResolver{r}
}

type loanResolver struct{ *Resolver }
type loanInstalmentResolver struct{ *Resolver }
