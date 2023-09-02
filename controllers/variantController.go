package controllers

import (
	"basictrade/database"
	models "basictrade/models/entity"
	"basictrade/models/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

func CreateVariant(ctx *gin.Context) {
	db := database.GetDB()

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	productID := uint(adminData["id"].(float64))

	// productUUID := ctx.Param("productUUID")

	// parsedUUID, err := uuid.Parse(productUUID)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error":   "Bad request",
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	var variantReq requests.VariantRequest
	if err := ctx.ShouldBind(&variantReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	variant := models.Variant{
		VariantName: variantReq.VariantName,
		Quantity:    variantReq.Quantity,
		ProductID:   productID,
	}

	if err := db.Debug().Create(&variant).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": variant,
	})
}

func GetAllVariant(ctx *gin.Context) {
	db := database.GetDB()

	var variants []models.Variant

	if err := db.Preload("Admin").Find(&variants).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": variants,
	})
}
