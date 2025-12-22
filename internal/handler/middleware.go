package handler

import (
	"strings"
	"github.com/labstack/echo/v4"
)

func ActivePageMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path
			
			var activePage string
			switch {
			case path == "/":
				activePage = "home"
			case strings.HasPrefix(path, "/projects"):
				activePage = "projects"
			case strings.HasPrefix(path, "/posts"):
				activePage = "posts"
			case strings.HasPrefix(path, "/experience"):
				activePage = "experience"
			default:
				activePage = ""
			}
			
			c.Set("activePage", activePage)
			return next(c)
		}
	}
}
