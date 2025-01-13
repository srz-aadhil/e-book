package dto

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BookResponse struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string         `json:"title" gorm:"size:255;not null"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	AuthorID  int            `json:"author_id" gorm:"column:author_id;not null"`
	CreatedBy int            `json:"created_by" gorm:"column:created_by;not null"`
	UpdatedBy *int           `json:"updated_by,omitempty" gorm:"column:updated_by"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	DeletedBy *int           `json:"deleted_by,omitempty" gorm:"index"`
	Status    int            `json:"status" gorm:"column:status;default:1"`
}

// for path param
type BookRequest struct {
	ID int `validation:"required"`
}

func (b *BookRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	b.ID = intID
	return nil
}

func (b *BookRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}

type BookDeleteRequest struct {
	ID        int `json:"id"`
	DeletedBy int `json:"deleted_by"`
}

func (b *BookDeleteRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	b.ID = intID
	if err := json.NewDecoder(r.Body).Decode(b); err != nil {
		return err
	}
	return nil
}

func (b *BookDeleteRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}

// for body param
type BookCreateRequest struct {
	ID        int    `gorm:"primary key" json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	AuthorID  int    `json:"author_id" validate:"required"` //Author.ID
	Status    int    `json:"status"`
	CreatedBy int    `json:"created_by" validate:"required"` // User.ID
}

func (b *BookCreateRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(b); err != nil {
		return err
	}
	return nil
}

func (b *BookCreateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}

type BookUpdateRequest struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    int    `validate:"required"`
	UpdatedBy int    `json:"updated_by" validate:"required"`
}

func (b *BookUpdateRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	b.ID = intID
	if err := json.NewDecoder(r.Body).Decode(b); err != nil {
		return err
	}
	return nil
}

func (b *BookUpdateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}
