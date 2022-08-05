package routes

import (
	"github.com/Ismadrade/api-book/controllers"
	"github.com/Ismadrade/api-book/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("users")
		{
			user.POST("/", controllers.CreateUser)
		}
		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/", controllers.ShowAllBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
		main.POST("login", controllers.Login)
	}
	return router
}
