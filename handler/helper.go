package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderComponent(ctx echo.Context, statusCode int, component templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}
