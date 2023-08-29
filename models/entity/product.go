package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UUID        string `json:uuid`
	ProductName string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	ImageURL    string `gorm:"not null" json:"file" form:"file"`
	AdminID     uint   `gorm:"not null" json:"stock" form:"stock" valid:"required~Stock of book is required, numeric~Invalid stock format"`
	Admin       *Admin
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
