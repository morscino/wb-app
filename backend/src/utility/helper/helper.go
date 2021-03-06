package helper

import (
	"bytes"
	"context"
	"io"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/mastocred/web-app/language"
	"gitlab.com/mastocred/web-app/models"
)

const (
	// LogStrKeyModule log service name value
	LogStrKeyModule = "ser_name"
)

// GinContextFromContext gets a gin context from a context.Context
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("ctxkey")
	if ginContext == nil {
		return nil, language.ErrText()[language.ErrGinContextRetrieveFailed]
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, language.ErrText()[language.ErrGinContextWrongType]
	}

	return gc, nil
}

func ConvertStringPointerToString(s *string) string {
	return *s
}

func ConvertModeToIntPointer(s string) int {
	i := models.WaitListModeMap[s]
	return i
}
func StringToTime(s string) (time.Time, error) {
	return time.Parse("02/01/2006", s)
}

func StringToTimePointer(s string) (*time.Time, error) {
	t, err := time.Parse("02/01/2006", s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func StringToUuid(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func StringToFloat64(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}

func Float64ToString(f float64) string {
	s := strconv.FormatFloat(f, 'E', -1, 64)
	return s
}

func ConverInt64ToString(i int64) string {
	return strconv.Itoa(int(i))
}

// StreamToByte converts an io Stream to a slice of byte
func StreamToByte(stream io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(stream)

	return buf.Bytes(), nil
}
