package translator

import (
	"regexp"
	"strings"
	"unicode"

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

func emailIsValid(email string) bool {
	if len(email) < 4 || len(email) > 254 {
		return false
	}
	return regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9]+\\.(?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString(email)
}

func passwordIsValid(password string) (lenghtOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}
	lenghtOrMore = letters >= 8
	return
}
