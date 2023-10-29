package repository

import (
	"fmt"
	"github.com/aspandyar/todo-list/internal/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TodoListRepository struct {
	db *sqlx.DB
}

func NewTodoListRepository(db *sqlx.DB) *TodoListRepository {
	return &TodoListRepository{db: db}
}

func (r *TodoListRepository) CreateList(userId int, list models.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listId int

	stmt := "INSERT INTO %s (title, description) " +
		"VALUES ($1, $2) " +
		"RETURNING id"

	createListsQuery := fmt.Sprintf(stmt, todoListsTable)
	row := tx.QueryRow(createListsQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	stmt = "INSERT INTO %s (user_id, list_id) " +
		"VALUES ($1, $2) " +
		"RETURNING id"

	createUsersListsQuery := fmt.Sprintf(stmt, usersListsTable)
	_, err = tx.Exec(createUsersListsQuery, userId, listId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}

func (r *TodoListRepository) GetAllLists(userId int) ([]models.TodoList, error) {
	var lists []models.TodoList

	stmt := "SELECT tl.id, tl.title, tl.description " +
		"FROM %s tl " +
		"INNER JOIN %s ul ON tl.id = ul.list_id " +
		"WHERE ul.user_id = $1"

	getAllListQuery := fmt.Sprintf(stmt, todoListsTable, usersListsTable)

	err := r.db.Select(&lists, getAllListQuery, userId)

	return lists, err
}

func (r *TodoListRepository) GetListById(userId, listId int) (models.TodoList, error) {
	var list models.TodoList

	stmt := "SELECT tl.id, tl.title, tl.description " +
		"FROM %s tl " +
		"INNER JOIN %s ul ON tl.id = ul.list_id " +
		"WHERE ul.user_id = $1 AND ul.list_id = $2"

	getListByIdQuery := fmt.Sprintf(stmt, todoListsTable, usersListsTable)

	err := r.db.Get(&list, getListByIdQuery, userId, listId)

	return list, err
}

func (r *TodoListRepository) UpdateList(userId, listId int, updateListInput models.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1

	if updateListInput.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argsId))
		args = append(args, *updateListInput.Title)
		argsId++
	}

	if updateListInput.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argsId))
		args = append(args, *updateListInput.Description)
		argsId++
	}

	setUpdateListQuery := strings.Join(setValues, ", ")

	stmt := "UPDATE %s tl SET %s " +
		"FROM %s ul " +
		"WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id=$%d"

	updateListQuery := fmt.Sprintf(stmt, todoListsTable, setUpdateListQuery, usersListsTable, argsId, argsId+1)

	args = append(args, listId, userId)

	_, err := r.db.Exec(updateListQuery, args...)
	return err
}

func (r *TodoListRepository) DeleteList(userId, listId int) error {
	stmt := "DELETE FROM %s tl USING %s ul " +
		"WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2"

	deleteQuery := fmt.Sprintf(stmt, todoListsTable, usersListsTable)

	_, err := r.db.Exec(deleteQuery, userId, listId)

	return err
}
