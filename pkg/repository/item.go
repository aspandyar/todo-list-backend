package repository

import (
	"fmt"
	"github.com/aspandyar/todo-list/internal/models"
	"github.com/jmoiron/sqlx"
)

type TodoItemRepository struct {
	db *sqlx.DB
}

func NewTodoItemRepository(db *sqlx.DB) *TodoItemRepository {
	return &TodoItemRepository{db: db}
}

func (r *TodoItemRepository) CreateItem(listId int, item models.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int

	stmt := "INSERT INTO %s (title, description) " +
		"VALUES ($1, $2) " +
		"RETURNING id"

	createItemQuery := fmt.Sprintf(stmt, todoItemTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)

	err = row.Scan(&itemId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	stmt = "INSERT INTO %s (list_id, item_id)" +
		"VALUES ($1, $2)"

	createListItemQuery := fmt.Sprintf(stmt, listsItemsTable)

	_, err = tx.Exec(createListItemQuery, listId, itemId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemRepository) GetAllItem(userId, listId int) ([]models.TodoItem, error) {
	var items []models.TodoItem

	stmt := "SELECT ti.id, ti.title, ti.description " +
		"FROM %s ti " +
		"INNER JOIN %s li ON li.item_id = ti.id " +
		"INNER JOIN %s ul ON ul.list_id = li.list_id " +
		"WHERE li.list_id = $1 AND ul.user_id = $2"

	getAllQuery := fmt.Sprintf(stmt, todoItemTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&items, getAllQuery, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemRepository) GetItemById(userId, itemId int) (models.TodoItem, error) {
	var item models.TodoItem

	stmt := "SELECT ti.id, ti.title, ti.description " +
		"FROM %s ti " +
		"INNER JOIN %s li ON li.item_id = ti.id " +
		"INNER JOIN %s ul ON ul.list_id = li.list_id " +
		"WHERE ti.id = $1 AND ul.user_id = $2"

	getAllQuery := fmt.Sprintf(stmt, todoItemTable, listsItemsTable, usersListsTable)

	if err := r.db.Get(&item, getAllQuery, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}
