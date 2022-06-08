package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gitlab.com/mastocred/web-app/database"
	"gitlab.com/mastocred/web-app/language"
	"gitlab.com/mastocred/web-app/models"
)

//go:generate mockgen -source association.go -destination ./mock/mock_association.go -package mock AssociationStore
type AssociationStore interface {
	RegisterAssociation(ctx context.Context, association models.Association) (models.Association, error)
	GetAllAssociations(ctx context.Context, page models.Page) ([]*models.Association, *models.PageInfo, error)
	GetAssociationById(ctx context.Context, id uuid.UUID) (models.Association, error)
}

// Association objectx
type Association struct {
	logger  zerolog.Logger
	storage *database.Storage
}

// NewUser creates a new reference to the User storage entity
func NewAssociation(s *database.Storage) *AssociationStore {
	l := s.Logger.With().Str("app", "user").Logger()
	assoc := &Association{
		logger:  l,
		storage: s,
	}
	assocDB := AssociationStore(assoc)
	return &assocDB
}

// RegisterAssociation adds a new association record to the database
func (a *Association) RegisterAssociation(ctx context.Context, association models.Association) (models.Association, error) {

	db := a.storage.DB.WithContext(ctx).Create(&association)
	if db.Error != nil {
		a.logger.Err(db.Error).Msgf("Association::RegisterAssociation error: %v, (%v)", language.ErrText()[language.ErrRecordCreatingFailed], db.Error)
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return models.Association{}, language.ErrText()[language.ErrDuplicateRecord]
		}
		return models.Association{}, language.ErrText()[language.ErrRecordCreatingFailed]
	}
	return association, nil
}

func (a *Association) GetAssociationById(ctx context.Context, id uuid.UUID) (models.Association, error) {
	var assoc models.Association
	db := a.storage.DB.WithContext(ctx).Where("id = ?", id).Find(&assoc)
	if db.Error != nil {
		a.logger.Err(db.Error).Msgf("Association::GetAssociationById error: %v, (%v)", language.ErrText()[language.ErrRecordNotFound], db.Error)
		return assoc, language.ErrText()[language.ErrRecordNotFound]
	}
	return assoc, nil
}

func (a *Association) GetAllAssociations(ctx context.Context, page models.Page) ([]*models.Association, *models.PageInfo, error) {
	var associations []*models.Association
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

	queryDraft := a.storage.DB.WithContext(ctx).Model(models.Association{})
	dbCount := a.storage.DB.WithContext(ctx).Model(models.Association{})

	db := queryDraft.Offset(offset).Limit(*page.Size).
		Order(fmt.Sprintf("associations.%s %s", *page.SortBy, sortDirection)).
		Find(&associations)
	if db.Error != nil {
		a.logger.Err(db.Error).Msgf("Association::GetAllAssociations error: %v, (%v)", language.ErrText()[language.ErrRecordNotFound], db.Error)
		return nil, nil, language.ErrText()[language.ErrRecordNotFound]
	}

	// then do counting
	var count int64
	dbCount.Count(&count)

	return associations, &models.PageInfo{
		Page:            *page.Number,
		Size:            *page.Size,
		HasNextPage:     int64(offset+*page.Size) < count,
		HasPreviousPage: *page.Number > 1,
		TotalCount:      count,
	}, nil
}
