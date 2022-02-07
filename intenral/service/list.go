package service

import (
	"github.com/lov3allmy/todo-app/intenral/domain"
	"github.com/lov3allmy/todo-app/intenral/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userID int, list domain.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s *TodoListService) ReadAll(userID int) ([]domain.TodoList, error) {
	return s.repo.ReadAll(userID)
}

func (s *TodoListService) ReadByID(userID, listID int) (domain.TodoList, error) {
	return s.repo.ReadByID(userID, listID)
}

func (s *TodoListService) Update(usedID, listID int, listInput UpdateListInput) error {
	if err := listInput.Validate(); err != nil {
		return err
	}
	updateInput := repository.UpdateListInput{
		Title:       listInput.Title,
		Description: listInput.Description,
	}

	return s.repo.Update(usedID, listID, updateInput)
}

func (s *TodoListService) Delete(userID, listID int) error {
	return s.repo.Delete(userID, listID)
}
