package repo

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"created_at"`
	UpdateAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy *int64         `gorm:"index"`
	IsDeleted bool           `gorm:"is_deleted"`
}
