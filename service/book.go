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

type BookService interface {
	CreateBook(bookReq *http.Request) (lastInsertedID int, err error)
	GetBook(r *http.Request) (bookResp *dto.BookResponse, err error)
	GetAllBooks() (allBooks []*dto.BookResponse, err error)
	UpdateBook(r *http.Request) error
	DeleteBook(r *http.Request) error
}

type bookServiceImpl struct {
	bookRepo repo.BookRepo
}

func NewBookService(bookRepo repo.BookRepo) BookService {
	return &bookServiceImpl{
		bookRepo: bookRepo,
	}
}

func (s *bookServiceImpl) CreateBook(r *http.Request) (lastInsertedID int, err error) {
	body := &dto.BookCreateRequest{}
	if err := body.Parse(r); err != nil {
		return 0, e.NewError(e.ErrDecodeRequestBody, "Book creation parse error", err)
	}

	if err := body.Validate(); err != nil {
		return 0, e.NewError(e.ErrValidateRequest, "Book creation validate error", err)
	}

	log.Info().Msg("Book validation and parsing successfully completed")

	bookID, err := s.bookRepo.CreateBook(body)
	if err != nil {
		return 0, e.NewError(e.ErrInternalServer, "Book creation failed due to internal error", err)
	}

	log.Info().Msgf("Book created with id %d successfully", bookID)
	return bookID, nil
}

func (s *bookServiceImpl) GetBook(r *http.Request) (*dto.BookResponse, error) {
	body := &dto.BookRequest{}
	if err := body.Parse(r); err != nil {
		return nil, e.NewError(e.ErrDecodeRequestBody, "Book fetching parse error", err)
	}

	if err := body.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "Book fetching validate error", err)
	}

	bookResp, err := s.bookRepo.GetBook(body.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Msg("Record not found")
			return nil, e.NewError(e.ErrResourceNotFound, "Book record not found", err)
		}
		return nil, e.NewError(e.ErrResourceNotFound, "Book not found with mentioned id", err)
	}

	//Storing values to return type struct
	var book dto.BookResponse
	book.ID = bookResp.ID
	book.Title = bookResp.Title
	book.Content = bookResp.Content
	book.AuthorID = bookResp.AuthorID
	book.CreatedBy = bookResp.CreatedBy
	book.CreatedAt = bookResp.CreatedAt
	book.UpdatedBy = bookResp.UpdatedBy
	book.UpdatedAt = bookResp.UpdateAt
	book.Status = bookResp.Status
	book.DeletedBy = bookResp.DeletedBy
	book.DeletedAt = bookResp.DeletedAt

	log.Info().Msgf("Book retrieved successfully")
	return &book, nil
}

func (s *bookServiceImpl) GetAllBooks() (allBooks []*dto.BookResponse, err error) {
	result, err := s.bookRepo.GetAllBooks()
	if err != nil {
		return nil, e.NewError(e.ErrGetAllRequest, "All books fetching parse error", err)
	}

	for _, val := range result {

		var book dto.BookResponse
		book.ID = val.ID
		book.Title = val.Title
		book.Content = val.Content
		book.AuthorID = val.AuthorID
		book.CreatedBy = val.CreatedBy
		book.CreatedAt = val.CreatedAt
		book.UpdatedBy = val.UpdatedBy
		book.UpdatedAt = val.UpdateAt
		book.DeletedBy = val.DeletedBy
		book.DeletedAt = val.DeletedAt
		book.Status = val.Status

		allBooks = append(allBooks, &book)
	}

	log.Info().Msg("All books retrieved successfully")
	return allBooks, nil
}

func (s *bookServiceImpl) UpdateBook(r *http.Request) error {
	body := &dto.BookUpdateRequest{}
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "Book updation parse error", err)
	}

	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "Book updation validate error", err)
	}

	if err := s.bookRepo.UpdateBook(body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Msg("Record not found")
			return e.NewError(e.ErrResourceNotFound, "Book not found with mentioned id", err)
		}
		return e.NewError(e.ErrResourceNotFound, "No book with mentioned details", err)
	}

	log.Info().Msg("Book updation successfully completed")
	return nil
}

func (s *bookServiceImpl) DeleteBook(r *http.Request) error {
	body := &dto.BookRequest{}
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "Book deletion parse error", err)
	}

	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "Book deletion validate error", err)
	}

	log.Info().Msg("Parsing and validation successfully completed")

	if err := s.bookRepo.DeleteBook(body.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Msg("Record not found with mention id")
			return e.NewError(e.ErrResourceNotFound, "No book with mention id", err)
		}
		return e.NewError(e.ErrResourceNotFound, "Book not found with mentioned id", err)
	}

	log.Info().Msg("Book Deletion successfully completed")
	return nil
}
