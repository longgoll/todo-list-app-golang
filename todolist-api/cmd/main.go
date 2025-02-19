package main

import (
	"todolist-api/internal/config"
	"todolist-api/internal/handlers"
	"todolist-api/internal/repository"
	"todolist-api/internal/service"

	_ "todolist-api/docs" // Import the generated docs package

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Todo List API
// @version 1.0
// @description This is a simple Todo List application built with Golang and Gin framework.
// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := gin.Default()

	config.InitDB()

	todoRepo := repository.NewTodoRepository(config.DB)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	v1 := r.Group("/api/v1")

	v1.GET("/todos", todoHandler.GetTodosHandler)
	v1.POST("/todos", todoHandler.CreateTodoHandler)
	v1.GET("/todos/:id", todoHandler.GetTodoByIDHandler)
	v1.PUT("/todos/:id", todoHandler.UpdateTodoByIDHandler)
	v1.DELETE("/todos/:id", todoHandler.DeleteTodoByIDHandler)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
