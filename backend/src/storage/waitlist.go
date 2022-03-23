package storage

import (
	"context"
	"strings"

	"github.com/MastoCred-Inc/web-app/database"
	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/rs/zerolog"
)

//go:generate mockgen -source waitlist.go -destination ./mock/mock_waitlist.go -package mock WaitlistStore
type WaitlistStore interface {
	CreateWaitList(ctx context.Context, waitlist models.Waitlist) (bool, error)
	GetWaitlistByEmail(ctx context.Context, email string) (models.Waitlist, error)
}

// Waitlist objectx
type Waitlist struct {
	logger  zerolog.Logger
	storage *database.Storage
}

// NewUser creates a new reference to the User storage entity
func NewWaitlist(s *database.Storage) *WaitlistStore {
	l := s.Logger.With().Str("app", "waitlist").Logger()
	waitlist := &Waitlist{
		logger:  l,
		storage: s,
	}
	waitlistDB := WaitlistStore(waitlist)
	return &waitlistDB
}

func (w *Waitlist) CreateWaitList(ctx context.Context, waitlist models.Waitlist) (bool, error) {
	db := w.storage.DB.WithContext(ctx).Create(&waitlist)
	if db.Error != nil {
		w.logger.Err(db.Error).Msgf("User::RegisterUser error: %v, (%v)", language.ErrText()[language.ErrRecordCreatingFailed], db.Error)
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return false, language.ErrText()[language.ErrDuplicateRecord]
		}
		return false, language.ErrText()[language.ErrRecordCreatingFailed]
	}
	return true, nil
}

func (w *Waitlist) GetWaitlistByEmail(ctx context.Context, email string) (models.Waitlist, error) {
	var waitlist models.Waitlist
	db := w.storage.DB.WithContext(ctx).Where("email = ?", email).Find(&waitlist)
	if db.Error != nil {
		w.logger.Err(db.Error).Msgf("User::GetUserByEmail error: %v, (%v)", language.ErrText()[language.ErrRecordNotFound], db.Error)
		return waitlist, language.ErrText()[language.ErrRecordNotFound]
	}
	return waitlist, nil
}
