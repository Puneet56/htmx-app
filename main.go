package main

import (
	"os"

	"github.com/Puneet56/planner/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	appMiddleware "github.com/Puneet56/planner/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "${time_rfc3339} ${status} ${path} ${method}\n",
			Output: os.Stdout,
		},
	))

	e.Static("/assets", "assets")

	e.GET("/auth/login", handler.HandleLoginIndex)
	e.GET("/auth/login/form", handler.HandleGetLoginForm)
	e.POST("/auth/login", handler.HandleLoginSubmit)
	e.GET("/auth/logout", handler.HandleLogout)

	pr := e.Group("/")
	pr.Use(appMiddleware.Auth)

	pr.GET("", handler.HandleHomeIndex)
	pr.GET("todos", handler.HandleGetTodos)
	pr.POST("todos", handler.HandleCreateTodo)
	pr.GET("todos/new", handler.HandleAddTodoForm)

	e.Logger.Fatal(e.Start(":1323"))
}
