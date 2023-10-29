package repository

import (
	"fmt"
	"github.com/aspandyar/todo-list/internal/models"
	"github.com/jmoiron/sqlx"
)

type TodoListRepository struct {
	db *sqlx.DB
}

func NewTodoListRepository(db *sqlx.DB) *TodoListRepository {
	return &TodoListRepository{db: db}
}

func (r *TodoListRepository) Create(userId int, list models.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listId int
	createListsQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListsQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id, lists_id) VALUES ($1, $2) RETURNING id", usersListsTable)
	_, err = tx.Exec(createUsersListsQuery, userId, listId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}

func (r *TodoListRepository) GetAll(userId int) ([]models.TodoList, error) {
	var lists []models.TodoList

	getAllListQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.lists_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)

	err := r.db.Select(&lists, getAllListQuery, userId)

	return lists, err
}
