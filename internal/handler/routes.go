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

	staticHandler := http.FileServer(http.FS(assets.Files))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static", staticHandler)))

	e.GET("/", h.homeHandler)
	e.GET("/projects", h.projectsHandler)
	e.GET("/projects/:slug", h.projectDetailHandler)
	e.GET("/posts", h.postsHandler)
	e.GET("/posts/:slug", h.postDetailHandler)
	e.GET("/experience", h.experienceHandler)
	
	// 404 handler
	e.RouteNotFound("/*", h.notFoundHandler)
}
