package service

import (
	"todolist-api/internal/models"
	"todolist-api/internal/repository"
)

type TodoService struct {
	Repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{Repo: repo}
}

func (s *TodoService) CreateTodo(todo *models.TodoItem) error {
	return s.Repo.Create(todo)
}

func (s *TodoService) GetAllTodos() ([]models.TodoItem, error) {
	return s.Repo.GetAll()
}
