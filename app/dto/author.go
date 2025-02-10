package dto

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type AuthorResponse struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"size:255;not null"`
	CreatedBy int            `json:"created_by,omitempty" gorm:"column:created_by"`
	Status    bool           `json:"status" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	UpdatedBy *int           `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	DeletedBy *int           `json:"deleted_by,omitempty" gorm:"index"`
}

type AuthorRequest struct {
	ID int `validate:"required"`
}

// for Path param
func (a *AuthorRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	// converting string ID to int ID
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	a.ID = intID
	return nil
}

func (a *AuthorRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil
}

// For Body param
type AuthorCreateRequest struct {
	// ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	CreatedBy int    `json:"created_by"`
}

func (a *AuthorCreateRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		return err
	}
	return nil
}
func (a *AuthorCreateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil

}

type AuthorUpdateRequest struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UpdatedBy int    `json:"updated_by"`
}

func (a *AuthorUpdateRequest) Parse(r *http.Request) error {
	//Get ID from request
	strID := chi.URLParam(r, "id")
	IntID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	a.ID = IntID
	//Decode to AuthorUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		return err
	}
	return nil

}

func (a *AuthorUpdateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil

}

type AuthorDeleteRequest struct {
	ID        int  `json:"id"`
	DeletedBy *int `json:"deleted_by"`
}

func (a *AuthorDeleteRequest) Parse(r *http.Request) error {
	//Get ID from request
	strID := chi.URLParam(r, "id")
	IntID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	a.ID = IntID
	//Decode to AuthorUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		return err
	}
	return nil

}

func (a *AuthorDeleteRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil

}
