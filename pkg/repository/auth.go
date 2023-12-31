package repository

import (
	"fmt"
	"github.com/aspandyar/todo-list/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) (int, error) {
	var id int

	stmt := "INSERT INTO %s (name, username, password_hash) " +
		"VALUES ($1, $2, $3) " +
		"RETURNING id"

	query := fmt.Sprintf(stmt, usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) GetUser(username, password string) (models.User, error) {
	var user models.User

	stmt := "SELECT id FROM %s " +
		"WHERE username=$1 AND password_hash=$2"

	query := fmt.Sprintf(stmt, usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
