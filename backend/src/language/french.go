package language

import "errors"

var ErrFrench = map[string]error{
	"ErrInvalidEmail": errors.New("email is not valid ---- french version"),
}
