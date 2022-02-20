package language

import "errors"

var ErrEnglish = map[string]error{
	ErrInvalidEmail:         errors.New("email is not valid"),
	ErrInvalidPassword:      errors.New("Invalid password, password must contain at least 8 characters, one letter, one number and one special character"),
	ErrPasswordDoesNotMatch: errors.New("password does not match"),
	ErrRecordCreatingFailed: errors.New("record creation failed"),
	ErrDuplicateRecord:      errors.New("duplicate record not allowed"),
	ErrRecordNotFound:       errors.New("record not found"),
	ErrEmailAlreadyExist:    errors.New("user with supplied email already exists"),
}
