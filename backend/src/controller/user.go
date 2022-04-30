package controller

import (
	"context"
	"time"

	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/models"
)

func (c *Controller) RegisterUser(ctx context.Context, u models.User) (*models.User, error) {
	// check if user already exists
	existingUser, err := c.userStorage.GetUserByEmail(ctx, u.Email)
	if err != nil {
		c.logger.Err(err).Msgf("RegisterUser:GetUserByEmail [%v] : (%v)", u.Email, err)
		return nil, err
	}

	if (existingUser != models.User{}) {
		return nil, language.ErrText()[language.ErrEmailAlreadyExist]
	}

	salt := u.GenerateSalt()
	encryptedPassword, err := u.EncyptPassword(u.Password, salt)

	if err != nil {
		c.logger.Err(err).Msgf("RegisterUser:EncyptPassword [%v] : (%v)", u.Email, err)
		return nil, err
	}

	u.Salt = salt
	u.Password = encryptedPassword

	// send data to the storage
	user, err := c.userStorage.RegisterUser(ctx, u)
	if err != nil {
		c.logger.Err(err).Msgf("RegisterUser:RegisterUser [%v] : (%v)", u.Email, err)
		return nil, err
	}

	return &user, nil
}

func (c *Controller) UpdateUserByID(ctx context.Context, u models.User) (*models.User, error) {

	u.UpdatedAt = time.Now()
	user, err := c.userStorage.UpdateUserByID(ctx, u.ID, u)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
