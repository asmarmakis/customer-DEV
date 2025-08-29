package entity

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type AccountManager struct {
	ID          string         `gorm:"type:varchar(5);primaryKey" json:"id"`
	ManagerName string         `gorm:"type:varchar(255);not null" json:"manager_name"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Customers []Customer `json:"customers,omitempty" gorm:"foreignKey:AccountManagerID"`
}

// BeforeCreate hook untuk generate UUID dan ID
func (s *AccountManager) BeforeCreate(tx *gorm.DB) (err error) {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	u := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	// only take last 5 chars
	if len(u) >= 5 {
		s.ID = u[len(u)-5:]
	} else {
		s.ID = u
	}

	return
}

// TableName untuk menentukan nama tabel
func (AccountManager) TableName() string {
	return "account_managers"
}
