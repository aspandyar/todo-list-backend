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

func (s *TodoListService) CreateList(userId int, list models.TodoList) (int, error) {
	return s.repo.CreateList(userId, list)
}

func (s *TodoListService) GetAllLists(userId int) ([]models.TodoList, error) {
	return s.repo.GetAllLists(userId)
}

func (s *TodoListService) GetListById(userId, listId int) (models.TodoList, error) {
	return s.repo.GetListById(userId, listId)
}

func (s *TodoListService) UpdateList(userId, listId int, updateListInput models.UpdateListInput) error {
	if err := updateListInput.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateList(userId, listId, updateListInput)
}

func (s *TodoListService) DeleteList(userId, listId int) error {
	return s.repo.DeleteList(userId, listId)
}
