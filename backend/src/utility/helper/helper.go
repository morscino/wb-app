package helper

import (
	"context"

	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/gin-gonic/gin"
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
