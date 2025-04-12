package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gorm/config"
	"github.com/gin-gorm/handlers"
	"github.com/gin-gorm/models"
)

func main() {
	// Initialize database
	config.ConnectDB()
	
	// Auto migrate the schema
	config.DB.AutoMigrate(&models.User{})

	// Initialize router
	r := gin.Default()

	// User routes
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", handlers.CreateUser)
		userRoutes.GET("/", handlers.GetUsers)
		userRoutes.GET("/:id", handlers.GetUser)
		userRoutes.PUT("/:id", handlers.UpdateUser)
		userRoutes.DELETE("/:id", handlers.DeleteUser)
	}

	// Start server
	r.Run(":8080")
} 