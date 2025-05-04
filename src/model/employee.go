package model

import "time"

type Employee struct {
	ID         uint      `json:"id"`
	Names      string    `json:"names"`
	LastNames  string    `json:"lastNames" gorm:"column:last_names"`
	HiringDate time.Time `json:"hiringDate" gorm:"column:hiring_date"`
}
