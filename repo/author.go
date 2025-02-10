package repo

import (
	"ebookmod/app/dto"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type AuthorRepo interface {
	CreateAuthor(authorReq *dto.AuthorCreateRequest) (lastInsertedID int, err error)
	GetAuthor(id int) (authorResp *Author, err error)
	GetAllAuthors() (authorResp []*Author, err error)
	UpdateAuthor(updateReq *dto.AuthorUpdateRequest) error
	DeleteAuthor(deleteReq *dto.AuthorDeleteRequest) error
}
type Author struct {
	ID        int            `gorm:"primary-key"`
	Name      string         `gorm:"column:name"`
	CreatedBy int            `gorm:"column:created_by"`
	UpdatedBy *int           `gorm:"column:updated_by"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdateAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy *int           `gorm:"index"`
	Status    bool           `gorm:"status"`
}

type AuthorRepoImpl struct {
	db *gorm.DB
}

func NewAuthorRepo(db *gorm.DB) AuthorRepo {
	return &AuthorRepoImpl{
		db: db,
	}
}

// Create an Author
func (r *AuthorRepoImpl) CreateAuthor(authorReq *dto.AuthorCreateRequest) (lastInsertedID int, err error) {
	author := &Author{
		Name:      authorReq.Name,
		CreatedBy: authorReq.CreatedBy,
		Status:    true,
	}
	result := r.db.Table("authors").Create(author)
	if result.Error != nil {
		return 0, fmt.Errorf("Author creation failed due to- %v ", result.Error)
	}

	return author.ID, nil
}

// Read a single user
func (r *AuthorRepoImpl) GetAuthor(id int) (authorResp *Author, err error) {
	result := r.db.Unscoped().First(&authorResp, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Println("NO author found") // Print message for non-existent user
			return nil, fmt.Errorf("Author not found")
		} else {
			return nil, fmt.Errorf("Author fetching failed due to -%v", result.Error)
		}
	}

	if authorResp.DeletedBy != nil {
		return nil, fmt.Errorf("Author already deleted")
	}
	return authorResp, nil
}

// Read All Authors
func (r *AuthorRepoImpl) GetAllAuthors() (authorResp []*Author, err error) {
	result := r.db.Table("authors").Where("status = ?", true).Find(&authorResp)

	if result.Error != nil {
		return nil, fmt.Errorf("Authors fetching failed due to- %v ", result.Error)
	}

	// for _, author := range authorResp {
	// 	fmt.Printf("Author details with authorid '%d'- AuthorName:- %s Createdby:- '%d' CreatedAt:- %v \n", author.ID, author.AuthorName, author.CreatedBy, author.CreatedAt)
	// }
	return authorResp, nil
}

// Update Author
func (r *AuthorRepoImpl) UpdateAuthor(updateReq *dto.AuthorUpdateRequest) error {
	result := r.db.Table("authors").Where("id = ? AND status = ?", updateReq.ID, true).Updates(map[string]interface{}{
		"name":       updateReq.Name,
		"updated_by": updateReq.UpdatedBy,
		"updated_at": time.Now().UTC(),
	})

	if result.Error != nil {
		return fmt.Errorf("Author Updation failed due to- %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("No author with id '%d' exist to update ", updateReq.ID)
	}

	fmt.Println("Author updation successfully completed")
	return nil
}

// Delete Author
func (r *AuthorRepoImpl) DeleteAuthor(deleteReq *dto.AuthorDeleteRequest) error {

	result := r.db.Table("authors").Where("id = ? AND status = ?", deleteReq.ID, true).Updates(map[string]interface{}{
		"status":     false,
		"deleted_by": deleteReq.DeletedBy,
		"deleted_at": time.Now().UTC(),
	})

	if result.Error != nil {
		return fmt.Errorf("Author Deletion failed due to %v", result.Error)
	}

	// updateRecord := r.db.Table("authors").Where("id = ? AND status = ?", id, true).Update("status", false)

	// if updateRecord.Error != nil {
	// 	return fmt.Errorf("Author status updation failed due to %v", updateRecord.Error)
	// }
	if result.RowsAffected == 0 {
		return fmt.Errorf("No Author found with id '%d' to delete", deleteReq.ID)
	}

	fmt.Printf("Author with id '%d' deletion successfully completed", deleteReq.ID)
	return nil
}
