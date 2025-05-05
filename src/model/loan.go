package model

import "time"

type Loan struct {
	ID         uint      `json:"id"`
	ClientID   uint      `gorm:"column:id_client"`
	Client     Client    `json:"client"`
	LoanDate   time.Time `json:"loanDate" gorm:"column:loan_date"`
	ReturnDate time.Time `json:"returnDate" gorm:"column:return_date"`
}
