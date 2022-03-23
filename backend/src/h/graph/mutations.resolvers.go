package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MastoCred-Inc/web-app/h/graph/generated"
	"github.com/MastoCred-Inc/web-app/h/graph/model"
	"github.com/MastoCred-Inc/web-app/h/graph/translator"
	"github.com/MastoCred-Inc/web-app/models"
)

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

func (r *mutationResolver) CreateWaitList(ctx context.Context, input model.RegisterWaitlist) (bool, error) {
	var created bool
	// validate user input
	waitlist, err := translator.ConvertWaitlistInputToWaitlistModel(input)
	if err != nil {
		return false, err
	}

	// send waitlist to controller
	created, err = r.controller.CreateWaitlist(ctx, waitlist)
	if err != nil {
		return false, err
	}

	return created, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
