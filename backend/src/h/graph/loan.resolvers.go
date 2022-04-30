package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/MastoCred-Inc/web-app/h/graph/generated"
	"github.com/MastoCred-Inc/web-app/models"
)

func (r *loanResolver) ID(ctx context.Context, obj *models.Loan) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *loanResolver) Payback(ctx context.Context, obj *models.Loan) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *loanResolver) OtherLoansAmount(ctx context.Context, obj *models.Loan) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *loanResolver) LoanAmount(ctx context.Context, obj *models.Loan) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *loanResolver) MonthlyRepayment(ctx context.Context, obj *models.Loan) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *loanResolver) MonthlyInterestRate(ctx context.Context, obj *models.Loan) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *loanResolver) ProcessingFee(ctx context.Context, obj *models.Loan) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Loan returns generated.LoanResolver implementation.
func (r *Resolver) Loan() generated.LoanResolver { return &loanResolver{r} }

type loanResolver struct{ *Resolver }
