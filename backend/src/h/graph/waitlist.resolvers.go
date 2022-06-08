package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/mastocred/web-app/h/graph/generated"
	"gitlab.com/mastocred/web-app/models"
)

func (r *waitlistResolver) ID(ctx context.Context, obj *models.Waitlist) (string, error) {
	return obj.ID.String(), nil
}

func (r *waitlistResolver) Mode(ctx context.Context, obj *models.Waitlist) (models.WaitlistMode, error) {
	var m models.WaitlistMode
	for k, v := range models.WaitListModeMap {
		if v == obj.Mode {
			m = models.WaitlistMode(k)
		}
	}
	return m, nil
}

func (r *waitlistResolver) RegisteredAt(ctx context.Context, obj *models.Waitlist) (*string, error) {
	t := obj.CreatedAt.String()
	return &t, nil
}

// Waitlist returns generated.WaitlistResolver implementation.
func (r *Resolver) Waitlist() generated.WaitlistResolver { return &waitlistResolver{r} }

type waitlistResolver struct{ *Resolver }
