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
	v1.GET("/todos/:id", todoHandler.GetTodoByIDHandler)
	v1.PUT("/todos/:id", todoHandler.UpdateTodoByIDHandler)
	v1.DELETE("/todos/:id", todoHandler.DeleteTodoByIDHandler)

	r.Run(":8080")
}
