package service

import (
	"errors"
	"github.com/lov3allmy/todo-app/intenral/domain"
	"github.com/lov3allmy/todo-app/intenral/repository"
)

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i *UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userID int, list domain.TodoList) (int, error)
	ReadAll(userID int) ([]domain.TodoList, error)
	ReadByID(userID, listID int) (domain.TodoList, error)
	Update(usedID, listID int, input UpdateListInput) error
	Delete(userID, listID int) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
	}
}
