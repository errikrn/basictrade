package router

import (
	"basictrade/controllers"
	"basictrade/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	userRouter := router.Group("/auth")
	{
		userRouter.POST("/register", controllers.AdminRegister)
		userRouter.POST("/login", controllers.AdminLogin)

	}
	productRouter := router.Group("/products")
	{
		productRouter.GET("", controllers.GetAllProduct)
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/:productUUID", middlewares.ProductAuthorization(), controllers.GetProduct)
		productRouter.PUT("/:productUUID", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productUUID", middlewares.ProductAuthorization(), controllers.DeleteProduct)

	}

	variantRouter := router.Group("/products/variants")
	{
		variantRouter.GET("", controllers.GetAllVariant)
		variantRouter.Use(middlewares.Authentication())
		variantRouter.POST("/", controllers.CreateVariant)
		variantRouter.GET("/:variantUUID", middlewares.VariantAuthorization(), controllers.GetVariant)
		variantRouter.PUT("/:variantUUID", middlewares.VariantAuthorization(), controllers.UpdateVariant)
		variantRouter.DELETE("/:variantUUID", middlewares.VariantAuthorization(), controllers.DeleteVariant)
	}

	return router
}
