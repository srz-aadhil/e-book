package repo

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

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

// Create an Author
func (authorInfo *Author) CreateAuthor(db *gorm.DB) (id int, err error) {
	result := db.Table("authors").Create(&authorInfo)
	if result.Error != nil {
		return 0, fmt.Errorf("Author creation failed due to- %v ", result.Error)
	}

	return authorInfo.ID, nil
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

//Update Author
//Delete Author
