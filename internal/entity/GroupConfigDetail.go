package entity

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type GroupConfigDetail struct {
	ID            string         `json:"id" gorm:"primaryKey;size:26"`
	GroupConfigID string         `json:"-" gorm:"not null"` // ULID string, hidden from JSON
	Name          string         `json:"name" gorm:"not null;unique"`
	Icon          string         `json:"icon" gorm:"default:'default.icon';not null"`
	IsActive      bool           `json:"-" gorm:"default:true"` // Hidden from JSON
	CreatedAt     time.Time      `json:"-" gorm:""`             // Hidden from JSON
	UpdatedAt     time.Time      `json:"-" gorm:""`             // Hidden from JSON
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
	// Relations - hilangkan dari JSON response
	GroupConfig GroupConfig `json:"-" gorm:"foreignKey:GroupConfigID"`
}

// before save generate id
func (s *GroupConfigDetail) BeforeCreate(tx *gorm.DB) (err error) {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	s.ID = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	return
}
