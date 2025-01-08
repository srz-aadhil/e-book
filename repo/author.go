package repo

import (
	"ebookmod/app/dto"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type AuthorRepo interface {
	CreateAuthor(db *gorm.DB) (lastInsertedID int, err error)
	GetAuthor(db *gorm.DB, id int) (*Author, error)
	GetAllAuthors(db *gorm.DB) ([]*Author, error)
	UpdateAuthor(db *gorm.DB, author *Author) error
	DeleteAuthor(db *gorm.DB, id int) error
}
type Author struct {
	ID         int            `gorm:"primary-key"`
	AuthorName string         `gorm:"column:name"`
	CreatedBy  int            `gorm:"column:created_by"`
	UpdatedBy  *int           `gorm:"column:updated_by"`
	CreatedAt  time.Time      `gorm:"created_at"`
	UpdateAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	DeletedBy  *int64         `gorm:"index"`
	Status     bool           `gorm:"status"`
}

type AuthorRepoImpl struct {
	db *gorm.DB
}

var _ AuthorRepo = (*AuthorRepoImpl)(nil)

func NewAuthorRepo(db *gorm.DB) *AuthorRepo {
	return &AuthorRepoImpl{
		db: db,
	}
}

// Create an Author
func (r *AuthorRepoImpl) CreateAuthor(authorReq *dto.AuthorCreateRequest) (lastInsertedID int, err error) {
	result := r.db.Table("authors").Create(authorReq)
	if result.Error != nil {
		return 0, fmt.Errorf("Author creation failed due to- %v ", result.Error)
	}

	return lastInsertedID, nil
}

// Read a single user
func GetAuthor(db *gorm.DB, id int) (*Author, error) {
	authordetails := &Author{}
	result := db.Unscoped().First(&authordetails, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Println("NO user found") // Print message for non-existent user
			return nil, fmt.Errorf("User not found")
		} else {
			return nil, fmt.Errorf("User fetching failed due to -%v", result.Error)
		}
	}

	if authordetails.Status == false {
		return nil, fmt.Errorf("User already deleted")
	}
	return authordetails, nil
}

// Read All Authors
func GetAllAuthors(db *gorm.DB) ([]*Author, error) {
	allAuthors := []*Author{}
	result := db.Table("authors").Where("status = ?", true).Find(&allAuthors)

	if result.Error != nil {
		return nil, fmt.Errorf("Authors fetching failed due to- %v ", result.Error)
	}

	for _, author := range allAuthors {
		fmt.Printf("Author details with authorid '%d'- AuthorName:- %s Createdby:- '%d' CreatedAt:- %v \n", author.ID, author.AuthorName, author.CreatedBy, author.CreatedAt)
	}
	return allAuthors, nil
}

// Update Author
func UpdateAuthor(db *gorm.DB, author *Author) error {
	result := db.Table("authors").Where("id = ? AND status = ?", author.ID, true).Updates(author)

	if result.Error != nil {
		return fmt.Errorf("Author Updation failed due to- %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("No author with id '%d' exist to update ", author.ID)
	}

	fmt.Println("Author updation successfully completed")
	return nil
}

// Delete Author
func DeleteAuthor(db *gorm.DB, id int) error {
	result := db.Table("authors").Where("id = ? AND status = ?", id, true).Delete(&Author{})

	if result.Error != nil {
		return fmt.Errorf("Author Deletion failed due to %v", result.Error)
	}

	updateRecord := db.Table("authors").Where("id = ? AND status = ?", id, true).Update("status", false)

	if updateRecord.Error != nil {
		return fmt.Errorf("Author status updation failed due to %v", updateRecord.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("No Author found with id '%d' to delete", id)
	}

	fmt.Printf("Author with id '%d' deletion successfully completed", id)
	return nil
}
