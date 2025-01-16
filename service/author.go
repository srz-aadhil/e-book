package service

import (
	"ebookmod/app/dto"
	"ebookmod/pkg/e"
	"ebookmod/repo"
	"net/http"
)

type AuthorService interface {
	CreateAuthor(r *http.Request) (lastInsertedID int, err error)
	GetAuthor(r *http.Request) (authorResp *dto.AuthorResponse, err error)
	GetAllAuthors() (authorsResp []*dto.AuthorResponse, err error)
	UpdateAuthor(r *http.Request) error
	DeleteAuthor(r *http.Request) error
}

type authorServiceImpl struct {
	authorRepo repo.AuthorRepo
}

// NewAuthorService creates a new instance of AuthorService
func NewAuthorService(authorRepo repo.AuthorRepo) AuthorService {
	return &authorServiceImpl{
		authorRepo: authorRepo,
	}
}

func (s *authorServiceImpl) CreateAuthor(r *http.Request) (lastInsertedID int, err error) {
	body := &dto.AuthorCreateRequest{}
	if err := body.Parse(r); err != nil {
		return 0, e.NewError(e.ErrDecodeRequestBody, "Author request parse error", err)
	}

	if err := body.Validate(); err != nil {
		return 0, e.NewError(e.ErrValidateRequest, "Validation error during author creation", err)
	}

	authorID, err := s.authorRepo.CreateAuthor(body)
	if err != nil {
		return 0, e.NewError(e.ErrInvalidRequest, "Author creation error", err)
	}

	return authorID, nil
}

func (s *authorServiceImpl) GetAuthor(r *http.Request) (authorResp *dto.AuthorResponse, err error) {
	body := &dto.AuthorRequest{}
	if err := body.Parse(r); err != nil {
		return nil, e.NewError(e.ErrInvalidRequest, "Author request parse error", err)
	}

	if err := body.Validate(); err != nil {

		return nil, e.NewError(e.ErrValidateRequest, "Validation error", err)
	}

	author, err := s.authorRepo.GetAuthor(body.ID)
	if err != nil {
		return nil, e.NewError(e.ErrInvalidRequest, "No author found with mentioned id", err)
	}

	return author, nil
}

func (s *authorServiceImpl) GetAllAuthors() (authorsResp []*dto.AuthorResponse, err error) {
	result, err := s.authorRepo.GetAllAuthors()
	if err != nil {
		return nil, e.NewError(e.ErrGetAllRequest, "All Authors parsing error", err)
	}
	var authorsList []*dto.AuthorResponse
	for _, value := range result {

		var author dto.AuthorResponse
		author.ID = value.ID
		author.Name = value.Name
		author.Status = value.Status
		author.CreatedBy = value.CreatedBy
		author.CreatedAt = value.CreatedAt
		author.UpdatedBy = value.UpdatedBy
		author.UpdatedAt = value.UpdatedAt
		author.DeletedAt = value.DeletedAt
		author.DeletedBy = value.DeletedBy
		author.IsDeleted = value.IsDeleted

		authorsList = append(authorsList, &author)
	}
	return authorsList, nil
}

func (s *authorServiceImpl) UpdateAuthor(r *http.Request) error {
	body := &dto.AuthorUpdateRequest{}
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "Author update decode error", err)
	}

	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "author update validation error", err)
	}

	if err := s.authorRepo.UpdateAuthor(body); err != nil {
		return e.NewError(e.ErrInternalServer, "Author updation error", err)
	}

	return nil
}

func (s *authorServiceImpl) DeleteAuthor(r *http.Request) error {
	body := &dto.AuthorRequest{}
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrInvalidRequest, "Author delete parse error", err)
	}

	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "Author deletion validation error", err)
	}

	if err := s.authorRepo.DeleteAuthor(body.ID); err != nil {
		return e.NewError(e.ErrInvalidRequest, "Author deletion failed", err)
	}

	return nil
}
