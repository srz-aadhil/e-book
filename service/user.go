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

type UserService interface {
	CreateUser(r *http.Request) (lastInsertedID int, err error)
	GetUser(r *http.Request) (userResp *dto.UserResponse, err error)
	GetAllUsers(r *http.Request) (allUsers []*dto.UserResponse, err error)
	UpdateUser(r *http.Request) error
	DeleteUser(r *http.Request) error
}

type userServiceImpl struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
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

	log.Info().Msg("Parsing and Validation successfully completed")
	userID, err := s.userRepo.CreateUser(body)
	if err != nil {
		return 0, e.NewError(e.ErrInternalServer, "User creation failed", err)
	}

	log.Info().Msg("User creation successfully completed")
	return userID, nil
}

func (s *userServiceImpl) GetUser(r *http.Request) (*dto.UserResponse, error) {
	body := &dto.UserRequest{}
	if err := body.Parse(r); err != nil {
		return nil, e.NewError(e.ErrInvalidRequest, "User fetching parse error", err)
	}

	if err := body.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "User fetching validation error", err)
	}

	log.Info().Msg("Parsing and validaton successfully completed")

	user, err := s.userRepo.GetUser(body.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Msg("Record not found with mentioned id")
			return nil, e.NewError(e.ErrResourceNotFound, "No user with mentioned id", err)
		}
		return nil, e.NewError(e.ErrInternalServer, "User fetching failed", err)
	}

	//storing data from user struct to dto.UserResponse struct
	var userResp dto.UserResponse
	userResp.ID = user.ID
	userResp.UserName = user.Username
	userResp.CreatedAt = user.CreatedAt
	userResp.UpdatedAt = user.UpdateAt

	log.Info().Msg("User retrieved successfully")
	return &userResp, nil
}

func (s *userServiceImpl) GetAllUsers(r *http.Request) (allUsers []*dto.UserResponse, err error) {
	result, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, e.NewError(e.ErrGetAllRequest, "All users fetching parse error", err)
	}

	var usersList []*dto.UserResponse

	for _, val := range result {
		var user dto.UserResponse

		user.ID = val.ID
		user.UserName = val.Username
		user.CreatedAt = val.CreatedAt
		user.UpdatedAt = val.UpdateAt
		user.IsDeleted = val.IsDeleted

		usersList = append(usersList, &user)
	}

	log.Info().Msg("All users retrieved successfully")
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

	log.Info().Msg("User parsing and validation successfully completed")

	if err := s.userRepo.UpdateUser(body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Msg("Record not found")
			return e.NewError(e.ErrResourceNotFound, "No user found mentioned id ", err)
		}
		return e.NewError(e.ErrInternalServer, "User updation failed", err)
	}

	log.Info().Msg("User successfully updated")
	return nil
}

func (s *userServiceImpl) DeleteUser(r *http.Request) error {
	body := &dto.UserDeleteRequest{}
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "User deletion parse error", err)
	}

	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "User deletion validate error", err)
	}

	log.Info().Msg("Request parsing and validation successfully completed")

	if err := s.userRepo.DeleteUser(body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Msg("Record not found")
			return e.NewError(e.ErrResourceNotFound, "No user with mentioned id to delete", err)
		}
		return e.NewError(e.ErrResourceNotFound, "No user found with mentioned id", err)
	}

	log.Info().Msg("User deletion successfully completed")
	return nil
}
