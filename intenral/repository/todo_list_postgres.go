package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lov3allmy/todo-app/intenral/domain"
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

	var id int
	createListQuery := fmt.Sprintf("insert into %s (title, description) values ($1, $2) returning id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("insert into %s (user_id, list_id) values ($1, $2)", usersListsTable)
	if _, err := tx.Exec(createUsersListQuery, userID, id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListPostgres) ReadAll(userID int) ([]domain.TodoList, error) {
	var lists []domain.TodoList
	query := fmt.Sprintf("select tl.id, tl.title, tl.description from %s tl inner join %s ul on tl.id = ul.list_id where ul.user_id = $1", todoListsTable, usersListsTable)

	err := r.db.Select(&lists, query, userID)

	return lists, err
}

func (r *TodoListPostgres) ReadByID(userID, listID int) (domain.TodoList, error) {
	var list domain.TodoList
	query := fmt.Sprintf("select tl.id, tl.title, tl.description from %s tl inner join %s ul on tl.id = ul.list_id where ul.user_id = $1 and ul.list_id = $2", todoListsTable, usersListsTable)

	err := r.db.Get(&list, query, userID, listID)

	return list, err
}
