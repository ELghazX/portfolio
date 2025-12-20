package server

import (
	"github.com/elghazx/goth-echo-setup/internal/components"
	"github.com/labstack/echo/v4"
)

func root(c echo.Context) error {
	return Render(c, components.Home())
}
