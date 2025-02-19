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

func (r *TodoRepository) GetByID(id string) (models.TodoItem, error) {
	var todo models.TodoItem
	err := r.DB.First(&todo, id).Error
	return todo, err
}

func (r *TodoRepository) UpdateByID(id string, todo *models.TodoItem) error {
	return r.DB.Model(&models.TodoItem{}).Where("id = ?", id).Updates(todo).Error
}

func (r *TodoRepository) DeleteByID(id string) error {
	return r.DB.Where("id = ?", id).Delete(&models.TodoItem{}).Error
}
