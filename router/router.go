package router

import (
	"basictrade/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	router.POST("auth/register", controllers.AdminRegister)

	return router
}

// 	router := gin.Default()

// 	router.GET("/students", handler.GetAllStudents)
// 	studentRouter := router.Group("/student")
// 	{
// 		studentRouter.POST("/", handler.CreateStudent)
// 		studentRouter.PUT("/:id", handler.UpdateStudent)
// 		studentRouter.DELETE("/:id", handler.DeleteStudent)
// 	}

// 	router.GET("/scores", handler.GetAllScores)
// 	scoreRouter := router.Group("/score")
// 	{
// 		scoreRouter.POST("/", handler.CreateScore)
// 		scoreRouter.PUT("/:id", handler.UpdateScore)
// 		scoreRouter.DELETE("/:id", handler.DeleteScore)
// 	}

// 	return router
// }
