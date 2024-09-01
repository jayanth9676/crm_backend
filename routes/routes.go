package routes

import (
	"github.com/gin-gonic/gin"
	"your_project/controllers"
	"your_project/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/login", controllers.Login) // TODO: Implement login controller

	// Protected routes
	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		// User routes
		protected.POST("/users", controllers.CreateUser)
		protected.GET("/users/:id", controllers.GetUser)
		protected.PUT("/users/:id", controllers.UpdateUser)
		protected.DELETE("/users/:id", controllers.DeleteUser)

		// Customer routes
		protected.POST("/customers", controllers.CreateCustomer)
		protected.GET("/customers/:id", controllers.GetCustomer)
		protected.PUT("/customers/:id", controllers.UpdateCustomer)
		protected.DELETE("/customers/:id", controllers.DeleteCustomer)
	}

	return router
}