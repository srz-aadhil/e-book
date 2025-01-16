package service

import (
	"ebookmod/app/dto"
	"ebookmod/pkg/e"
	"ebookmod/repo"
	"net/http"
)

type UserService interface {
	CreateUser(r *http.Request) (lastInsertedID int, err error)
	GetUser(r *http.Request) (userResp *dto.UserResponse, err error)
	GetAllUsers() (allUsers []*dto.UserResponse, err error)
	UpdateUser(r *http.Request) error
	DeleteUser(r *http.Request) error
}

type userServiceImpl struct {
	userRepo repo.UserRepo
}

func NewUserController(userRepo repo.User) UserService {
	return &userServiceImpl{
		userRepo: &userRepo,
	}
}

func (s *userServiceImpl) CreateUser(r *http.Request) (lastInsertedID int, err error) {
	body := &dto.UserCreateRequest{}
	if err := body.Parse(r); err != nil {
		return 0, e.NewError(e.ErrDecodeRequestBody, "User creation parse error", err)
	}

	if err := body.Validate(); err != nil {
		return 0, e.NewError(e.ErrValidateRequest, "User creation validate error", err)
	}

	userID, err := s.userRepo.CreateUser(body)
	if err != nil {
		return 0, e.NewError(e.ErrInternalServer, "User creation failed", err)
	}

	return userID, nil
}

func (s *userServiceImpl) GetUser(r *http.Request) (userResp *dto.UserResponse, err error) {
	body := &dto.UserRequest{}
	if err := body.Parse(r); err != nil {
		return nil, e.NewError(e.ErrInvalidRequest, "User fetching parse error", err)
	}

	if err := body.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "User fetching validation error", err)
	}

	userResp, err = s.userRepo.GetUser(body.ID)
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "User fetching failed", err)
	}

	return userResp, nil
}

func (s *userServiceImpl) GetAllUsers(r *http.Request) (userResp []*dto.UserResponse, err error) {
	result, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, e.NewError(e.ErrGetAllRequest, "All users fetching parse error", err)
	}

	var usersList []*dto.UserResponse

	for _, val := range result {
		var user dto.UserResponse

		user.ID = val.ID
		user.UserName = val.UserName
		user.IsDeleted = val.IsDeleted

		usersList = append(usersList, &user)
	}

	return usersList, nil
}

func (s *userServiceImpl) UpdateUser(r *http.Request) error {
	body := &dto.UserUpdateRequest{}
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "User updation parse error", err)
	}

	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "User updation validate error", err)
	}

	if err := s.userRepo.UpdateUser(body); err != nil {
		return e.NewError(e.ErrInternalServer, "User updation failed", err)
	}
	return nil
}

func (s *userServiceImpl) DeleteUser(r *http.Request) error {
	body := &dto.UserRequest{}
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "User deletion parse error", err)
	}

	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "User deletion validate error", err)
	}

	if err := s.userRepo.DeleteUser(body.ID); err != nil {
		return e.NewError(e.ErrResourceNotFound, "No user found with mentioned id", err)
	}

	return nil
}
