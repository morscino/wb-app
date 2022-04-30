package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MastoCred-Inc/web-app/h/graph/generated"
	"github.com/MastoCred-Inc/web-app/h/graph/model"
	"github.com/MastoCred-Inc/web-app/h/graph/translator"
	"github.com/MastoCred-Inc/web-app/integrations/upload"
	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/MastoCred-Inc/web-app/utility/helper"
	"github.com/google/uuid"
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

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserRequest) (*models.User, error) {
	var (
		profilePictureUrl, docFileUrl string
	)
	ginC, err := helper.GinContextFromContext(ctx)
	if err != nil {
		r.logger.Err(err).Msgf("UpdateUser:GinContextFromContext [%v] : (%v)", input.ID, err)
		return nil, err
	}

	actorUser, err := r.controller.Middleware().PasetoUserAuth(ginC)
	if err != nil {
		r.logger.Err(err).Msgf("UpdateUser:PasetoUserAuth : (%v)", err)
		return nil, err
	}

	// validate user input
	userModel, err := translator.ConvertUpdateUserInputToUserModel(input)
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
	}

	userModel.ProfilePictureURL = &profilePictureUrl
	userModel.DocumentURL = &docFileUrl

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

func (r *mutationResolver) ApplyLoan(ctx context.Context, input model.ApplyLoanRequest) (*models.Loan, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
