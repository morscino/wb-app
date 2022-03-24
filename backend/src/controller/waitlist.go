package controller

import (
	"context"

	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/models"
)

// CreateWaitlist creates a new waitlist
func (c *Controller) CreateWaitlist(ctx context.Context, waitlist *models.Waitlist) (bool, error) {
	// check if user already exists
	existingUser, err := c.waitlistStorage.GetWaitlistByEmail(ctx, waitlist.Email)
	if err != nil {
		c.logger.Err(err).Msgf("RegisterUser:GetUserByEmail [%v] : (%v)", existingUser.Email, err)
		return false, err
	}

	if (existingUser != models.Waitlist{}) {
		return false, language.ErrText()[language.ErrEmailAlreadyExist]
	}

	return c.waitlistStorage.CreateWaitList(ctx, *waitlist)
}

func (c *Controller) GetAllWaitlists(ctx context.Context, page models.Page, mode int) ([]*models.Waitlist, *models.PageInfo, error) {
	return c.waitlistStorage.GetAllWaitlists(ctx, page, mode)
}
