package controller

import (
	"ebookmod/pkg/api"
	"ebookmod/pkg/e"
	"ebookmod/service"
	"net/http"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{
		userService: userService,
	}
}

func (c *userControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := c.userService.CreateUser(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "User creation failed")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, userID)
}

func (c *userControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := c.userService.GetUser(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "User fetching failed")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, user)
}

func (c *userControllerImpl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers, err := c.userService.GetAllUsers(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "Fetching all users failed")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, allUsers)
}

func (c *userControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if err := c.userService.UpdateUser(r); err != nil {
		httpErr := e.NewAPIError(err, "User updating failed")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, "User Updation successfull")
}

func (c *userControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if err := c.userService.DeleteUser(r); err != nil {
		httpErr := e.NewAPIError(err, "No user found with mentioned id")
		api.Fail(w, httpErr.Statuscode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	api.Success(w, http.StatusOK, "User deletion successfull")
}
