package repo

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	ID       int    `gorm:"primarykey"`
	Username string `gorm:"type:varchar(255);not null;unique"`
	Mail     string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varhcar(255);not null"`
	Salt     string `gorm:"type:varchar(255);not null"`
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

// func (userInfo *User) GetUser(db *gorm.DB, id int) (user User, err error) {
// 	result := db.First(&userInfo, id)

// 	if result.Error != nil {
// 		fmt.Println("User fetching failed due to -", result.Error)
// 	}

// 	fmt.Printf("User with id %d is ", id)
// 	return *userInfo, nil
// }
