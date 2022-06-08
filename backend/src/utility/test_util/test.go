package test_util

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/mastocred/web-app/models"
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return dummy
}
