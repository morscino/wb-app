package graph

import (
	"context"

	"github.com/MastoCred-Inc/web-app/h/graph/generated"
	"github.com/MastoCred-Inc/web-app/h/graph/model"
	"github.com/MastoCred-Inc/web-app/h/graph/translator"
	"github.com/MastoCred-Inc/web-app/models"
)

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUser) (*models.User, error) {

	// validate user input
	userModel, err := translator.ConvertUserInputToUserModel(input)
	if err != nil {
		return nil, err
	}

	// send user to controller
	user, err := r.controller.RegisterUser(ctx, *userModel)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
