package handler

import (
	"net/http"

	"github.com/Puneet56/planner/view/home"
	"github.com/labstack/echo/v4"
)

func HandleHomeIndex(c echo.Context) error {
	return RenderComponent(c, http.StatusOK, home.Index("puneet"))
}
