package handler

import (
	"net/http"
	"strconv"

	"github.com/Puneet56/planner/db"
	"github.com/Puneet56/planner/planner_orm"
	view "github.com/Puneet56/planner/view/todos"
	"github.com/labstack/echo/v4"
)

func HandleGetTodos(c echo.Context) error {
	todos, err := db.GetQuery().ListTodos(c.Request().Context())

	if err != nil {
		return err
	}

	component := view.TodoIndex(todos)
	return RenderComponent(c, http.StatusOK, component)
}

func HandleCreateTodo(c echo.Context) error {
	todo := planner_orm.Todo{}

	if err := c.Bind(&todo); err != nil {
		return err
	}

	component := view.NewTodoRow(todo)

	return RenderComponent(c, http.StatusOK, component)
}

func HandleAddTodoForm(c echo.Context) error {
	return RenderComponent(c, http.StatusOK, view.AddTodoForm(planner_orm.Todo{}))
}

func HandleEditTodoForm(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	todo, err := db.GetQuery().GetTodo(c.Request().Context(), uint64(id))

	return RenderComponent(c, http.StatusOK, view.AddTodoForm(todo))
}
