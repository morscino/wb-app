package translator

import (
	"regexp"
	"unicode"
)

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

func phoneNumberIsValid(phone string) bool {
	if len(phone) > 10 {
		phone = phone[len(phone)-10:]
	}

	return regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`).MatchString(phone)
}
