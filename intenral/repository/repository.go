package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lov3allmy/todo-app/intenral/domain"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type TodoList interface {
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
	}
}
