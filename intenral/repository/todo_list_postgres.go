package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lov3allmy/todo-app/intenral/domain"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userID int, list domain.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listID int
	createListQuery := fmt.Sprintf("insert into %s (title, description) values ($1, $2) returning id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&listID); err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("insert into %s (user_id, list_id) values ($1, $2)", usersListsTable)
	if _, err := tx.Exec(createUsersListQuery, userID, listID); err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return listID, nil
}

func (r *TodoListPostgres) ReadAll(userID int) ([]domain.TodoList, error) {
	var lists []domain.TodoList
	query := fmt.Sprintf("select tl.id, tl.title, tl.description from %s tl inner join %s ul on tl.id = ul.list_id where ul.user_id = $1", todoListsTable, usersListsTable)

	if err := r.db.Select(&lists, query, userID); err != nil {
		return nil, err
	}

	return lists, nil
}

func (r *TodoListPostgres) ReadByID(userID, listID int) (domain.TodoList, error) {
	var list domain.TodoList
	query := fmt.Sprintf("select tl.id, tl.title, tl.description from %s tl inner join %s ul on tl.id = ul.list_id where ul.user_id = $1 and ul.list_id = $2", todoListsTable, usersListsTable)

	err := r.db.Get(&list, query, userID, listID)

	return list, err
}

func (r *TodoListPostgres) Update(usedID, listID int, listInput UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if listInput.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argID))
		args = append(args, *listInput.Title)
		argID++
	}

	if listInput.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argID))
		args = append(args, *listInput.Description)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("update %s tl set %s from %s ul where tl.id = ul.list_id and ul.user_id = $%d and ul.list_id = $%d", todoListsTable, setQuery, usersListsTable, argID, argID+1)
	args = append(args, usedID, listID)
	exec, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	rows, err := exec.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("nothing to update")
	}

	return nil
}

func (r *TodoListPostgres) Delete(userID, listID int) error {
	query := fmt.Sprintf("delete from %s tl using %s ul where tl.id = ul.list_id and ul.user_id = $1 and ul.list_id = $2", todoListsTable, usersListsTable)
	exec, err := r.db.Exec(query, userID, listID)
	if err != nil {
		return err
	}

	rows, err := exec.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("nothing to delete")
	}

	return nil
}
