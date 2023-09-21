package model

type UserModel struct {
	Username  string `json:"username" gorm:"column:username"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
}
