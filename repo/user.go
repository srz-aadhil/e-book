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
	result := db.Table("users").Where("id = ?", id).Delete(&User{})
	if result.Error != nil {
		fmt.Println("User deletion failed due to- ", result.Error)
		return result.Error
	}
	// Update the is_deleted field to true for the specified user
	updaterecord := db.Table("users").Where("id = ? AND is_deleted = ?", id, false).Update("is_deleted", true)

	// Check for errors during the update operation
	if updaterecord.Error != nil {
		fmt.Println("User deletion failed due to:", updaterecord.Error)
		return updaterecord.Error
	}

	if result.RowsAffected == 0 {
		fmt.Printf("No user found with id %d to delete.\n", id)
		return fmt.Errorf("no user found with id %d to delete", id)
	}

	fmt.Printf("User with id %d is deleted successfully\n", id)
	return nil
}

func UpdateUser(db *gorm.DB, id int, user *User) error {
	result := db.Where("id = ? AND is_deleted = ?", id, false).Updates(user)

	if result.Error != nil {
		fmt.Println("User updating failed due to- ", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Printf("No user found with id %d to update.\n", id)
		return fmt.Errorf("no user found with id %d to update", id)
	}
	if result.RowsAffected == 0 {
		fmt.Printf("No user found with id %d to update.\n", id)
		return fmt.Errorf("no user found with id %d to update", id)
	}

	fmt.Printf("User with userid %d updation successfully completed", id)
	return nil
}
