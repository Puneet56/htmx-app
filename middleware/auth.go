package middleware

import (
	"net/http"

	"github.com/Puneet56/planner/types"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Logger().Error("Path: %v", c.Path())

		cookie, err := c.Cookie("username")

		var user types.User
		if err != nil {
			return c.Redirect(http.StatusFound, "/auth/login")
			// user = types.User{
			// 	Username: "No user",
			// }

			// c.Set("user", user)
			// return next(c)
		}

		user = types.User{
			Username: cookie.Value,
		}

		c.Set("user", user)

		return next(c)
	}
}
