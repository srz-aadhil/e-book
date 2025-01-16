package controller

import (
	"ebookmod/pkg/api"
	"ebookmod/pkg/e"
	"ebookmod/service"
	"net/http"
)

type BookController interface {
	CreateBook(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
}

type bookControllerImpl struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &bookControllerImpl{
		bookService: bookService,
	}
}

func (c *bookControllerImpl) CreateBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := c.bookService.CreateBook(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "Book creation failed")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, bookID)
}

func (c *bookControllerImpl) GetBook(w http.ResponseWriter, r *http.Request) {
	book, err := c.bookService.GetBook(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "Book fetching failed")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, book)
}

func (c *bookControllerImpl) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	allBooks, err := c.bookService.GetAllBooks()
	if err != nil {
		httpErr := e.NewAPIError(err, "Fetching all books failed")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, allBooks)
}

func (c *bookControllerImpl) UpdateBook(w http.ResponseWriter, r *http.Request) {
	if err := c.bookService.UpdateBook(r); err != nil {
		httpErr := e.NewAPIError(err, "Book updating failed")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, "Book Updation successfull")
}

func (c *bookControllerImpl) DeleteBook(w http.ResponseWriter, r *http.Request) {
	if err := c.bookService.DeleteBook(r); err != nil {
		httpErr := e.NewAPIError(err, "No book found with mentioned id")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, "Book deletion successfull")
}
