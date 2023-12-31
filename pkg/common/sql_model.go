package common

import (
	"time"

	"gorm.io/gorm"
)

type SQLModel struct {
	Id        int            `json:"id" gorm:"column:id"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}
