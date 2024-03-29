package handler

import (
	"net/http"

	"github.com/Puneet56/planner/types"
	"github.com/Puneet56/planner/view/auth"
	"github.com/Puneet56/planner/view/home"
	"github.com/labstack/echo/v4"
)

func HandleLoginIndex(c echo.Context) error {
	_, err := c.Cookie("username")

	if err == nil {
		return c.Redirect(http.StatusFound, "/")
	}

	component := auth.Index()
	return RenderComponent(c, http.StatusOK, component)
}

func HandleGetLoginForm(c echo.Context) error {
	return RenderComponent(c, http.StatusOK, auth.Login(types.LoginRequest{}, types.LoginError{}))
}

func HandleLoginSubmit(c echo.Context) error {
	lr := types.LoginRequest{}

	if err := c.Bind(&lr); err != nil {
		return RenderComponent(c, http.StatusOK, auth.Login(lr, types.LoginError{}))
	}

	le := types.LoginError{}

	if lr.Username == "" {
		le.Username = "Username is required"
	}

	if lr.Password == "" {
		le.Password = "Password is required"
	}

	if le.Username != "" || le.Password != "" {
		return RenderComponent(c, http.StatusUnauthorized, auth.Login(lr, le))
	}

	if lr.Username != "puneet" || lr.Password != "password" {
		c.Request().Header.Set(echo.HeaderAuthorization, echo.ErrUnauthorized.Error())
		le.Username = "Invalid username or password"
		return RenderComponent(c, http.StatusUnauthorized, auth.Login(lr, le))
	}

	c.SetCookie(&http.Cookie{
		Name:  "username",
		Value: lr.Username,
		Path:  "/",
	})

	user := &types.User{
		Username: lr.Username,
	}

	return RenderComponent(c, http.StatusOK, home.Index(user))
}

func HandleLogout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:   "username",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	return c.Redirect(http.StatusFound, "/auth/login")
}
