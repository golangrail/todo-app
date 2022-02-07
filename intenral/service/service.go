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

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i *UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

func (i *UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
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
	Update(usedID, listID int, listInput UpdateListInput) error
	Delete(userID, listID int) error
}

type TodoItem interface {
	Create(userID, listID int, item domain.TodoItem) (int, error)
	ReadAll(userID, listID int) ([]domain.TodoItem, error)
	ReadByID(userID, itemID int) (domain.TodoItem, error)
	Update(usedID, itemID int, itemInput UpdateItemInput) error
	Delete(userID, itemID int) error
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
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
