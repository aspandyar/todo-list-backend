package repository

import (
	"github.com/aspandyar/todo-list/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type TodoList interface {
	CreateList(userId int, list models.TodoList) (int, error)
	GetAllLists(userId int) ([]models.TodoList, error)
	GetListById(userId, listId int) (models.TodoList, error)
	UpdateList(userId, listId int, updateListInput models.UpdateListInput) error
	DeleteList(userId, listId int) error
}

type TodoItem interface {
	CreateItem(listId int, item models.TodoItem) (int, error)
	GetAllItem(userId, listId int) ([]models.TodoItem, error)
	GetItemById(userId, itemId int) (models.TodoItem, error)
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		TodoList:      NewTodoListRepository(db),
		TodoItem:      NewTodoItemRepository(db),
	}
}
