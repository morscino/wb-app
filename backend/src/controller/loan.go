package controller

import (
	"context"

	"github.com/MastoCred-Inc/web-app/language"
	"github.com/MastoCred-Inc/web-app/models"
	"github.com/MastoCred-Inc/web-app/utility/helper"
	"github.com/google/uuid"
)

func (c *Controller) ApplyForLoan(ctx context.Context, loan models.Loan, userID uuid.UUID) (*models.Loan, error) {
	// get user details
	loanUser, err := c.userStorage.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// check if various verifications have been done
	if loanUser.DocumentURL == nil {
		switch loanUser.UserType {
		case int64(models.UserTypeIndividual):
			return nil, language.ErrText()[language.ErrMeansOfIdRequiredForLoan]
		case int64(models.UserTypeSME):
			return nil, language.ErrText()[language.ErrCACDocRequiredForLoan]
		}
	}
	// if loanUser.BVN == nil {
	// 	return nil, language.ErrText()[language.ErrBvnRequiredForLoan]
	// }

	if loanUser.ProfilePictureURL == nil {
		return nil, language.ErrText()[language.ErrProfilePictureRequiredForLoan]
	}

	if loanUser.Salary == nil {
		return nil, language.ErrText()[language.ErrSalaryRequiredForLoan]
	}
	interestRate, _ := helper.StringToFloat64(c.env.Get("INTEREST_RATE"))
	processingFee, _ := helper.StringToFloat64(c.env.Get("PROCESSING_FEE"))
	balance := loan.GetTotalLoan(interestRate, processingFee)

	// apply for loan
	loanObject := models.Loan{
		UserID:            userID,
		RepaymentDuration: loan.RepaymentDuration,
		OtherLoansAmount:  loan.OtherLoansAmount,
		LoanAmount:        loan.LoanAmount,
		AccountNumber:     loan.AccountNumber,
		AccountName:       loan.AccountName,
		Bank:              loan.Bank,
		RepaymentStatus:   models.RepaymentStatusNotPaid,
		Status:            models.LoanStatusPending,
		Balance:           &balance,
		AmountPaid:        nil,
	}

	newLoan, err := c.loanStorage.CreateLoan(ctx, loanObject)
	if err != nil {
		return nil, err
	}

	return &newLoan, nil
}

func (c *Controller) GetAllLoans(ctx context.Context, page models.Page) ([]*models.Loan, *models.PageInfo, error) {
	return c.loanStorage.GetAllLoans(ctx, page)
}
