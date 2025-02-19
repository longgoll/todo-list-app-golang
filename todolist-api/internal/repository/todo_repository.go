package repository

import (
	"todolist-api/internal/models"

	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) Create(todo *models.TodoItem) error {
	return r.DB.Create(todo).Error
}

func (r *TodoRepository) GetAll() ([]models.TodoItem, error) {
	var todos []models.TodoItem
	err := r.DB.Find(&todos).Error
	return todos, err
}
