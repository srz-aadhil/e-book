package repo

import (
	"ebookmod/app/dto"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(userReq *dto.UserCreateRequest) (lastInsertedID int, err error)
	GetUser(id int) (userResp *User, err error)
	GetAllUsers() (userResp []*User, err error)
	UpdateUser(updateReq *dto.UserUpdateRequest) error
	DeleteUser(deleteReq *dto.UserDeleteRequest) error
}

type User struct {
	ID       int    `gorm:"primarykey"`
	Username string `gorm:"type:varchar(255);not null;unique"`
	Mail     string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varhcar(255);not null"`
	Salt     string `gorm:"type:varchar(255);not null"`
	BaseModel
}

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}

// Create User
func (r *UserRepoImpl) CreateUser(userReq *dto.UserCreateRequest) (lastInsertedID int, err error) {
	user := &User{
		Username: userReq.UserName,
		Password: userReq.Password,
		Salt:     userReq.Salt,
		Mail:     userReq.Mail,
	}
	result := r.db.Create(&user)

	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

// Read a single user
func (r *UserRepoImpl) GetUser(id int) (userResp *User, err error) {
	result := r.db.Unscoped().First(&userResp, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Print("User not found")
		} else {
			fmt.Println("User fetching failed due to -", result.Error)
			return nil, result.Error
		}
	}
	if userResp.IsDeleted {
		return nil, fmt.Errorf("User not found because the user record is deleted")
	} else {
		fmt.Println("User retrieved successfully")
	}
	return userResp, nil
}

// Read All Users
func (r *UserRepoImpl) GetAllUsers() (userResp []*User, err error) {
	result := r.db.Table("users").Where("is_deleted = ?", false).Find(&userResp)

	if result.Error != nil {
		log.Print("Users fetching failed due to-", result.Error)
		return nil, result.Error
	}

	return userResp, nil
}

// Update user
func (r *UserRepoImpl) UpdateUser(updateReq *dto.UserUpdateRequest) error {
	result := r.db.Table("users").Where("id = ? AND is_deleted = ?", updateReq.ID, false).Updates(map[string]interface{}{
		"username":   updateReq.UserName,
		"password":   updateReq.Password,
		"updated_at": time.Now().UTC(),
	})

	if result.Error != nil {
		fmt.Println("User updating failed due to- ", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		// fmt.Printf("No user found with id %d to update.\n", id)
		return fmt.Errorf("no user found with id %d to update", updateReq.ID)
	}
	fmt.Printf("User with userid %d updation successfully completed", updateReq.ID)
	return nil
}

// Delete user
func (r *UserRepoImpl) DeleteUser(deleteReq *dto.UserDeleteRequest) error {
	result := r.db.Table("users").Where("id = ? AND is_deleted = ?", deleteReq.ID, false).Updates(map[string]interface{}{
		"is_deleted": true,
		"deleted_at": time.Now().UTC(),
		"deleted_by": deleteReq.DeletedBy,
	})
	if result.Error != nil {
		fmt.Println("User deletion failed due to- ", result.Error)
		return result.Error
	}
	// Update the is_deleted field to true for the specified user
	//updaterecord := db.Table("users").Where("id = ? AND is_deleted = ?", id, false).Update("is_deleted", true)

	// Check for errors during the update operation
	// if updaterecord.Error != nil {
	// 	fmt.Println("User deletion failed due to:", updaterecord.Error)
	// 	return updaterecord.Error
	// }

	if result.RowsAffected == 0 {
		fmt.Printf("No user found with id %d to delete.\n", deleteReq.ID)
		return fmt.Errorf("no user found with id %d to delete", deleteReq.ID)
	}

	fmt.Printf("User with id %d is deleted successfully\n", deleteReq.ID)
	return nil
}
