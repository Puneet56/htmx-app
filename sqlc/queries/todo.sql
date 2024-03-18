-- name: GetTodo :one
SELECT * FROM todos
WHERE id = ? LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id;

-- name: ListIncompleteTodos :many
SELECT * FROM todos
WHERE completed = FALSE;

-- name: ListCompleteTodos :many
SELECT * FROM todos
WHERE completed = TRUE;

-- name: CreateTodo :exec
INSERT INTO todos (title, `description`, completed)
VALUES (?, ?, ?);

-- name: UpdateTodo :exec
UPDATE todos
SET title = ?, `description` = ?, completed = ?
WHERE id = ?;


-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;