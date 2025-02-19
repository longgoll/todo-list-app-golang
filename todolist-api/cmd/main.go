package main

import (
	"todolist-api/internal/config"
	"todolist-api/internal/handlers"
	"todolist-api/internal/repository"
	"todolist-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.InitDB()

	todoRepo := repository.NewTodoRepository(config.DB)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	v1 := r.Group("/api/v1")

	v1.GET("/todos", todoHandler.GetTodosHandler)
	v1.POST("/todos", todoHandler.CreateTodoHandler)

	r.Run(":8080")
}
