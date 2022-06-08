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

//go:generate mockgen -source loan_instalment.go -destination ./mock/loan_instalment.go -package mock LoanInstalmentStore
type LoanInstalmentStore interface {
	CreateLoanInstalment(ctx context.Context, loanInstalment models.LoanInstalment) (models.LoanInstalment, error)
	GetAllLoanInstalments(ctx context.Context, page models.Page) ([]*models.LoanInstalment, *models.PageInfo, error)
	GetInstalmentByLoanID(ctx context.Context, loanID uuid.UUID) (*models.LoanInstalment, error)
	GetLoanInstalmentsByUserID(ctx context.Context, userID uuid.UUID, page models.Page) ([]*models.LoanInstalment, *models.PageInfo, error)
}

// LoanInstalment object
type LoanInstalment struct {
	logger  zerolog.Logger
	storage *database.Storage
}

// NewLoan creates a new reference to the Loan storage entity
func NewLoanInstalment(s *database.Storage) *LoanInstalmentStore {
	l := s.Logger.With().Str("app", "user").Logger()
	loanInstalment := &LoanInstalment{
		logger:  l,
		storage: s,
	}
	loanInstalmentDB := LoanInstalmentStore(loanInstalment)
	return &loanInstalmentDB
}

// CreateLoanInstalment adds a new loan instalment record to the database
func (l *LoanInstalment) CreateLoanInstalment(ctx context.Context, loanInstalment models.LoanInstalment) (models.LoanInstalment, error) {

	db := l.storage.DB.WithContext(ctx).Create(&loanInstalment)
	if db.Error != nil {
		l.logger.Err(db.Error).Msgf("LoanInstalment::CreateLoanInstalment error: %v, (%v)", language.ErrText()[language.ErrRecordCreatingFailed], db.Error)
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return models.LoanInstalment{}, language.ErrText()[language.ErrDuplicateRecord]
		}
		return models.LoanInstalment{}, language.ErrText()[language.ErrRecordCreatingFailed]
	}
	return loanInstalment, nil
}

func (l *LoanInstalment) GetAllLoanInstalments(ctx context.Context, page models.Page) ([]*models.LoanInstalment, *models.PageInfo, error) {
	var loanInstalments []*models.LoanInstalment
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

	queryDraft := l.storage.DB.WithContext(ctx).Model(models.LoanInstalment{})
	dbCount := l.storage.DB.WithContext(ctx).Model(models.LoanInstalment{})

	db := queryDraft.Offset(offset).Limit(*page.Size).
		Order(fmt.Sprintf("loan_instalments.%s %s", *page.SortBy, sortDirection)).
		Find(&loanInstalments)
	if db.Error != nil {
		l.logger.Err(db.Error).Msgf("LoanInstalment::GetAllLoanInstalments error: %v, (%v)", language.ErrText()[language.ErrRecordNotFound], db.Error)
		return nil, nil, language.ErrText()[language.ErrRecordNotFound]
	}

	// then do counting
	var count int64
	dbCount.Count(&count)

	return loanInstalments, &models.PageInfo{
		Page:            *page.Number,
		Size:            *page.Size,
		HasNextPage:     int64(offset+*page.Size) < count,
		HasPreviousPage: *page.Number > 1,
		TotalCount:      count,
	}, nil
}

func (l *LoanInstalment) GetInstalmentByLoanID(ctx context.Context, loanID uuid.UUID) (*models.LoanInstalment, error) {
	var loanInstalment models.LoanInstalment
	db := l.storage.DB.WithContext(ctx).Where("loan_id = ?", loanID).Find(&loanInstalment)
	if db.Error != nil {
		l.logger.Err(db.Error).Msgf("LoanInstalment::GetInstalmentByLoanID error: %v, (%v)", language.ErrText()[language.ErrRecordNotFound], db.Error)
		return nil, language.ErrText()[language.ErrRecordNotFound]
	}
	return &loanInstalment, nil
}

func (l *LoanInstalment) GetLoanInstalmentsByUserID(ctx context.Context, userID uuid.UUID, page models.Page) ([]*models.LoanInstalment, *models.PageInfo, error) {
	var loanInstalments []*models.LoanInstalment
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
		tmpPageSortBy := "repayment_date"
		page.SortBy = &tmpPageSortBy
	}
	if page.SortDirectionDesc == nil {
		tmpPageSortDirectionDesc := models.PageDefaultSortDirectionDesc
		page.SortDirectionDesc = &tmpPageSortDirectionDesc
	}

	if *page.Number > 1 {
		offset = *page.Size * (*page.Number - 1)
	}
	sortDirection := models.PageSortDirectionAscending
	if !*page.SortDirectionDesc {
		sortDirection = models.PageSortDirectionAscending
	}

	queryDraft := l.storage.DB.WithContext(ctx).Model(models.LoanInstalment{}).Where("user_id = ?", userID)
	dbCount := l.storage.DB.WithContext(ctx).Model(models.LoanInstalment{}).Where("user_id = ?", userID)

	db := queryDraft.Offset(offset).Limit(*page.Size).
		Order(fmt.Sprintf("loan_instalments.%s %s", *page.SortBy, sortDirection)).
		Find(&loanInstalments)
	if db.Error != nil {
		l.logger.Err(db.Error).Msgf("LoanInstalment::GetLoanInstalmentsByUserID error: %v, (%v)", language.ErrText()[language.ErrRecordNotFound], db.Error)
		return nil, nil, language.ErrText()[language.ErrRecordNotFound]
	}

	// then do counting
	var count int64
	dbCount.Count(&count)

	return loanInstalments, &models.PageInfo{
		Page:            *page.Number,
		Size:            *page.Size,
		HasNextPage:     int64(offset+*page.Size) < count,
		HasPreviousPage: *page.Number > 1,
		TotalCount:      count,
	}, nil
}
