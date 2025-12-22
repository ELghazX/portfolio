package handler

import (
	"net/http"

	"github.com/elghazx/portfolio/static/assets"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, h *Handler) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/static/*", echo.WrapHandler(http.FileServer(http.FS(assets.Files))))

	e.GET("/", h.homeHandler)
	e.GET("/projects", h.projectsHandler)
	e.GET("/experience", h.experienceHandler)
	e.GET("/posts", h.postsHandler)
	
	// 404 handler
	e.RouteNotFound("/*", h.notFoundHandler)
}
