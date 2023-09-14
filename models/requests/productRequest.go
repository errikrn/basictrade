package requests

import "mime/multipart"

type ProductRequest struct {
	Name  string                `form:"name" valid:"required~Product name is required"`
	Image *multipart.FileHeader `form:"file" valid:"file,optional"`
}
