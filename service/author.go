package service

import (
	"ebookmod/app/dto"
	"ebookmod/pkg/e"
	"ebookmod/repo"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
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

	log.Info().Msg("Successsfully completed validation and parsing")

	authorID, err := s.authorRepo.CreateAuthor(body)
	if err != nil {
		return 0, e.NewError(e.ErrInvalidRequest, "Author creation error", err)
	}

	log.Info().Msgf("Successfully created author with id-%d", authorID)
	return authorID, nil
}

func (s *authorServiceImpl) GetAuthor(r *http.Request) (*dto.AuthorResponse, error) {
	body := &dto.AuthorRequest{}
	if err := body.Parse(r); err != nil {
		return nil, e.NewError(e.ErrInvalidRequest, "Author request parse error", err)
	}

	if err := body.Validate(); err != nil {

		return nil, e.NewError(e.ErrValidateRequest, "Validation error", err)
	}

	log.Info().Msg("Successsfully completed validation and parsing")

	author, err := s.authorRepo.GetAuthor(body.ID)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Msg("Record Not found")
			return nil, e.NewError(e.ErrResourceNotFound, "Author record not found", err)
		}

		return nil, e.NewError(e.ErrInvalidRequest, "No author found with mentioned id", err)
	}
	var authorResp dto.AuthorResponse
	authorResp.ID = author.ID
	authorResp.Name = author.Name
	authorResp.CreatedBy = author.CreatedBy
	authorResp.CreatedAt = author.CreatedAt
	authorResp.UpdatedBy = author.UpdatedBy
	authorResp.UpdatedAt = author.UpdateAt
	authorResp.DeletedBy = author.DeletedBy
	authorResp.DeletedAt = author.DeletedAt
	authorResp.Status = author.Status

	log.Info().Msgf("Author successfully retrieved with id- %d", authorResp.ID)
	return &authorResp, nil
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
		author.UpdatedAt = value.UpdateAt
		author.DeletedAt = value.DeletedAt
		author.DeletedBy = value.DeletedBy

		authorsList = append(authorsList, &author)
	}

	log.Info().Msg("All authors successsfully retrieved")
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info().Msgf("Record not found")
			return e.NewError(e.ErrResourceNotFound, "No author found to update", err)
		}

		return e.NewError(e.ErrInternalServer, "Author updation error", err)
	}

	log.Info().Msgf("Author updation successfully completed")
	return nil
}

func (s *authorServiceImpl) DeleteAuthor(r *http.Request) error {
	body := &dto.AuthorDeleteRequest{}
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrInvalidRequest, "Author delete parse error", err)
	}

	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "Author deletion validation error", err)
	}

	if err := s.authorRepo.DeleteAuthor(body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info().Msgf("Record not found")
			return e.NewError(e.ErrResourceNotFound, "Author record not found", err)
		}
		return e.NewError(e.ErrInvalidRequest, "Author deletion failed", err)
	}

	log.Info().Msgf("Author with id %d deleted successfully", body.ID)
	return nil
}
