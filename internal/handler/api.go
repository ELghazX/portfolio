package handler

import (
	"embed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(staticFiles embed.FS) *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	h := NewHandler(staticFiles)
	RegisterRoutes(e, h)

	return e
}
