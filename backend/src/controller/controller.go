package controller

import (
	"context"

	"github.com/MastoCred-Inc/web-app/database"
	"github.com/MastoCred-Inc/web-app/integrations/upload"
	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/middleware"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/MastoCred-Inc/web-app/storage"
	"github.com/MastoCred-Inc/web-app/utility/environment"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

//go:generate mockgen -source controller.go -destination ./mock/mock_controller.go -package mock Operations
type Operations interface {
	Middleware() *middleware.Middleware
	UploadFileToAwsS3(file upload.FileInput, maxSize int64, allowedFileTypes []upload.AttachmentKind, uploadFolder string) (string, error)

	RegisterUser(ctx context.Context, user models.User) (*models.User, error)
	UpdateUserByID(ctx context.Context, u models.User) (*models.User, error)
	CreateWaitlist(ctx context.Context, waitlist *models.Waitlist) (bool, error)
	GetAllWaitlists(ctx context.Context, page models.Page, mode int) ([]*models.Waitlist, *models.PageInfo, error)
	RegisterAssociation(ctx context.Context, a models.Association) (*models.Association, error)
	GetAllAssociations(ctx context.Context, page models.Page) ([]*models.Association, *models.PageInfo, error)
	GetAssociationById(ctx context.Context, id uuid.UUID) (models.Association, error)
	GetAllLoans(ctx context.Context, page models.Page) ([]*models.Loan, *models.PageInfo, error)

	ApplyForLoan(ctx context.Context, loan models.Loan, userID uuid.UUID) (*models.Loan, error)
	ApproveLoanToggle(ctx context.Context, loanID uuid.UUID, status string) (bool, error)

	GetLoanInstalmentsByUserID(ctx context.Context, userID uuid.UUID, page models.Page) ([]*models.LoanInstalment, *models.PageInfo, error)
}

type Controller struct {
	env                   *environment.Env
	uploadService         upload.Uploader
	logger                zerolog.Logger
	userStorage           storage.UserStore
	loanStorage           storage.LoanStore
	associationstorage    storage.AssociationStore
	waitlistStorage       storage.WaitlistStore
	loanInstalmentStorage storage.LoanInstalmentStore
	middleware            *middleware.Middleware
}

func New(l zerolog.Logger, s *database.Storage, middleware *middleware.Middleware) *Operations {
	user := storage.NewUser(s)
	waitlist := storage.NewWaitlist(s)
	assoc := storage.NewAssociation(s)
	loan := storage.NewLoan(s)
	upload := upload.NewUpload(l, s.Env)
	loanInstalment := storage.NewLoanInstalment(s)

	// build controller struct
	c := &Controller{
		logger:                l,
		uploadService:         *upload,
		userStorage:           *user,
		loanStorage:           *loan,
		waitlistStorage:       *waitlist,
		middleware:            middleware,
		loanInstalmentStorage: *loanInstalment,
		env:                   s.Env,
		associationstorage:    *assoc,
	}
	op := Operations(c)
	return &op
}

func (c *Controller) Middleware() *middleware.Middleware {
	return c.middleware
}

func (c *Controller) UploadFileToAwsS3(file upload.FileInput, maxFileSize int64, allowedFileTypes []upload.AttachmentKind, uploadFolder string) (string, error) {
	var fileURL string

	//validate image
	err := c.uploadService.ValidateFile(file, maxFileSize, allowedFileTypes) // validating the file(s)
	if err != nil {
		c.logger.Err(err).Msgf("UploadPartnerLogo:validateFile [%v] : (%v)", file.Name, err)
		return "", err
	}

	//upload file into destination  folder
	savedFile, err := c.uploadService.UploadFile(file, uploadFolder)
	if err != nil {
		c.logger.Err(err).Msgf("UploadFileToAwsS3:UploadFile file error: (%v)", err)
		return "", language.ErrText()[language.ErrFileUpload]
	}
	fileURL = savedFile.URL
	return fileURL, nil
}
