package dto

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type UserResponse struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type UserRequest struct {
	ID int `validate:"required"`
}

// for path param
func (u *UserRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	u.ID = intID
	return nil

}

func (u *UserRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}

// for body param
type UserCreateRequest struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	UserName string `json:"username" validate:"required"`
	Mail     string `json:"mail" validate:"required"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

func (u *UserCreateRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		return err
	}
	return nil
}

func (u *UserCreateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}

type UserUpdateRequest struct {
	ID       int    `gorm:"primarykey" json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (u *UserUpdateRequest) Parse(r *http.Request) error {
	// get ID from request
	strID := chi.URLParam(r, "id")
	// conversion from string to int type
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	u.ID = intID
	// decode to UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		return err
	}
	return nil

}

func (u *UserUpdateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}

type UserDeleteRequest struct {
	ID        int `gorm:"primarykey" json:"id"`
	DeletedBy int `json:"deleted_by"`
}

func (u *UserDeleteRequest) Parse(r *http.Request) error {
	// get ID from request
	strID := chi.URLParam(r, "id")
	// conversion from string to int type
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	u.ID = intID
	// decode to UserDeleteRequest
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		return err
	}
	return nil

}

func (u *UserDeleteRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}
