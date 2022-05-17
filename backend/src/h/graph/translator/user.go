package translator

import (
	"database/sql"
	"regexp"
	"strings"
	"time"

	"github.com/MastoCred-Inc/web-app/h/graph/model"
	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/MastoCred-Inc/web-app/utility/helper"
	"github.com/google/uuid"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ConvertUserInputToUserModel(userInput model.RegisterUser) (*models.User, error) {
	var (
		t       int64
		regDate time.Time
		err     error
		assocID uuid.UUID
	)

	if !emailIsValid(strings.ToLower(userInput.Email)) {
		return nil, language.ErrText()[language.ErrInvalidEmail]
	}

	//ensure password is valid
	lenghtOrMore, number, upper, special := passwordIsValid(userInput.Password)

	if !lenghtOrMore || !number || !upper || !special {
		return nil, language.ErrText()[language.ErrInvalidPassword]
	}

	if userInput.Password != userInput.PasswordMatch {
		return nil, language.ErrText()[language.ErrPasswordDoesNotMatch]
	}

	userTypeInput := string(userInput.UserType)

	// check the type of user being created and validate accordingly
	switch userTypeInput {

	case models.UserTypeSME.String():
		if userInput.BusinessName == nil || userInput.BusinessRCNumber == nil {
			return nil, language.ErrText()[language.ErrBusinessNameRcNumberIsRequired]
		}
		t = models.UserTypeSME.Value()
	case models.UserTypeAdmin.String():
		t = models.UserTypeAdmin.Value()
	case models.UserTypeIndividual.String():
		t = models.UserTypeIndividual.Value()
	}
	if userInput.BusinessRegistrationDate != nil {
		regDate, err = helper.StringToTime(*userInput.BusinessRegistrationDate)
		if err != nil {
			return nil, language.ErrText()[language.ErrParseError]
		}
	}
	if userInput.AssociationID != nil {
		assocID, err = helper.StringToUuid(*userInput.AssociationID)
		if err != nil {
			return nil, language.ErrText()[language.ErrParseError]
		}
	}

	user := models.User{
		LastName:                 userInput.Lastname,
		FirstName:                userInput.Firstname,
		Email:                    strings.ToLower(userInput.Email),
		Password:                 userInput.Password,
		UserType:                 t,
		Telephone:                userInput.PhoneNumber,
		BusinessName:             userInput.BusinessName,
		BusinessRCNumber:         userInput.BusinessRCNumber,
		BusinessRegistrationDate: &sql.NullTime{Time: regDate, Valid: true},
		AssociationID:            &assocID,
		AssociationBranch:        userInput.AssociationBranch,
	}
	return &user, nil
}

func ConvertUpdateUserInputToUserModel(userInput model.UserKYCRequest, userID uuid.UUID) (*models.User, error) {
	var (
		DOB, regDate                       time.Time
		err                                error
		maritalStatusInput, meansOfIDInput string
		m                                  int64
		assocID                            uuid.UUID
		salary                             float64
	)

	if userInput.MeansOfIdentification != nil {
		meansOfIDInput = string(*userInput.MeansOfIdentification)
	}

	// check the type of user being created and validate accordingly
	switch meansOfIDInput {

	case models.MeansOfIdentificationDriversLicense.String():
		m = models.MeansOfIdentificationDriversLicense.Value()
	case models.MeansOfIdentificationIntlPassport.String():
		m = models.MeansOfIdentificationIntlPassport.Value()
	case models.MeansOfIdentificationNIN.String():
		m = models.MeansOfIdentificationNIN.Value()
	}

	if userInput.MaritalStatus != nil {
		maritalStatusInput = userInput.MaritalStatus.String()
		if strings.EqualFold(maritalStatusInput, model.UserMaritalStatusPreferNotToSay.String()) {
			maritalStatusInput = "prefer not to say"
		}
	}

	if userInput.DateOfBirth != nil {
		DOB, err = helper.StringToTime(*userInput.DateOfBirth)
		if err != nil {
			return nil, language.ErrText()[language.ErrParseError]
		}
	}
	if userInput.BusinessRegistrationDate != nil {
		regDate, err = helper.StringToTime(*userInput.BusinessRegistrationDate)
		if err != nil {
			return nil, language.ErrText()[language.ErrParseError]
		}
	}
	if userInput.AssociationID != nil {
		assocID, err = helper.StringToUuid(*userInput.AssociationID)
		if err != nil {
			return nil, language.ErrText()[language.ErrParseError]
		}
	}

	if userInput.Salary != nil {
		salary = *userInput.Salary
		if err != nil {
			return nil, language.ErrText()[language.ErrParseError]
		}
	}

	//userID, err = helper.StringToUuid(*&userInput.ID)
	if err != nil {
		return nil, language.ErrText()[language.ErrParseError]
	}

	user := models.User{
		ID:                       userID,
		Occupation:               userInput.Occupation,
		Salary:                   &salary,
		DateOfBirth:              &sql.NullTime{Time: DOB, Valid: true},
		MaritalStatus:            &maritalStatusInput,
		MeansOfIdentification:    &m,
		BusinessName:             userInput.BusinessName,
		BusinessRCNumber:         userInput.BusinessRCNumber,
		BusinessRegistrationDate: &sql.NullTime{Time: regDate, Valid: true},
		AssociationID:            &assocID,
		AssociationBranch:        userInput.AssociationBranch,
		State:                    userInput.State,
		LocalGovernment:          userInput.LocalGovernment,
		BVN:                      userInput.Bvn,
	}

	return &user, nil
}
