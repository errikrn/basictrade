package requests

type VariantRequest struct {
	VariantName string `gorm:"not null" json:"variant_name" form:"variant_name" valid:"required~Variant name is required"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	ProductID   uint
}
