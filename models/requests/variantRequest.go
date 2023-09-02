package requests

type VariantRequest struct {
	VariantName string `json:"variant_name" form:"variant_name" valid:"required~Variant name is required"`
	Quantity    uint   `json:"quantity" valid:"required~Quantity is required,numeric~Invalid quantity format"`
	ProductID   uint   `json:"product_id" valid:"required~Product ID is required"`
}
