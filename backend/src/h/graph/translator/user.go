package translator

import (
	"regexp"
	"strings"

	"github.com/MastoCred-Inc/web-app/h/graph/model"
	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/models"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ConvertUserInputToUserModel(userInput model.RegisterUser) (*models.User, error) {

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

	user := models.User{
		LastName:  userInput.Lastname,
		FirstName: userInput.Firstname,
		Email:     strings.ToLower(userInput.Email),
		Password:  userInput.Password,
	}
	return &user, nil
}
