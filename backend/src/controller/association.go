package controller

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/mastocred/web-app/models"
)

func (c *Controller) RegisterAssociation(ctx context.Context, a models.Association) (*models.Association, error) {
	assoc, err := c.associationstorage.RegisterAssociation(ctx, a)
	if err != nil {
		return nil, err
	}

	return &assoc, nil
}

func (c *Controller) GetAllAssociations(ctx context.Context, page models.Page) ([]*models.Association, *models.PageInfo, error) {
	return c.associationstorage.GetAllAssociations(ctx, page)
}

func (c *Controller) GetAssociationById(ctx context.Context, id uuid.UUID) (models.Association, error) {
	return c.associationstorage.GetAssociationById(ctx, id)
}
