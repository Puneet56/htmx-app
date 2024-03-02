package handler

import (
	"net/http"

	"github.com/Puneet56/planner/types"
	"github.com/Puneet56/planner/view/home"
	"github.com/labstack/echo/v4"
)

func HandleHomeIndex(c echo.Context) error {
	user, ok := c.Get("user").(types.User)

	if !ok {
		c.Logger().Error("User not found")
		return c.Redirect(http.StatusUnauthorized, "/auth/login")
	}

	return RenderComponent(c, http.StatusOK, home.Index(&user))
}
