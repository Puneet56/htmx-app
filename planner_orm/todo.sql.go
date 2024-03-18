// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: todo.sql

package planner_orm

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :exec
INSERT INTO todos (title, ` + "`" + `description` + "`" + `, completed)
VALUES (?, ?, ?)
`

type CreateTodoParams struct {
	Title       string
	Description sql.NullString
	Completed   bool
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) error {
	_, err := q.db.ExecContext(ctx, createTodo, arg.Title, arg.Description, arg.Completed)
	return err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?
`

func (q *Queries) DeleteTodo(ctx context.Context, id uint64) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getTodo = `-- name: GetTodo :one
SELECT id, title, description, completed FROM todos
WHERE id = ? LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id uint64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Completed,
	)
	return i, err
}

const listCompleteTodos = `-- name: ListCompleteTodos :many
SELECT id, title, description, completed FROM todos
WHERE completed = TRUE
`

func (q *Queries) ListCompleteTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listCompleteTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Completed,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listIncompleteTodos = `-- name: ListIncompleteTodos :many
SELECT id, title, description, completed FROM todos
WHERE completed = FALSE
`

func (q *Queries) ListIncompleteTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listIncompleteTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Completed,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTodos = `-- name: ListTodos :many
SELECT id, title, description, completed FROM todos
ORDER BY id
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Completed,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :exec
UPDATE todos
SET title = ?, ` + "`" + `description` + "`" + ` = ?, completed = ?
WHERE id = ?
`

type UpdateTodoParams struct {
	Title       string
	Description sql.NullString
	Completed   bool
	ID          uint64
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) error {
	_, err := q.db.ExecContext(ctx, updateTodo,
		arg.Title,
		arg.Description,
		arg.Completed,
		arg.ID,
	)
	return err
}