package entity

import (
	"time"

	"gorm.io/gorm"
)

// Address model - update untuk menambahkan field name dan active
// Address model
type Address struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	CustomerID uint `json:"customer_id" gorm:"not null"`
	// SupplierID *uint          `json:"supplier_id"` // HAPUS field ini
	Name string `json:"name" gorm:"not null"`

	Address string `json:"address" gorm:"not null"`

	Main bool `json:"main" gorm:"default:false"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations - hilangkan dari JSON response
	Customer Customer `json:"-" gorm:"foreignKey:CustomerID"`
}
