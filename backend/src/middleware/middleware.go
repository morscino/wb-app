package middleware

import (
	"context"

	"github.com/MastoCred-Inc/web-app/database"
	graphQLmodel "github.com/MastoCred-Inc/web-app/h/graph/model"
	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/MastoCred-Inc/web-app/storage"
	"github.com/MastoCred-Inc/web-app/utility/environment"
	"github.com/rs/zerolog"
)

type TokenMaker interface {
	CreateToken(email string) (string, error)
	VerifyToken(token string) (*Payload, error)
}

type Middleware struct {
	PasetoMaker TokenMaker
	logger      zerolog.Logger
	env         *environment.Env
	storage     *database.Storage
	userStorage storage.UserStore
}

func NewMiddleware(z zerolog.Logger, env environment.Env, s *database.Storage) (*Middleware, error) {
	l := z.With().Str("middleware", "api").Logger()
	paseto, err := NewPasetoMaker(env)
	if err != nil {
		return nil, err
	}
	userStore := storage.NewUser(s)
	m := &Middleware{
		PasetoMaker: paseto,
		logger:      l,
		env:         &env,
		storage:     s,
		userStorage: *userStore,
	}
	return m, nil
}

func (m *Middleware) AuthenticateUser(ctx context.Context, email, password string) (*graphQLmodel.UserAuthenticated, error) {
	// check if user exists
	user, err := m.userStorage.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, language.ErrText()[language.ErrIncorrectUsernameOrPassword]
	}
	// check if password is correct
	if !user.VerifyPassword(password, user.Salt, user.Password) {
		return nil, language.ErrText()[language.ErrIncorrectUsernameOrPassword]
	}
	// generate token
	token, err := m.PasetoMaker.CreateToken(email)
	// return response
	authUser := &graphQLmodel.UserAuthenticated{
		Token: token,
		User: &models.User{
			ID:        user.ID,
			LastName:  user.LastName,
			FirstName: user.FirstName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
	}
	return authUser, nil

}