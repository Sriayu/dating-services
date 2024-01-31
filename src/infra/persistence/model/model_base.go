package model

import (
	"time"

	"gorm.io/gorm"
)

// Base .
type Base struct {
	ID        uint           `gorm:"column:id" json:"id"`
	CreatedAt time.Time      `gorm:"index:,sort:desc" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"index,sort:desc" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
