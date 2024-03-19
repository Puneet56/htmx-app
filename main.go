package main

import (
	"os"

	"github.com/Puneet56/planner/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"

	appMiddleware "github.com/Puneet56/planner/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format:           "${time_custom} ${path} ${method} ${status}\n",
			Output:           os.Stdout,
			CustomTimeFormat: "2006-01-02 15:04:05",
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

	e.Logger.Fatal(e.Start(":1323"))
}
