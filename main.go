package main

import (
	"github.com/Puneet56/planner/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/assets", "assets")

	e.GET("/", handler.HandleHomeIndex)

	e.GET("/auth/login", handler.HandleLoginIndex)
	e.GET("/auth/login/form", handler.HandleGetLoginForm)
	e.POST("/auth/login", handler.HandleLoginSubmit)

	e.Logger.Fatal(e.Start(":1323"))
}
