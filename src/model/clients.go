package model

type Client struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Names     string `json:"names"`
	LastNames string `json:"lastNames" gorm:"column:last_names"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
