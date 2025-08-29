package entity

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

// Customer model - update untuk menambahkan field baru
// Customer model - update untuk menambahkan field baru
type Customer struct {
	ID               string         `json:"id" gorm:"primaryKey;size:26"`
	Name             string         `json:"name" gorm:"not null"`
	BrandName        string         `json:"brand_name"`
	Code             string         `json:"code" gorm:"unique"`
	AccountManagerID *string        `json:"account_manager_id" gorm:"size:5"`
	Email            string         `json:"email"`
	Phone            string         `json:"phone"`
	Website          string         `json:"website"`
	Description      string         `json:"description"`
	Logo             string         `json:"logo"`
	Status           string         `json:"status" gorm:"default:'Active'"` // Status internal
	Category         string         `json:"category"`
	Rating           float64        `json:"rating" gorm:"default:0"`
	AverageCost      float64        `json:"average_cost" gorm:"default:0"`
	LogoSmall        string         `json:"logo_small"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	AccountManager *AccountManager `json:"account_manager,omitempty" gorm:"foreignKey:AccountManagerID;references:ID"`
	Addresses      []Address       `json:"addresses,omitempty" gorm:"foreignKey:CustomerID"`
	Sosmeds        []Sosmed        `json:"sosmeds,omitempty" gorm:"foreignKey:CustomerID"`
	Contacts       []Contact       `json:"contacts,omitempty" gorm:"foreignKey:CustomerID"`
	Structures     []Structure     `json:"structures,omitempty" gorm:"foreignKey:CustomerID"`
	Groups         []Group         `json:"groups,omitempty" gorm:"many2many:customer_groups;"`
	Others         []Other         `json:"others,omitempty" gorm:"foreignKey:CustomerID"`
	Activities     []Activity      `json:"activities,omitempty" gorm:"foreignKey:CustomerID"`
	Events         []Event         `json:"events,omitempty" gorm:"foreignKey:CustomerID"`
}

// BeforeCreate hook - generate ID before create
func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	id, err := ulid.New(ulid.Timestamp(time.Now()), rand.Reader)
	if err != nil {
		return err
	}
	c.ID = id.String()
	return nil
}
