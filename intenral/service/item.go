package service

import (
	"errors"
	"github.com/lov3allmy/todo-app/intenral/domain"
	"github.com/lov3allmy/todo-app/intenral/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userID, listID int, item domain.TodoItem) (int, error) {
	if _, err := s.listRepo.ReadByID(userID, listID); err != nil {
		return 0, errors.New("list doesn't exists or doesn't belongs to user")
	}

	return s.repo.Create(listID, item)
}

func (s *TodoItemService) ReadAll(userID, listID int) ([]domain.TodoItem, error) {
	return s.repo.ReadAll(userID, listID)
}

func (s *TodoItemService) ReadByID(userID, itemID int) (domain.TodoItem, error) {
	return s.repo.ReadByID(userID, itemID)
}

func (s *TodoItemService) Update(usedID, itemID int, itemInput UpdateItemInput) error {
	if err := itemInput.Validate(); err != nil {
		return err
	}
	updateInput := repository.UpdateItemInput{
		Title:       itemInput.Title,
		Description: itemInput.Description,
		Done:        itemInput.Done,
	}

	return s.repo.Update(usedID, itemID, updateInput)
}

func (s *TodoItemService) Delete(userID, itemID int) error {
	return s.repo.Delete(userID, itemID)
}
