package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lov3allmy/todo-app/intenral/domain"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listID int, item domain.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemID int
	createItemQuery := fmt.Sprintf("insert into %s (title, description) values ($1, $2) returning id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err := row.Scan(&itemID); err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	createListsItemQuery := fmt.Sprintf("insert into %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	if _, err := tx.Exec(createListsItemQuery, listID, itemID); err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return itemID, nil
}

func (r *TodoItemPostgres) ReadAll(userID, listID int) ([]domain.TodoItem, error) {
	var items []domain.TodoItem
	query := fmt.Sprintf("select ti.id, ti.title, ti.description, ti.done from %s ti inner join %s li on ti.id = li.item_id inner join %s ul on li.list_id = ul.list_id where li.list_id = $1 and ul.user_id = $2", todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&items, query, listID, userID); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) ReadByID(userID, itemID int) (domain.TodoItem, error) {
	var item domain.TodoItem
	query := fmt.Sprintf("select ti.id, ti.title, ti.description, ti.done from %s ti inner join %s li on ti.id = li.item_id inner join %s ul on li.list_id = ul.list_id where li.item_id = $1 and ul.user_id = $2", todoItemsTable, listsItemsTable, usersListsTable)

	err := r.db.Get(&item, query, itemID, userID)

	return item, err
}

func (r *TodoItemPostgres) Update(usedID, itemID int, itemInput UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if itemInput.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argID))
		args = append(args, *itemInput.Title)
		argID++
	}

	if itemInput.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argID))
		args = append(args, *itemInput.Description)
		argID++
	}

	if itemInput.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done = $%d", argID))
		args = append(args, *itemInput.Done)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("update %s ti set %s from %s li, %s ul where ti.id = li.item_id and li.list_id = ul.list_id and ul.user_id = $%d and ti.id = $%d", todoItemsTable, setQuery, listsItemsTable, usersListsTable, argID, argID+1)
	args = append(args, usedID, itemID)
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

func (r *TodoItemPostgres) Delete(userID, itemID int) error {
	query := fmt.Sprintf("delete from %s ti using %s li, %s ul where ti.id = li.item_id and li.list_id = ul.list_id and ul.list_id = $1 and ti.id = $2", todoItemsTable, listsItemsTable, usersListsTable)
	exec, err := r.db.Exec(query, userID, itemID)
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
