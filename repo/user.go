package repo

import (
	"gorm.io/gorm"
)

type User struct {
	ID       int
	Username string
	Mail     string
	Password string
	Salt     string
	BaseModel
}

// Create User
func (userInfo *User) CreateUser(db *gorm.DB) (Id int, err error) {
	result := db.Create(&userInfo)

	if result.Error != nil {
		return 0, result.Error
	}

	return userInfo.ID, nil
}
