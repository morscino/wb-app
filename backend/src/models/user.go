package models

import (
	"database/sql"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserType int64

	MaritalStatus string

	MeansOfIdentification int64

	// User holds the user object struct
	User struct {
		ID                       uuid.UUID `gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
		LastName                 string
		FirstName                string
		Email                    string
		Password                 string
		Telephone                string
		Salt                     string
		UserType                 int64
		AssociationID            *uuid.UUID
		AssociationBranch        *string
		BusinessName             *string
		BusinessRegistrationDate *sql.NullTime
		BusinessRCNumber         *string
		Occupation               *string
		SalaryRange              *string
		DateOfBirth              *sql.NullTime
		MaritalStatus            *string
		MeansOfIdentification    *int64
		ProfilePictureURL        *string
		DocumentURL              *string
		CreatedAt                time.Time
		UpdatedAt                time.Time
	}
)

var (
	UserTypeMap = map[UserType]string{
		UserTypeIndividual: "individual",
		UserTypeGroup:      "group",
		UserTypeSME:        "sme",
		UserTypeAdmin:      "admin",
		UserTypeSuperAdmin: "super admin",
	}

	MeansOfIdentificationMap = map[MeansOfIdentification]string{
		MeansOfIdentificationDriversLicense: "drivers licence",
		MeansOfIdentificationNIN:            "NIN",
		MeansOfIdentificationIntlPassport:   "intl passport",
	}
)

// password constants
const (
	SaltLen        = 32
	HashLen        = 14
	CharacterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWZYZ1234567890"

	UserTypeIndividual UserType = 1
	UserTypeGroup      UserType = 2
	UserTypeSME        UserType = 3
	UserTypeAdmin      UserType = 4
	UserTypeSuperAdmin UserType = 5

	MeansOfIdentificationDriversLicense MeansOfIdentification = 1
	MeansOfIdentificationNIN            MeansOfIdentification = 2
	MeansOfIdentificationIntlPassport   MeansOfIdentification = 3
)

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

func (u UserType) Value() int64 {
	return int64(u)
}

func (u UserType) String() string {
	return UserTypeMap[u]
}

func (u UserType) StringToInt(s string) int64 {
	var val int64
	for k, v := range UserTypeMap {
		if strings.EqualFold(v, s) {
			val = k.Value()
			break
		}
	}
	return val
}

func (m MeansOfIdentification) Value() int64 {
	return int64(m)
}

func (m MeansOfIdentification) String() string {
	return MeansOfIdentificationMap[m]
}

func (u MeansOfIdentification) StringToInt(s string) int64 {
	var val int64
	for k, v := range MeansOfIdentificationMap {
		if strings.EqualFold(v, s) {
			val = k.Value()
			break
		}
	}
	return val
}
