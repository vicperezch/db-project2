package model

type Author struct {
	ID          uint   `json:"id"`
	Names       string `json:"names"`
	LastNames   string `json:"lastNames" gorm:"column:last_names"`
	Nationality string `json:"nationality"`
}
