package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"gitlab.com/mastocred/web-app/h/graph/generated"
	"gitlab.com/mastocred/web-app/h/graph/model"
	"gitlab.com/mastocred/web-app/h/graph/translator"
	"gitlab.com/mastocred/web-app/integrations/upload"
	"gitlab.com/mastocred/web-app/language"
	"gitlab.com/mastocred/web-app/models"
	"gitlab.com/mastocred/web-app/utility/helper"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUser) (*models.User, error) {
	// validate user input
	userModel, err := translator.ConvertUserInputToUserModel(input)
	if err != nil {
		return nil, err
	}

	// send user to controller
	user, err := r.controller.RegisterUser(ctx, *userModel)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) UserKyc(ctx context.Context, input model.UserKYCRequest) (*models.User, error) {
	var (
		profilePictureUrl, docFileUrl string
	)
	ginC, err := helper.GinContextFromContext(ctx)
	if err != nil {
		r.logger.Err(err).Msgf("UpdateUser:GinContextFromContext : (%v)", err)
		return nil, err
	}

	actorUser, err := r.controller.Middleware().PasetoUserAuth(ginC)
	if err != nil {
		r.logger.Err(err).Msgf("UpdateUser:PasetoUserAuth : (%v)", err)
		return nil, err
	}

	// validate user input
	userModel, err := translator.ConvertUpdateUserInputToUserModel(input, actorUser.ID)
	if err != nil {
		r.logger.Err(err).Msgf("UpdateUser:ConvertUpdateUserInputToUserModel : (%v)", err)
		return nil, err
	}

	if actorUser.ID != userModel.ID {
		return nil, language.ErrText()[language.ErrAccessDenied]
	}

	// uplpad files if they are attached
	if input.ProfilePictureFile != nil {
		profilePictureFile := input.ProfilePictureFile
		b, err := helper.StreamToByte(profilePictureFile.File)
		if err != nil {
			r.logger.Err(err).Msgf("UpdateUser file error: [%v] : (%v)", profilePictureFile.File, err)
			return nil, language.ErrText()[language.ErrInvalidFileUpload]
		}
		image := upload.FileInput{Content: b, Name: profilePictureFile.Filename, Kind: profilePictureFile.ContentType, Size: profilePictureFile.Size}
		maxFileSize, _ := strconv.ParseInt(r.env.Get("PROFILE_PICTURE_MAX_FILE_SIZE"), 10, 64)
		attachmentKinds := []upload.AttachmentKind{upload.AttachmentKindImageJPEG, upload.AttachmentKindImageJPG, upload.AttachmentKindImagePNG}
		uploadFolder := r.env.Get("PROFILE_PICTURES_FOLDER")

		profilePictureUrl, err = r.controller.UploadFileToAwsS3(image, maxFileSize, attachmentKinds, uploadFolder)
		if err != nil {
			return nil, err
		}
		userModel.ProfilePictureURL = &profilePictureUrl
	}

	if input.DocumentFile != nil {
		docFile := input.DocumentFile
		b, err := helper.StreamToByte(docFile.File)
		if err != nil {
			r.logger.Err(err).Msgf("UpdateUser file error: [%v] : (%v)", docFile.File, err)
			return nil, language.ErrText()[language.ErrInvalidFileUpload]
		}
		image := upload.FileInput{Content: b, Name: docFile.Filename, Kind: docFile.ContentType, Size: docFile.Size}
		maxFileSize, _ := strconv.ParseInt(r.env.Get("DOCUMENT_MAX_FILE_SIZE"), 10, 64)
		attachmentKinds := []upload.AttachmentKind{upload.AttachmentKindPDF}
		uploadFolder := r.env.Get("DOCUMENT_FILES_FOLDER")

		docFileUrl, err = r.controller.UploadFileToAwsS3(image, maxFileSize, attachmentKinds, uploadFolder)
		if err != nil {
			return nil, err
		}
		userModel.DocumentURL = &docFileUrl
	}

	// send user to controller
	user, err := r.controller.UpdateUserByID(ctx, *userModel)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) RegisterAssociation(ctx context.Context, name string) (*models.Association, error) {
	ginC, err := helper.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	actorUser, err := r.controller.Middleware().PasetoUserAuth(ginC)
	if err != nil {
		return nil, err
	}

	if actorUser.UserType != int64(models.UserTypeAdmin) {
		return nil, language.ErrText()[language.ErrAccessDenied]
	}
	// send association to controller
	created, err := r.controller.RegisterAssociation(ctx, models.Association{Name: name, ID: uuid.New()})
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (r *mutationResolver) CreateWaitList(ctx context.Context, input model.RegisterWaitlist) (bool, error) {
	var created bool
	// validate user input
	waitlist, err := translator.ConvertWaitlistInputToWaitlistModel(input)
	if err != nil {
		return false, err
	}

	// send waitlist to controller
	created, err = r.controller.CreateWaitlist(ctx, waitlist)
	if err != nil {
		return false, err
	}

	return created, nil
}

func (r *mutationResolver) LoanApplication(ctx context.Context, input model.LoanApplicationRequest) (*models.Loan, error) {
	ginC, err := helper.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := r.controller.Middleware().PasetoUserAuth(ginC)
	if err != nil {
		return nil, err
	}
	// prevent admin from applying for loans
	if user.UserType != int64(models.UserTypeIndividual) && user.UserType != int64(models.UserTypeSME) {
		return nil, language.ErrText()[language.ErrAccessDenied]
	}

	loan := models.Loan{
		RepaymentDuration: int64(input.RepaymentDuration),
		OtherLoansAmount:  input.OtherLoansAmount,
		LoanAmount:        input.LoanAmount,
		AccountNumber:     input.AccountNumber,
		AccountName:       input.AccountName,
		Bank:              input.Bank,
	}

	newLoan, err := r.controller.ApplyForLoan(ctx, loan, user.ID)
	if err != nil {
		return nil, err
	}

	return newLoan, nil
}

func (r *mutationResolver) ApproveLoanToggle(ctx context.Context, loanID string, loanStatus model.LoanStatusEnum) (bool, error) {
	ginC, err := helper.GinContextFromContext(ctx)
	if err != nil {
		return false, err
	}

	actorUser, err := r.controller.Middleware().PasetoUserAuth(ginC)
	if err != nil {
		return false, err
	}

	if actorUser.UserType != int64(models.UserTypeAdmin) {
		return false, language.ErrText()[language.ErrAccessDenied]
	}

	l, err := helper.StringToUuid(loanID)
	if err != nil {
		return false, language.ErrText()[language.ErrParseError]
	}

	_, err = r.controller.ApproveLoanToggle(ctx, l, string(loanStatus))
	if err != nil {
		return false, err
	}

	return true, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
