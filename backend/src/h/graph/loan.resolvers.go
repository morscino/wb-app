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
		t := helper.ConverInt64ToString(obj.RepaymentDuration % 4)
		d = fmt.Sprintf("%v months", t)
	}

	return &d, nil
}

func (r *loanResolver) RepaymentDates(ctx context.Context, obj *models.Loan) ([]*string, error) {
	var dates []*string
	for i := 0; i < len(obj.RepaymentDates); i++ {
		s := obj.RepaymentDates[i].String()
		dates = append(dates, &s)
	}
	return dates, nil
}

func (r *loanResolver) LoanApprovalDate(ctx context.Context, obj *models.Loan) (*string, error) {
	l := obj.LoanApprovalDate.Time.String()
	return &l, nil
}

// Loan returns generated.LoanResolver implementation.
func (r *Resolver) Loan() generated.LoanResolver { return &loanResolver{r} }

type loanResolver struct{ *Resolver }
