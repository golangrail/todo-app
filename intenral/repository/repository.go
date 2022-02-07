package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lov3allmy/todo-app/intenral/domain"
)

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type TodoList interface {
	Create(userID int, list domain.TodoList) (int, error)
	ReadAll(userID int) ([]domain.TodoList, error)
	ReadByID(userID, listID int) (domain.TodoList, error)
	Update(usedID, listID int, listInput UpdateListInput) error
	Delete(userID, listID int) error
}

type TodoItem interface {
	Create(listID int, item domain.TodoItem) (int, error)
	ReadAll(userID, listID int) ([]domain.TodoItem, error)
	ReadByID(userID, itemID int) (domain.TodoItem, error)
	Update(usedID, itemID int, itemInput UpdateItemInput) error
	Delete(userID, itemID int) error
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
		TodoItem:      NewTodoItemPostgres(db),
	}
}
