package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/MastoCred-Inc/web-app/h/graph/generated"
	"github.com/MastoCred-Inc/web-app/h/graph/model"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/MastoCred-Inc/web-app/utility/helper"
)

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AuthenticateUser(ctx context.Context, email string, password string) (*model.UserAuthenticated, error) {
	ginC, err := helper.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.controller.Middleware().AuthenticateUser(ginC, email, password)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
