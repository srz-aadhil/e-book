package repo

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        int            `gorm:"primarykey"`
	Title     string         `gorm:"title"`
	Content   string         `gorm:"content"`
	AuthorID  int            `gorm:"author_id"`
	CreatedBy int            `gorm:"column:created_by"`
	UpdatedBy *int           `gorm:"column:updated_by"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdateAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy *int           `gorm:"index"`
	Status    bool           `gorm:"status"`
}

func (bookinfo *Book) CreateBook(db *gorm.DB) (id int, err error) {
	result := db.Table("books").Create(&bookinfo)

	if result.Error != nil {
		return 0, fmt.Errorf("Book creation failed due to- %v", result.Error)
	}

	fmt.Println("Book Successfully uploaded")
	return bookinfo.ID, nil
}
