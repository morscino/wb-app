package storage

import (
	"context"
	"strings"

	"github.com/MastoCred-Inc/web-app/database"
	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/rs/zerolog"
)

type UserStore interface {
	RegisterUser(ctx context.Context, user models.User) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

// User objectx
type User struct {
	logger  zerolog.Logger
	storage *database.Storage
}

// NewUser creates a new reference to the User storage entity
func NewUser(s *database.Storage) *UserStore {
	l := s.Logger.With().Str("app", "user").Logger()
	user := &User{
		logger:  l,
		storage: s,
	}
	userDB := UserStore(user)
	return &userDB
}

// RegisterUser adds user record to the database
func (u *User) RegisterUser(ctx context.Context, user models.User) (models.User, error) {

	db := u.storage.DB.WithContext(ctx).Create(&user)
	if db.Error != nil {
		u.logger.Err(db.Error).Msgf("User::RegisterUser error: %v, (%v)", language.ErrText()[language.ErrRecordCreatingFailed], db.Error)
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return models.User{}, language.ErrText()[language.ErrDuplicateRecord]
		}
		return models.User{}, language.ErrText()[language.ErrRecordCreatingFailed]
	}
	return user, nil
}

func (u *User) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	db := u.storage.DB.WithContext(ctx).Where("email = ?", email).Find(&user)
	if db.Error != nil {
		u.logger.Err(db.Error).Msgf("User::GetUserByEmail error: %v, (%v)", language.ErrText()[language.ErrRecordNotFound], db.Error)
		return user, language.ErrText()[language.ErrRecordNotFound]
	}
	return user, nil
}
