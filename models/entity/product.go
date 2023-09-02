package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint      `gorm:"primaryKey;AutoIncrement" json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `gorm:"not null" json:"name" form:"name" valid:"required~Variant name is required"`
	ImageURL  string    `gorm:"not null" json:"file" form:"file"`
	AdminID   uint      `json:"admin_id"`
	Admin     *Admin    `json:"admin,omitempty" gorm:"association_autoupdate:false"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
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
