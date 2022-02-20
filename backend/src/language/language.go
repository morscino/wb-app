package language

import "strings"

const (
	ErrInvalidEmail         string = "ErrInvalidEmail"
	ErrInvalidPassword      string = "ErrInvalidPassword"
	ErrPasswordDoesNotMatch string = "ErrPasswordDoesNotMatch"
	ErrRecordCreatingFailed string = "ErrRecordCreatingFailed"
	ErrDuplicateRecord      string = "ErrDuplicateRecord"
	ErrRecordNotFound       string = "ErrRecordNotFound"
	ErrEmailAlreadyExist    string = "ErrEmailAlreadyExist"

	LanguageEnglish string = "english"
)

func ErrText() map[string]error {
	// TODO: this to be made dynamic
	lang := "english"
	if strings.EqualFold(lang, LanguageEnglish) {
		return ErrEnglish
	} else {
		return ErrFrench
	}

}
