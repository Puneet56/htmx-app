package handler

import (
	"math/rand"
	"net/http"

	"github.com/Puneet56/planner/types"
	view "github.com/Puneet56/planner/view/todos"
	"github.com/labstack/echo/v4"
)

func HandleGetTodos(c echo.Context) error {
	component := view.TodoIndex()
	return RenderComponent(c, http.StatusOK, component)
}

func HandleCreateTodo(c echo.Context) error {
	component := view.NewTodoRow(types.Todo{
		ID:        rand.Int(),
		Title:     "Todo 5",
		Category:  "Work",
		Completed: false,
	})

	return RenderComponent(c, http.StatusOK, component)
}

func HandleAddTodoForm(c echo.Context) error {
	return RenderComponent(c, http.StatusOK, view.AddTodoForm())
}
