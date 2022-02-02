package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lov3allmy/todo-app/intenral/domain"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type TodoList interface {
	Create(userID int, list domain.TodoList) (int, error)
	ReadAll(userID int) ([]domain.TodoList, error)
	ReadByID(userID, listID int) (domain.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
