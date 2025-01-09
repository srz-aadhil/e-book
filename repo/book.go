package repo

import (
	"ebookmod/app/dto"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BookRepo interface {
	CreateBook(bookReq *dto.BookCreateRequest) (lastInsertedID int, err error)
	GetBook(id int) (bookResp *dto.BookResponse, err error)
	GetAllBooks()(bookResp []*dto.BookResponse,err error)
	UpdateBook(updateBook *dto.BookUpdateRequest) (err error)
	DeleteBook(id int)(err error)
}

type Book struct {
	ID        int            `gorm:"primarykey"`
	Title     string         `gorm:"title"`
	Content   string         `gorm:"content"`
	AuthorID  int            `gorm:"author_id"`
	CreatedBy int            `gorm:"created_by"`
	UpdatedBy *int           `gorm:"updated_by"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdateAt  time.Time      `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy *int           `gorm:"index"`
	Status    int            `gorm:"status"`
}

// Create a book
func (bookinfo *Book) CreateBook(db *gorm.DB) (id int, err error) {
	result := db.Table("books").Create(&bookinfo)

	if result.Error != nil {
		return 0, fmt.Errorf("Book creation failed due to- %v", result.Error)
	}

	fmt.Println("Book Successfully uploaded")
	return bookinfo.ID, nil
}

// Read a book
func GetaBook(db *gorm.DB, id int) (singlebook *Book, err error) {
	var book *Book
	//getting a book from db
	result := db.Where("id = ?", id).First(&book)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Book with id '%d' not found", id)
		}
		return nil, fmt.Errorf("failed to fetch boo due to %s ", result.Error)
	}

	if book.Status == 3 {
		return nil, fmt.Errorf("\nBook with id '%d' is deleted", id)
	} else {
		// 	fmt.Printf("Book with id '%d' details are\n Title:- %s\nContent:- %s\n Author:- %d\n Createdby:- %d\n Created at:- %v\n", book.ID, book.Title, book.Content, book.AuthorID, book.CreatedBy, book.CreatedAt)
		// }

		return book, nil
	}
}

// Read All Books
func GetAllBooks(db *gorm.DB) ([]*Book, error) {
	allbooks := []*Book{}
	result := db.Unscoped().Where("status = ? OR status = ?", 1, 2).Find(&allbooks)

	if result.Error != nil {
		return nil, fmt.Errorf("Books fetching failed due to %v", result.Error)
	}

	return allbooks, nil

}

// Update a book
func UpdateABook(db *gorm.DB, book *Book) error {
	result := db.Where("status = ? or status = ?", 1, 2).Updates(book)

	if result.Error != nil {
		return fmt.Errorf("Book updation failed due to- %v", result.Error)
	}
	//Checking if record of book is found
	if result.RowsAffected == 0 {
		return fmt.Errorf("Updation incomplete, book record not found")
	}

	fmt.Println("Book Updation Completed")
	return nil
}

// Delete a book
func DeleteaABook(db *gorm.DB, id int) error {
	//Query execution
	result := db.Table("books").Where("id = ? AND status = ? OR status = ?", id, 1, 2).Delete(&Book{})

	if result.Error != nil {
		return fmt.Errorf("Book deletion failed due to %v", result.Error)
	}

	updatedRecord := db.Table("books").Where("id = ? AND status = ? OR status = ?", id, 1, 2).Update("status", 3)

	if updatedRecord.Error != nil {
		return fmt.Errorf("Updating deletion status of book failed due to %v", updatedRecord.Error)
	}

	// if result.RowsAffected == 0 {
	// 	return fmt.Errorf("No book found with id '%d' to delete", id)
	// }

	fmt.Println("Book deletion successfully completed")
	return nil
}
