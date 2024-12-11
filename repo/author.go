package repo

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID         int            `gorm:"primary-key"`
	AuthorName string         `gorm:"column:name"`
	CreatedBy  int            `gorm:"column:created_by"`
	UpdatedBy  int            `gorm:"column:updated_by"`
	CreatedAt  time.Time      `gorm:"created_at"`
	UpdateAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	DeletedBy  *int64         `gorm:"index"`
}

// Create an Author
func (authorInfo *Author) CreateAuthor(db *gorm.DB) (id int, err error) {
	result := db.Table("authors").Create(&authorInfo)
	if result.Error != nil {
		return 0, fmt.Errorf("Author creation failed due to- %v ", result.Error)
	}

	return authorInfo.ID, nil
}
