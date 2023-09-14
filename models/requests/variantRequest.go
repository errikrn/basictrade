package requests

type VariantRequest struct {
	ProductID string `json:"product_id" form:"product_id"`
	Name      string `json:"variant_name" form:"variant_name" binding:"required"`
	Quantity  int    `json:"quantity" form:"quantity" binding:"required"`
}
