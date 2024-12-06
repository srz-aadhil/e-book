package repo

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time
	// CreatedBy *int64
	// UpdateAt  time.Time
	UpdateBy  int
	DeletedAt gorm.DeletedAt
	DeletedBy *int64
}
