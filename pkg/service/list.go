package service

import (
	"github.com/aspandyar/todo-list/internal/models"
	"github.com/aspandyar/todo-list/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list models.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]models.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetListById(userId, listId int) (models.TodoList, error) {
	return s.repo.GetListById(userId, listId)
}