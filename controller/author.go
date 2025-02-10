package controller

import (
	"ebookmod/pkg/api"
	"ebookmod/pkg/e"
	"ebookmod/service"
	"net/http"
)

type AuthorController interface {
	CreateAuthor(w http.ResponseWriter, r *http.Request)
	GetAuthor(w http.ResponseWriter, r *http.Request)
	GetAllAuthors(w http.ResponseWriter, r *http.Request)
	UpdateAuthor(w http.ResponseWriter, r *http.Request)
	DeleteAuthor(w http.ResponseWriter, r *http.Request)
}

type authorControllerImpl struct {
	authorService service.AuthorService
}

func NewAuthorController(authorService service.AuthorService) AuthorController {
	return &authorControllerImpl{
		authorService: authorService,
	}
}

func (c *authorControllerImpl) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	authorID, err := c.authorService.CreateAuthor(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "Author creation failed")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, authorID)

}

func (c *authorControllerImpl) GetAuthor(w http.ResponseWriter, r *http.Request) {
	author, err := c.authorService.GetAuthor(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "Author fetching failed")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, author)
}

func (c *authorControllerImpl) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	allAuthors, err := c.authorService.GetAllAuthors()
	if err != nil {
		httpErr := e.NewAPIError(err, "Fetching all authors failed")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, allAuthors)
}

func (c *authorControllerImpl) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	if err := c.authorService.UpdateAuthor(r); err != nil {
		httpErr := e.NewAPIError(err, "Updating author failed")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, "Author Upation Success")
}

func (c *authorControllerImpl) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	if err := c.authorService.DeleteAuthor(r); err != nil {
		httpErr := e.NewAPIError(err, "Deleting author failed")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, "Author Deletion successfully completed")
}
