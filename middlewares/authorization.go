package middlewares

import (
	"basictrade/database"
	"net/http"

	models "basictrade/models/entity"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		productUUID := ctx.Param("productUUID")

		parsedUUID, err := uuid.Parse(productUUID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": err.Error(),
			})
			return
		}

		adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
		adminID := uint(adminData["id"].(float64))
		Product := models.Product{}

		err = db.Select("admin_id").Where("uuid = ?", parsedUUID).First(&Product).Error
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Data Not Found",
			})
			return
		}

		if Product.AdminID != adminID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}
