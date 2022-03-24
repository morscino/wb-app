package storage

import (
	"context"
	"fmt"
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
	GetAllWaitlists(ctx context.Context, page models.Page, mode int) ([]*models.Waitlist, *models.PageInfo, error)
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

func (w *Waitlist) GetAllWaitlists(ctx context.Context, page models.Page, mode int) ([]*models.Waitlist, *models.PageInfo, error) {
	var waitlists []*models.Waitlist
	offset := 0
	// load defaults
	if page.Number == nil {
		tmpPageNumber := models.PageDefaultNumber
		page.Number = &tmpPageNumber
	}
	if page.Size == nil {
		tmpPageSize := models.PageDefaultSize
		page.Size = &tmpPageSize
	}
	if page.SortBy == nil {
		tmpPageSortBy := models.PageDefaultSortBy
		page.SortBy = &tmpPageSortBy
	}
	if page.SortDirectionDesc == nil {
		tmpPageSortDirectionDesc := models.PageDefaultSortDirectionDesc
		page.SortDirectionDesc = &tmpPageSortDirectionDesc
	}

	if *page.Number > 1 {
		offset = *page.Size * (*page.Number - 1)
	}
	sortDirection := models.PageSortDirectionDescending
	if !*page.SortDirectionDesc {
		sortDirection = models.PageSortDirectionAscending
	}

	query := models.Waitlist{
		Mode: mode,
	}

	queryDraft := w.storage.DB.WithContext(ctx).Model(models.Waitlist{})
	dbCount := w.storage.DB.WithContext(ctx).Model(models.Waitlist{})

	if mode > 0 {
		queryDraft = queryDraft.Where(query)
		dbCount = dbCount.Where(query)
	}

	db := queryDraft.Offset(offset).Limit(*page.Size).
		Order(fmt.Sprintf("waitlists.%s %s", *page.SortBy, sortDirection)).
		Find(&waitlists)
	if db.Error != nil {
		w.logger.Err(db.Error).Msgf("Waitlist::GetAllWaitlists error: %v, (%v)", language.ErrText()[language.ErrRecordNotFound], db.Error)
		return nil, nil, language.ErrText()[language.ErrRecordNotFound]
	}

	// then do counting
	var count int64
	dbCount.Count(&count)

	return waitlists, &models.PageInfo{
		Page:            *page.Number,
		Size:            *page.Size,
		HasNextPage:     int64(offset+*page.Size) < count,
		HasPreviousPage: *page.Number > 1,
		TotalCount:      count,
	}, nil
}
