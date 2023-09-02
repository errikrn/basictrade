package router

import (
	"basictrade/controllers"
	"basictrade/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/auth")
	{
		userRouter.POST("/register", controllers.AdminRegister)
		userRouter.POST("/login", controllers.AdminLogin)

	}

	productRouter := router.Group("/products")
	{
		productRouter.GET("/", controllers.GetAllProduct)

		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productUUID", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productUUID", middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	variantRouter := router.Group("/products/variants")
	{
		variantRouter.GET("/", controllers.GetAllVariant)

		variantRouter.Use(middlewares.Authentication())
		variantRouter.POST("/", controllers.CreateVariant)
		// variantRouter.PUT("/:variantUUID", middlewares.ProductAuthorization(), controllers.UpdateVariant)
		// variantRouter.DELETE("/:variantUUID", middlewares.ProductAuthorization(), controllers.DeleteVariant)
	}

	return router
}
