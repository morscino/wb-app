package middleware

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"gitlab.com/mastocred/web-app/utility/environment"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto    *paseto.V2
	secretKey []byte
	env       *environment.Env
}

type Payload struct {
	ID        uuid.UUID
	Email     string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

func NewPayload(email string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return payload, nil
}
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiresAt) {
		return fmt.Errorf("token is expired")
	}
	return nil
}

func NewPasetoMaker(env *environment.Env) (TokenMaker, error) {
	secretKey := env.Get("TOKEN_SECRET_KEY")
	if len(secretKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size")
	}

	return &PasetoMaker{paseto: paseto.NewV2(), secretKey: []byte(secretKey)}, nil
}

func (p PasetoMaker) CreateToken(email string) (string, error) {
	duration, _ := time.ParseDuration(p.env.Get("TOKEN_EXPIRY"))
	payload, err := NewPayload(email, duration)
	if err != nil {
		return "", err
	}

	return p.paseto.Encrypt(p.secretKey, payload, nil)
}
func (p PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := p.paseto.Decrypt(token, p.secretKey, payload, nil)
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
