package requests

import "mime/multipart"

type ProductRequest struct {
	ProductName string                `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	Image       *multipart.FileHeader `form:"file"`
}
