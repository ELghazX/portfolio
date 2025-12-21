package server

import (
	"net/http"

	"github.com/elghazx/portfolio/internal"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/static/*", echo.WrapHandler(http.FileServer(http.FS(internal.StaticFiles))))

	e.GET("/", homeHandler)
	e.GET("/projects", projectsHandler)
	e.GET("/experience", experienceHandler)
	e.GET("/posts", postsHandler)
	
	// 404 handler
	e.RouteNotFound("/*", notFoundHandler)
}
