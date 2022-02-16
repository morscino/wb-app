package models

import (
	"time"

	"github.com/google/uuid"
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
