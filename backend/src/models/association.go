package models

import (
	"time"

	"github.com/google/uuid"
)

// create association model to only be cretaed by super admin
type Association struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
