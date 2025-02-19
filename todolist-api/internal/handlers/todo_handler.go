package handlers

import (
	"net/http"

	"todolist-api/internal/models"
	"todolist-api/internal/service"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents the error response
type ErrorResponse struct {
	Error string `json:"error"`
}

type TodoHandler struct {
	Service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{Service: service}
}

// CreateTodoHandler godoc
// @Summary Create a new todo
// @Description Create a new todo
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.TodoItem true "Todo Item"
// @Success 201 {object} models.TodoItem
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /todos [post]
func (h *TodoHandler) CreateTodoHandler(c *gin.Context) {
	var todo models.TodoItem
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// GetTodosHandler godoc
// @Summary Get all todos
// @Description Get all todos
// @Tags todos
// @Produce json
// @Success 200 {array} models.TodoItem
// @Failure 500 {object} ErrorResponse
// @Router /todos [get]
func (h *TodoHandler) GetTodosHandler(c *gin.Context) {
	todos, err := h.Service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// GetTodoByIDHandler godoc
// @Summary Get a todo by ID
// @Description Get a todo by ID
// @Tags todos
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} models.TodoItem
// @Failure 404 {object} ErrorResponse
// @Router /todos/{id} [get]
func (h *TodoHandler) GetTodoByIDHandler(c *gin.Context) {
	id := c.Param("id")
	todo, err := h.Service.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// UpdateTodoByIDHandler godoc
// @Summary Update a todo by ID
// @Description Update a todo by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body models.TodoItem true "Todo Item"
// @Success 200 {object} models.TodoItem
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /todos/{id} [put]
func (h *TodoHandler) UpdateTodoByIDHandler(c *gin.Context) {
	id := c.Param("id")
	var todo models.TodoItem
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.UpdateTodoByID(id, &todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodoByIDHandler godoc
// @Summary Delete a todo by ID
// @Description Delete a todo by ID
// @Tags todos
// @Produce json
// @Param id path string true "Todo ID"
// @Success 204
// @Failure 500 {object} ErrorResponse
// @Router /todos/{id} [delete]
func (h *TodoHandler) DeleteTodoByIDHandler(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.DeleteTodoByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
