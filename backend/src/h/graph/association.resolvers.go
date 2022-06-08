package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/mastocred/web-app/h/graph/generated"
	"gitlab.com/mastocred/web-app/models"
)

func (r *associationResolver) ID(ctx context.Context, obj *models.Association) (string, error) {
	return obj.ID.String(), nil
}

// Association returns generated.AssociationResolver implementation.
func (r *Resolver) Association() generated.AssociationResolver { return &associationResolver{r} }

type associationResolver struct{ *Resolver }
