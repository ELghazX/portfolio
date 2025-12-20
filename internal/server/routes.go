package server

import (
	"net/http"

	"github.com/elghazx/goth-echo-setup/internal"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/static/*", echo.WrapHandler(http.FileServer(http.FS(internal.StaticFiles))))

	e.GET("/", root)
}
