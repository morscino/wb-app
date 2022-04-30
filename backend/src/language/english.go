package language

import "errors"

var ErrEnglish = map[string]error{
	ErrInvalidEmail:                   errors.New("email is not valid"),
	ErrInvalidPassword:                errors.New("Invalid password, password must contain at least 8 characters, one letter, one number and one special character"),
	ErrPasswordDoesNotMatch:           errors.New("password does not match"),
	ErrRecordCreatingFailed:           errors.New("record creation failed"),
	ErrDuplicateRecord:                errors.New("duplicate record not allowed"),
	ErrRecordNotFound:                 errors.New("record not found"),
	ErrEmailAlreadyExist:              errors.New("user with supplied email already exists"),
	ErrIncorrectUsernameOrPassword:    errors.New("incorrect username or passwordf"),
	ErrGinContextRetrieveFailed:       errors.New("gin context retrieval failed"),
	ErrGinContextWrongType:            errors.New("gin context wrong type"),
	ErrInvalidPhoneNumber:             errors.New("the supplied phone number is invalid"),
	ErrBusinessNameIsRequired:         errors.New("a business name is required for a non-individual application"),
	ErrBusinessNameRcNumberIsRequired: errors.New("a business name, rc number and business registration date are required"),
	ErrParseError:                     errors.New("could not parse submitted data"),
	ErrAccessDenied:                   errors.New("access denied"),
	ErrInvalidToken:                   errors.New("invalid token"),
	ErrRecordUpdateFailed:             errors.New("record update failed"),
	ErrInvalidFileUpload:              errors.New("invalid file upload"),
	ErrFileUpload:                     errors.New("file upload error"),
}
