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

// Get user
func GetUser(db *gorm.DB, id int) (*User, error) {
	userdetails := &User{}
	result := db.Unscoped().First(&userdetails, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Print("User not found")
		} else {
			fmt.Println("User fetching failed due to -", result.Error)
			return nil, result.Error
		}
	}
	if userdetails.IsDeleted {
		return nil, fmt.Errorf("User not found because the user record is deleted")
	} else {
		fmt.Println("User retrieved successfully")
	}
	return userdetails, nil
}

func GetAllUsers(db *gorm.DB) ([]*User, error) {
	var users []*User
	result := db.Unscoped().Where("is_deleted = ?", false).Find(&users)

	if result.Error != nil {
		fmt.Println("Users fetching failed due to-", result.Error)
		return nil, result.Error
	}

	for _, user := range users {
		fmt.Printf("User details with userid %d are Username: %s Mail: %s Created at: %v Updated at: %v\n", user.ID, user.Username, user.Mail, user.CreatedAt, user.UpdateAt)
	}
	return users, nil
}

// Delete user
func DeleteUser(db *gorm.DB, id int) error {
	result := db.Table("users").Where("is_deleted = ?", false).Delete(id)

	if result != nil {
		fmt.Println("User deletion failed due to- ", result.Error)
		return result.Error
	}
	fmt.Printf("User with id %d is deleted successfully", id)
	return nil
}
