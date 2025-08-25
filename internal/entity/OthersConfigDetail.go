package entity

import (
	"time"
	"gorm.io/gorm"
)

type OthersConfigDetail struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	// Tambahkan field lain sesuai kebutuhan
}