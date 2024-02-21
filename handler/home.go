package handler

import (
	"net/http"

	"github.com/Puneet56/planner/view/auth"
	"github.com/Puneet56/planner/view/home"
	"github.com/labstack/echo/v4"
)

func HandleHomeIndex(c echo.Context) error {
	cookie, err := c.Cookie("username")

	if err != nil {
		return RenderComponent(c, http.StatusUnauthorized, auth.ClickToLogin())
	}

	c.Logger().Infof("Cookie: %v", cookie)

	return RenderComponent(c, http.StatusOK, home.Index(cookie.Value))
}
