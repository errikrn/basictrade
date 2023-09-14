package entity

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Variant struct {
	ID        uint      `gorm:"primaryKey;AutoIncrement" json:"id"`
	UUID      string    `json:"UUID"`
	Name      string    `gorm:"not null" json:"variant_name" form:"variant_name" valid:"required~Variant name is required"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	ProductID uint      `json:"productID"`
	Product   *Product  `json:"product" gorm:"foreignkey:ProductID;association_foreignkey:ID"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Variant before create()")

	if len(v.Name) < 2 {
		err = errors.New("Variant name is too short")
	}

	return
}
