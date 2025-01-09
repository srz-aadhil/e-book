package service

import (
	"ebookmod/repo"
	"fmt"
)

type AuthorService struct {
	authorRepo repo.AuthorRepo
}

// NewAuthorService creates a new instance of AuthorService
func NewAuthorService(authorRepo repo.AuthorRepo) *AuthorService {
	return &AuthorService{
		authorRepo: authorRepo,
	}
}

func (service *AuthorService) CreateAuthorService() (id int, err error) {
	if author.AuthorName == "" {
		return 0, fmt.Errorf("Author name cannot be empty")
	}

	id, err = author.Crea
	if err != nil {
		return 0, fmt.Errorf("Author creation failed due to - %v", err)
	}

	return id, nil
}

func (service *AuthorService) GetAuthorService(id int) (author *repo.Author, err error) {
	if id <= 0 {
		return nil, fmt.Errorf("Invalid Author ID input")
	}

	author, err = repo.GetAuthor(service.db, id)
	if err != nil {
		return nil, fmt.Errorf("Author fetching failed due to %v", err)
	}

	return author, nil
}

func (service *AuthorService) GetAllAuthorsService() (authors []*repo.Author, err error) {
	authors, err = repo.GetAllAuthors(service.db)
	if err != nil {
		return nil, fmt.Errorf("All Autthors list fetching failed due to %v ", err)
	}
	return authors, nil
}

func (service *AuthorService) UpdateAuthorService(author *repo.Author) error {
	if author.ID <= 0 {

		return fmt.Errorf("Invalid author ID")
	}
	if author.AuthorName == "" {
		return fmt.Errorf("author name cannot be empty")
	}

	if err := repo.UpdateAuthor(service.db, author); err != nil {
		return fmt.Errorf("Author updating failed due to %v", err)
	}

	return nil

}

func (service *AuthorService) DeleteAuthorService(id int) error {
	if id <= 0 {
		return fmt.Errorf("Invalid author ID ")
	}

	if err := repo.DeleteAuthor(service.db, id); err != nil {
		return fmt.Errorf("Author deletion failed due to %v", err)
	}

	return nil
}
