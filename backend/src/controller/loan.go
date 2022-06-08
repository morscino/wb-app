package controller

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gitlab.com/mastocred/web-app/language"
	"gitlab.com/mastocred/web-app/models"
	"gitlab.com/mastocred/web-app/utility/helper"
)

func (c *Controller) ApplyForLoan(ctx context.Context, loan models.Loan, userID uuid.UUID) (*models.Loan, error) {
	// get user details
	loanUser, err := c.userStorage.GetUserByID(ctx, userID)
	if err != nil {
		c.logger.Err(err).Msgf("ApplyForLoan:GetUserByID [%v] : (%v)", loanUser.ID, err)
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
		c.logger.Err(err).Msgf("ApplyForLoan:CreateLoan [%v] : (%v)", newLoan.ID, err)
		return nil, err
	}
	return &newLoan, nil
}

func (c *Controller) GetAllLoans(ctx context.Context, page models.Page) ([]*models.Loan, *models.PageInfo, error) {
	return c.loanStorage.GetAllLoans(ctx, page)
}

func (c *Controller) ApproveLoanToggle(ctx context.Context, loanID uuid.UUID, status string) (bool, error) {
	interestRate, _ := helper.StringToFloat64(c.env.Get("INTEREST_RATE"))
	processingFee, _ := helper.StringToFloat64(c.env.Get("PROCESSING_FEE"))
	loan := models.Loan{
		Status:           status,
		LoanApprovalDate: sql.NullTime{Time: time.Now(), Valid: true},
	}
	_, err := c.loanStorage.UpdateLoanByID(ctx, loanID, loan)
	if err != nil {
		c.logger.Err(err).Msgf("ApproveLoanToggle:GetByLoanByID : (%v)", err)
		return false, err
	}

	updatedLoan, err := c.loanStorage.GetLoanByID(ctx, loanID)
	if err != nil {
		c.logger.Err(err).Msgf("ApproveLoanToggle:GetByLoanByID [%v] : (%v)", updatedLoan.ID, err)
		return false, err
	}

	// create instalments
	loanInstalmentDates := updatedLoan.GetRepaymentDates()
	fmt.Println(loanInstalmentDates)
	fmt.Println(updatedLoan.RepaymentDuration)
	fmt.Println(updatedLoan.LoanAmount)

	amount := updatedLoan.GetRepayment(interestRate, processingFee)
	for i := 0; i < len(loanInstalmentDates); i++ {
		loanInstalment := models.LoanInstalment{
			LoanID:            updatedLoan.ID,
			UserID:            updatedLoan.UserID,
			RepaymentAmount:   amount / float64(len(loanInstalmentDates)),
			RepaymentDate:     loanInstalmentDates[i],
			RepaymentDuration: updatedLoan.RepaymentDuration,
			RepaymentStatus:   models.RepaymentStatusNotPaid,
		}

		_, err = c.loanInstalmentStorage.CreateLoanInstalment(ctx, loanInstalment)
		if err != nil {
			c.logger.Err(err).Msgf("ApproveLoanToggle:CreateLoanInstalment [%v] : (%v)", updatedLoan.ID, err)
		}
	}

	return true, nil
}

func (c *Controller) GetLoanInstalmentsByUserID(ctx context.Context, userID uuid.UUID, page models.Page) ([]*models.LoanInstalment, *models.PageInfo, error) {
	return c.loanInstalmentStorage.GetLoanInstalmentsByUserID(ctx, userID, page)
}
