package model

import "time"

type Loan struct {
	ID         uint      `json:"id"`
	ClientID   uint      `json:"clientId" gorm:"column:id_client"`
	EmployeeID uint      `json:"employeeId" gorm:"column:id_employee"`
	EditionID  uint      `json:"editionId" gorm:"column:id_edition"`
	LoanDate   time.Time `json:"loanDate" gorm:"column:loan_date"`
	ReturnDate time.Time `json:"returnDate" gorm:"column:return_date"`
}
