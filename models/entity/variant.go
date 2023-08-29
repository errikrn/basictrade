package entity

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Variant struct {
	ID          uint   `gorm:"primaryKey"`
	UUID        string `json:uuid`
	VariantName string `gorm:"not null" json:"variant_name" form:"variant_name" valid:"required~Variant name is required"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	ProductID   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Variant before create()")

	if len(v.VariantName) < 5 {
		err = errors.New("Variant name is too short")
	}

	return
}
