package storage

import (
	"context"
	"strings"

	"github.com/MastoCred-Inc/web-app/database"
	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

//go:generate mockgen -source user.go -destination ./mock/mock_user.go -package mock UserStore
type UserStore interface {
	RegisterUser(ctx context.Context, user models.User) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	UpdateUserByID(ctx context.Context, id uuid.UUID, user models.User) (models.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error)
}

// User object
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

func (u *User) GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	var user models.User
	db := u.storage.DB.WithContext(ctx).Where("id = ?", userID).Find(&user)
	if db.Error != nil {
		u.logger.Err(db.Error).Msgf("User::GetUserByID error: %v, (%v)", language.ErrText()[language.ErrRecordNotFound], db.Error)
		return user, language.ErrText()[language.ErrRecordNotFound]
	}
	return user, nil
}

// UpdateByID should update the user record in the storage
func (u *User) UpdateUserByID(ctx context.Context, id uuid.UUID, user models.User) (models.User, error) {
	db := u.storage.DB.WithContext(ctx).Model(models.User{
		ID: id,
	}).UpdateColumns(models.User{
		FirstName:                user.FirstName,
		Email:                    user.Email,
		Telephone:                user.Telephone,
		Password:                 user.Password,
		Salt:                     user.Salt,
		UserType:                 user.UserType,
		AssociationID:            user.AssociationID,
		AssociationBranch:        user.AssociationBranch,
		BusinessName:             user.BusinessName,
		BusinessRegistrationDate: user.BusinessRegistrationDate,
		BusinessRCNumber:         user.BusinessRCNumber,
		Occupation:               user.Occupation,
		Salary:                   user.Salary,
		DateOfBirth:              user.DateOfBirth,
		MaritalStatus:            user.MaritalStatus,
		MeansOfIdentification:    user.MeansOfIdentification,
		ProfilePictureURL:        user.ProfilePictureURL,
		DocumentURL:              user.DocumentURL,
		State:                    user.State,
		LocalGovernment:          user.LocalGovernment,
		BVN:                      user.BVN,

		UpdatedAt: user.UpdatedAt, //disabled hooks and manually adding updatedAt here by self

	})
	if db.Error != nil {
		u.logger.Err(db.Error).Msgf("User::UpdateByID error: %v, (%v)", language.ErrText()[language.ErrRecordUpdateFailed], db.Error)
		return user, language.ErrText()[language.ErrRecordUpdateFailed]
	}
	user.ID = id

	return user, nil
}
