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

//go:generate mockgen -source loan.go -destination ./mock/loan_user.go -package mock LoanStore
type LoanStore interface {
	CreateLoan(ctx context.Context, loan models.Loan) (models.Loan, error)
	GetAllLoans(ctx context.Context, page models.Page) ([]*models.Loan, *models.PageInfo, error)
}

// Loan object
type Loan struct {
	logger  zerolog.Logger
	storage *database.Storage
}

// NewLoan creates a new reference to the Loan storage entity
func NewLoan(s *database.Storage) *LoanStore {
	l := s.Logger.With().Str("app", "user").Logger()
	loan := &Loan{
		logger:  l,
		storage: s,
	}
	loanDB := LoanStore(loan)
	return &loanDB
}

// CreateLoan adds a new loan record to the database
func (l *Loan) CreateLoan(ctx context.Context, loan models.Loan) (models.Loan, error) {

	db := l.storage.DB.WithContext(ctx).Create(&loan)
	if db.Error != nil {
		l.logger.Err(db.Error).Msgf("Loan::CreateLoan error: %v, (%v)", language.ErrText()[language.ErrRecordCreatingFailed], db.Error)
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return models.Loan{}, language.ErrText()[language.ErrDuplicateRecord]
		}
		return models.Loan{}, language.ErrText()[language.ErrRecordCreatingFailed]
	}
	return loan, nil
}

func (l *Loan) GetAllLoans(ctx context.Context, page models.Page) ([]*models.Loan, *models.PageInfo, error) {
	var loans []*models.Loan
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

	queryDraft := l.storage.DB.WithContext(ctx).Model(models.Loan{})
	dbCount := l.storage.DB.WithContext(ctx).Model(models.Loan{})

	db := queryDraft.Offset(offset).Limit(*page.Size).
		Order(fmt.Sprintf("loans.%s %s", *page.SortBy, sortDirection)).
		Find(&loans)
	if db.Error != nil {
		l.logger.Err(db.Error).Msgf("Loan::GetAllLoans error: %v, (%v)", language.ErrText()[language.ErrRecordNotFound], db.Error)
		return nil, nil, language.ErrText()[language.ErrRecordNotFound]
	}

	// then do counting
	var count int64
	dbCount.Count(&count)

	return loans, &models.PageInfo{
		Page:            *page.Number,
		Size:            *page.Size,
		HasNextPage:     int64(offset+*page.Size) < count,
		HasPreviousPage: *page.Number > 1,
		TotalCount:      count,
	}, nil
}
