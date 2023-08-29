package middlewares

import (
	"basictrade/database"
	"net/http"
	"strconv"

	models "basictrade/models/entity"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(ctx.Param("productId"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid email",
			})
			return
		}

		adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
		adminID := uint(adminData["id"].(float64))
		Product := models.Product{}

		err = db.Select("user_id").First(&Product, uint(productId)).Error
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
