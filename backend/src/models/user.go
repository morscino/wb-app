package models

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// password constants
const (
	SaltLen        = 32
	HashLen        = 14
	CharacterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWZYZ1234567890"
)

type User struct {
	ID        uuid.UUID `gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	LastName  string
	FirstName string
	Email     string
	Password  string
	Salt      string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GenerateSalt generates salt for password encryption
func (u User) GenerateSalt() string {
	salt := make([]byte, SaltLen)

	for i := range salt {
		salt[i] = CharacterBytes[rand.Intn(len(CharacterBytes))]
	}

	return string(salt)
}

// EncyptPassword encrypts user password
func (u User) EncyptPassword(password, salt string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password+salt), HashLen)
	return string(bytes), err
}

// VerifyPassword verifies user encrypted password
func (u User) VerifyPassword(password, salt, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt))
	return err == nil
}
