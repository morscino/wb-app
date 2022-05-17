// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/MastoCred-Inc/web-app/models"
)

type GetAssociationsRequest struct {
	Page *models.Page `json:"page"`
}

type GetAssociationsResult struct {
	Page  *models.PageInfo      `json:"page"`
	Items []*models.Association `json:"items"`
}

type GetLoansResult struct {
	Page  *models.PageInfo `json:"page"`
	Items []*models.Loan   `json:"items"`
}

type GetUsersResult struct {
	Page  *models.PageInfo `json:"page"`
	Items []*models.User   `json:"items"`
}

type GetWaitlistsRequest struct {
	Page *models.Page         `json:"page"`
	Mode *models.WaitlistMode `json:"mode"`
}

type GetWaitlistsResult struct {
	Page  *models.PageInfo   `json:"page"`
	Items []*models.Waitlist `json:"items"`
}

type LoanApplicationRequest struct {
	RepaymentDuration float64 `json:"repaymentDuration"`
	OtherLoansAmount  float64 `json:"otherLoansAmount"`
	LoanAmount        float64 `json:"loanAmount"`
	AccountNumber     string  `json:"accountNumber"`
	AccountName       string  `json:"accountName"`
	Bank              string  `json:"bank"`
}

type RegisterUser struct {
	Email                    string       `json:"email"`
	Lastname                 string       `json:"lastname"`
	Firstname                string       `json:"firstname"`
	Password                 string       `json:"password"`
	PasswordMatch            string       `json:"passwordMatch"`
	PhoneNumber              string       `json:"phoneNumber"`
	UserType                 UserTypeEnum `json:"userType"`
	AssociationID            *string      `json:"associationID"`
	AssociationBranch        *string      `json:"associationBranch"`
	BusinessName             *string      `json:"businessName"`
	BusinessRegistrationDate *string      `json:"businessRegistrationDate"`
	BusinessRCNumber         *string      `json:"businessRCNumber"`
}

type RegisterWaitlist struct {
	FullName     string              `json:"fullName"`
	Email        string              `json:"email"`
	Telephone    string              `json:"telephone"`
	BusinessName *string             `json:"businessName"`
	Mode         models.WaitlistMode `json:"mode"`
}

type UserAuthenticated struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

type UserKYCRequest struct {
	Lastname                 *string                    `json:"lastname"`
	Firstname                *string                    `json:"firstname"`
	PhoneNumber              *string                    `json:"phoneNumber"`
	AssociationID            *string                    `json:"associationID"`
	AssociationBranch        *string                    `json:"associationBranch"`
	BusinessName             *string                    `json:"businessName"`
	BusinessRegistrationDate *string                    `json:"businessRegistrationDate"`
	BusinessRCNumber         *string                    `json:"businessRCNumber"`
	Occupation               *string                    `json:"occupation"`
	Salary                   *float64                   `json:"salary"`
	DateOfBirth              *string                    `json:"dateOfBirth"`
	State                    *string                    `json:"state"`
	LocalGovernment          *string                    `json:"localGovernment"`
	Bvn                      *string                    `json:"bvn"`
	MaritalStatus            *UserMaritalStatus         `json:"maritalStatus"`
	MeansOfIdentification    *UserMeansOfIdentification `json:"meansOfIdentification"`
	ProfilePictureFile       *graphql.Upload            `json:"profilePictureFile"`
	DocumentFile             *graphql.Upload            `json:"documentFile"`
}

type UserMaritalStatus string

const (
	UserMaritalStatusSingle         UserMaritalStatus = "single"
	UserMaritalStatusSeperated      UserMaritalStatus = "seperated"
	UserMaritalStatusEngaged        UserMaritalStatus = "engaged"
	UserMaritalStatusPreferNotToSay UserMaritalStatus = "prefer_not_to_say"
)

var AllUserMaritalStatus = []UserMaritalStatus{
	UserMaritalStatusSingle,
	UserMaritalStatusSeperated,
	UserMaritalStatusEngaged,
	UserMaritalStatusPreferNotToSay,
}

func (e UserMaritalStatus) IsValid() bool {
	switch e {
	case UserMaritalStatusSingle, UserMaritalStatusSeperated, UserMaritalStatusEngaged, UserMaritalStatusPreferNotToSay:
		return true
	}
	return false
}

func (e UserMaritalStatus) String() string {
	return string(e)
}

func (e *UserMaritalStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserMaritalStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserMaritalStatus", str)
	}
	return nil
}

func (e UserMaritalStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserMeansOfIdentification string

const (
	UserMeansOfIdentificationDriversLicense UserMeansOfIdentification = "drivers_license"
	UserMeansOfIdentificationNin            UserMeansOfIdentification = "NIN"
	UserMeansOfIdentificationIntlPassport   UserMeansOfIdentification = "intl_passport"
)

var AllUserMeansOfIdentification = []UserMeansOfIdentification{
	UserMeansOfIdentificationDriversLicense,
	UserMeansOfIdentificationNin,
	UserMeansOfIdentificationIntlPassport,
}

func (e UserMeansOfIdentification) IsValid() bool {
	switch e {
	case UserMeansOfIdentificationDriversLicense, UserMeansOfIdentificationNin, UserMeansOfIdentificationIntlPassport:
		return true
	}
	return false
}

func (e UserMeansOfIdentification) String() string {
	return string(e)
}

func (e *UserMeansOfIdentification) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserMeansOfIdentification(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserMeansOfIdentification", str)
	}
	return nil
}

func (e UserMeansOfIdentification) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserTypeEnum string

const (
	UserTypeEnumIndividual UserTypeEnum = "individual"
	UserTypeEnumSme        UserTypeEnum = "sme"
	UserTypeEnumAdmin      UserTypeEnum = "admin"
	UserTypeEnumSuperAdmin UserTypeEnum = "super_admin"
	UserTypeEnumGroup      UserTypeEnum = "group"
)

var AllUserTypeEnum = []UserTypeEnum{
	UserTypeEnumIndividual,
	UserTypeEnumSme,
	UserTypeEnumAdmin,
	UserTypeEnumSuperAdmin,
	UserTypeEnumGroup,
}

func (e UserTypeEnum) IsValid() bool {
	switch e {
	case UserTypeEnumIndividual, UserTypeEnumSme, UserTypeEnumAdmin, UserTypeEnumSuperAdmin, UserTypeEnumGroup:
		return true
	}
	return false
}

func (e UserTypeEnum) String() string {
	return string(e)
}

func (e *UserTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserTypeEnum", str)
	}
	return nil
}

func (e UserTypeEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
