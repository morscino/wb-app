package translator

import (
	"strings"

	"github.com/MastoCred-Inc/web-app/h/graph/model"
	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/models"
)

func ConvertWaitlistInputToWaitlistModel(input model.RegisterWaitlist) (*models.Waitlist, error) {
	//validate email
	if !emailIsValid(strings.ToLower(input.Email)) {
		return nil, language.ErrText()[language.ErrInvalidEmail]
	}
	// validate phone niumber
	if !phoneNumberIsValid(strings.ToLower(input.Telephone)) {
		return nil, language.ErrText()[language.ErrInvalidPhoneNumber]
	}

	waitlist := &models.Waitlist{
		FullName:  input.FullName,
		Email:     input.Email,
		Telephone: input.Telephone,
		Mode:      models.WaitListModeMap[string(input.Mode)],
	}

	if strings.EqualFold(string(input.Mode), string(models.WaitListModeBusiness)) {
		if input.BusinessName == nil {
			return nil, language.ErrText()[language.ErrBusinessNameIsRequired]
		}
		waitlist.BusinessName = *input.BusinessName
	}
	return waitlist, nil

}
