package models

import (
	"time"

	"github.com/google/uuid"
)

type WaitlistMode string

const (
	WaitListModeIndividual WaitlistMode = "individual"
	WaitListModeBusiness   WaitlistMode = "business"
	WaitListModeAll        WaitlistMode = "all"
)

var WaitListModeMap = map[string]int{
	string(WaitListModeAll):        0,
	string(WaitListModeIndividual): 1,
	string(WaitListModeBusiness):   2,
}

type Waitlist struct {
	ID           uuid.UUID `gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	FullName     string
	BusinessName string
	Email        string
	Mode         int
	Telephone    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
