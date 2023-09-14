package controllers

import (
	"basictrade/database"
	"basictrade/helpers"
	models "basictrade/models/entity"
	"basictrade/models/requests"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateVariant(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	adminID := uint(adminData["id"].(float64))
	productUUID := ctx.PostForm("product_id")
	parsedproductUUID, err := uuid.Parse(productUUID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid UUID format",
		})
		return
	}

	var variantReq requests.VariantRequest
	if contentType == appJSON {
		if err := ctx.ShouldBindJSON(&variantReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := ctx.ShouldBind(&variantReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	var product *models.Product
	if err := db.Where("uuid = ?", parsedproductUUID).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Product not found",
			})
			return
		}
	}
	if product.AdminID != adminID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Your account not allowed to create variant",
			"message": "Your adminID is not match with product adminID",
		})
		return
	}
	if product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Errorf("product with uuid %s not found", productUUID),
		})
		return
	}

	variant := models.Variant{
		UUID:      uuid.New().String(),
		ProductID: product.ID,
		Name:      variantReq.Name,
		Quantity:  variantReq.Quantity,
	}

	if err := db.Create(&variant).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": variant,
	})
}

func GetAllVariant(ctx *gin.Context) {
	db := database.GetDB()
	var page helpers.Pagination
	page.Limit, _ = strconv.Atoi(ctx.Query("limit"))
	page.Page, _ = strconv.Atoi(ctx.Query("page"))
	page.Sort = ctx.Query("sort")
	page.Search = ctx.Query("search")
	var variants []models.Variant

	fmt.Println("search", page.Search)

	db.Scopes(helpers.Paginate(&variants, &page, db)).Find(&variants)

	page.Rows = variants
	page.Column = nil

	ctx.JSON(http.StatusOK, gin.H{
		"data": page,
	})
}

func GetVariant(ctx *gin.Context) {
	db := database.GetDB()
	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	adminID := uint(adminData["id"].(float64))
	variantUUID := ctx.Param("variantUUID")
	parsedUUID, err := uuid.Parse(variantUUID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid UUID format",
		})
		return
	}

	var variant models.Variant

	if err := db.Preload("Product").Where("uuid = ?", parsedUUID).First(&variant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Variant not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	if variant.Product.AdminID != adminID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to access this data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": variant,
	})
}

func UpdateVariant(ctx *gin.Context) {
	db := database.GetDB()
	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	adminID := uint(adminData["id"].(float64))
	variantUUID := ctx.Param("variantUUID")
	parsedUUID, err := uuid.Parse(variantUUID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid UUID format",
		})
		return
	}

	var variantReq requests.VariantRequest
	if err := ctx.ShouldBind(&variantReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var variant models.Variant

	if err := db.Preload("Product").Where("uuid = ?", parsedUUID).First(&variant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Variant not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	if variant.Product.AdminID != adminID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to update this variant",
		})
		return
	}

	variant.Name = variantReq.Name
	variant.Quantity = variantReq.Quantity

	if err := db.Save(&variant).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": variant,
	})
}

func DeleteVariant(ctx *gin.Context) {
	db := database.GetDB()
	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	adminID := uint(adminData["id"].(float64))
	variantUUID := ctx.Param("variantUUID")
	parsedUUID, err := uuid.Parse(variantUUID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid UUID format",
		})
		return
	}

	var variant models.Variant

	if err := db.Preload("Product").Where("uuid = ?", parsedUUID).First(&variant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Variant not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	if variant.Product.AdminID != adminID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to delete this variant",
		})
		return
	}

	if err := db.Delete(&variant).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Variant deleted successfully",
	})
}
