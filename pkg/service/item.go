package service

import (
	"github.com/aspandyar/todo-list/internal/models"
	"github.com/aspandyar/todo-list/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) CreateItem(userId, listId int, item models.TodoItem) (int, error) {
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.CreateItem(listId, item)
}

func (s *TodoItemService) GetAllItem(userId, listId int) ([]models.TodoItem, error) {
	return s.repo.GetAllItem(userId, listId)
}
