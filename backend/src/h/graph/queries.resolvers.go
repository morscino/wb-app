package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	"github.com/MastoCred-Inc/web-app/h/graph/generated"
	"github.com/MastoCred-Inc/web-app/h/graph/model"
	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/MastoCred-Inc/web-app/utility/helper"
)

func (r *queryResolver) GetAllUsers(ctx context.Context, page models.Page) (*model.GetUsersResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AuthenticateUser(ctx context.Context, email string, password string) (*model.UserAuthenticated, error) {
	ginC, err := helper.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.controller.Middleware().AuthenticateUser(ginC, email, password)
}

func (r *queryResolver) GeAllWaitlists(ctx context.Context, input model.GetWaitlistsRequest) (*model.GetWaitlistsResult, error) {
	ginC, err := helper.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if input.Mode == nil {
		m := models.WaitListModeAll
		input.Mode = &m
	}

	waitlists, page, err := r.controller.GetAllWaitlists(ginC, *input.Page, helper.ConvertModeToIntPointer(string(*input.Mode)))
	if err != nil {
		return nil, err
	}

	return &model.GetWaitlistsResult{
		Items: waitlists,
		Page:  page,
	}, nil
}

func (r *queryResolver) GetAllAssociations(ctx context.Context, input model.GetAssociationsRequest) (*model.GetAssociationsResult, error) {
	ginC, err := helper.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	associations, page, err := r.controller.GetAllAssociations(ginC, *input.Page)
	if err != nil {
		return nil, err
	}

	return &model.GetAssociationsResult{
		Items: associations,
		Page:  page,
	}, nil
}

func (r *queryResolver) GetAllLoans(ctx context.Context, page models.Page) (*model.GetLoansResult, error) {
	ginC, err := helper.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	actorUser, err := r.controller.Middleware().PasetoUserAuth(ginC)
	if err != nil {
		return nil, err
	}

	if actorUser.UserType != int64(models.UserTypeAdmin) {
		return nil, language.ErrText()[language.ErrAccessDenied]
	}

	loans, loanPage, err := r.controller.GetAllLoans(ginC, page)
	if err != nil {
		return nil, err
	}

	return &model.GetLoansResult{
		Items: loans,
		Page:  loanPage,
	}, nil
}

func (r *queryResolver) GetLoansByUserID(ctx context.Context, userID string) ([]*models.Loan, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetLoanInstalmentsByUserID(ctx context.Context, userID string, page models.Page) (*model.GetLoanInstalmentsResult, error) {
	ginC, err := helper.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	actorUser, err := r.controller.Middleware().PasetoUserAuth(ginC)
	if err != nil {
		return nil, err
	}

	if actorUser.UserType != int64(models.UserTypeAdmin) && !strings.EqualFold(userID, actorUser.ID.String()) {
		return nil, language.ErrText()[language.ErrAccessDenied]
	}

	u, err := helper.StringToUuid(userID)
	if err != nil {
		return nil, language.ErrText()[language.ErrParseError]
	}

	loanInstalments, loanPage, err := r.controller.GetLoanInstalmentsByUserID(ginC, u, page)
	if err != nil {
		return nil, err
	}

	return &model.GetLoanInstalmentsResult{
		Items: loanInstalments,
		Page:  loanPage,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
