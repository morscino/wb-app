package controller

import (
	"context"

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
	u.Username = salt // temporary value, to be handled later

	// send data to the storage
	user, err := c.userStorage.RegisterUser(ctx, u)
	if err != nil {
		c.logger.Err(err).Msgf("RegisterUser:RegisterUser [%v] : (%v)", u.Email, err)
		return nil, err
	}

	return &user, nil
}
