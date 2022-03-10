package test_util

import (
	"time"

	"github.com/MastoCred-Inc/web-app/models"
	"github.com/google/uuid"
)

// NewTestUser returns a test user object
func NewTestUser() models.User {
	dummy := models.User{
		ID:        uuid.New(),
		FirstName: "Ali",
		LastName:  "Daudu",
		Email:     "test@email.com",
		Password:  "somepassword",
		Salt:      "some$alt",
		Username:  "username",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return dummy
}
