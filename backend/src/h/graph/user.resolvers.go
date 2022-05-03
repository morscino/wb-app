package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MastoCred-Inc/web-app/h/graph/generated"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/google/uuid"
)

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return obj.ID.String(), nil
}

func (r *userResolver) UserType(ctx context.Context, obj *models.User) (*string, error) {
	u := models.UserTypeMap[models.UserType(obj.UserType)]

	return &u, nil
}

func (r *userResolver) Association(ctx context.Context, obj *models.User) (*models.Association, error) {
	var assocID uuid.UUID

	if obj.AssociationID != nil {
		assocID = *obj.AssociationID
	}
	assoc, err := r.controller.GetAssociationById(ctx, assocID)

	if err != nil {
		return nil, err
	}

	return &assoc, nil
}

func (r *userResolver) BusinessRegistrationDate(ctx context.Context, obj *models.User) (*string, error) {
	var b string
	if obj.BusinessRegistrationDate != nil {
		b = obj.BusinessRegistrationDate.Time.String()
	}

	return &b, nil
}

func (r *userResolver) DateOfBirth(ctx context.Context, obj *models.User) (*string, error) {
	var d string
	if obj.DateOfBirth != nil {
		d = obj.DateOfBirth.Time.String()
	}

	return &d, nil
}

func (r *userResolver) MeansOfIdentification(ctx context.Context, obj *models.User) (*string, error) {
	var m string
	if obj.MeansOfIdentification != nil {
		m = models.MeansOfIdentificationMap[models.MeansOfIdentification(*obj.MeansOfIdentification)]
	}

	return &m, nil
}

func (r *userResolver) ProfilePicture(ctx context.Context, obj *models.User) (*string, error) {
	return obj.ProfilePictureURL, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *userResolver) LocalGovernment(ctx context.Context, obj *models.User) (*string, error) {
	return obj.LocalGovernment, nil
}
