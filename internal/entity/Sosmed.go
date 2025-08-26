package entity

import (
	"time"

	"gorm.io/gorm"
)

// Sosmed model - update untuk menambahkan field handle dan active
// Sosmed model
type Sosmed struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	CustomerID uint           `json:"customer_id" gorm:"not null"`
	Name       string         `json:"name" gorm:"not null"`
	Address    string         `json:"Address" gorm:"not null"`
	Active     bool           `json:"active" gorm:"default:true"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations - hilangkan dari JSON response
	Customer Customer `json:"-" gorm:"foreignKey:CustomerID"`
}
