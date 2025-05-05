package model

import "time"

type Fine struct {
	ID          uint      `json:"id"`
	LoanID      uint      `json:"loanId" gorm:"column:id_loan"`
	Amount      float32   `json:"amount"`
	Reason      *string   `json:"reason"`
	PaymentDate time.Time `json:"paymentDate" gorm:"column:date_payment"`
}
