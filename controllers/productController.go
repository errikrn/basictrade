package controllers

import (
	"basictrade/database"
	"basictrade/helpers"
	models "basictrade/models/entity"
	"basictrade/models/requests"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	adminID := uint(adminData["id"].(float64))

	newUUID := uuid.New()
	productUUID := newUUID.String()

	var productReq requests.ProductRequest
	if contentType == appJSON {
		if err := ctx.ShouldBindJSON(&productReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := ctx.ShouldBind(&productReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	fileName := helpers.RemoveExtension(productReq.Image.Filename)

	uploadResult, err := helpers.UploadFile(productReq.Image, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		UUID:     productUUID,
		Name:     productReq.Name,
		ImageURL: uploadResult,
		AdminID:  adminID,
	}

	if err := db.Debug().Create(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	admin := models.Admin{ID: adminID}
	if err := db.First(&admin).Omit("products").Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "cannot get admin data: " + err.Error(),
		})
		return
	}
	product.AdminID = adminID
	product.Admin = &admin
	ctx.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func GetAllProduct(ctx *gin.Context) {
	db := database.GetDB()

	var products []models.Product

	if err := db.Preload("Admin").Find(&products).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()

	fmt.Printf("Entering UpdateProduct\n")

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	adminID := uint(adminData["id"].(float64))
	fmt.Printf("AdminID: %d\n", adminID)

	productUUID := ctx.Param("productUUID")
	fmt.Printf("Received productUUID: %s\n", productUUID)

	parsedUUID, err := uuid.Parse(productUUID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	var productReq requests.ProductRequest
	if err := ctx.ShouldBind(&productReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := helpers.RemoveExtension(productReq.Image.Filename)
	fmt.Printf("Received filename: %s\n", fileName)

	uploadResult, err := helpers.UploadFile(productReq.Image, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := db.Debug().Where("uuid = ? AND admin_id = ?", parsedUUID, adminID).
		First(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "error when get product: " + err.Error(),
		})
		return
	}
	product.Name = productReq.Name
	product.ImageURL = uploadResult
	if err := db.Debug().Updates(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "error when updating product: " + err.Error(),
		})
		return
	}
	// kode ini tidak keluar
	fmt.Printf("Exiting UpdateProduct\n")

	ctx.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func DeleteProduct(ctx *gin.Context) {
	db := database.GetDB()
	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	adminID := uint(adminData["id"].(float64))

	productUUID := ctx.Param("productUUID")

	parsedUUID, err := uuid.Parse(productUUID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	product := models.Product{}
	if err := db.Where("uuid = ? AND admin_id = ?", parsedUUID, adminID).First(&product).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Product not found",
			"message": err.Error(),
		})
		return
	}

	if err := db.Delete(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
