package create_database_model

type User struct {
	Username  string `gorm:"primary_key"`
	FirstName string
	LastName  string
}

func (User) TableName() string {
	return "user"
}
