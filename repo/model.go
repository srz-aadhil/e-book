package repo

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time `gorm:"created_at"`
	//CreatedBy *int64         `gorm:"column:created_by"`
	UpdateAt time.Time `gorm:"column:updated_at"`
	//UpdateBy  int            `gorm:"column:updated_by"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy *int64         `gorm:"index"`
	IsDeleted bool           `gorm:"is_deleted"`
}
