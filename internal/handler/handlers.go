package handler

import (
	"embed"
	"net/http"
	"strings"

	"github.com/elghazx/portfolio/internal/models"
	"github.com/elghazx/portfolio/internal/ui/pages"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	StaticFiles embed.FS
}

func NewHandler(staticFiles embed.FS) *Handler {
	return &Handler{
		StaticFiles: staticFiles,
	}
}

func (h *Handler) homeHandler(c echo.Context) error {
	if c.Request().Header.Get("HX-Request") == "true" {
		return Render(c, pages.HomeContent(Projects[:3]))
	}
	return Render(c, pages.Home(Projects[:3]))
}

func (h *Handler) projectsHandler(c echo.Context) error {
	search := strings.ToLower(c.QueryParam("search"))
	var filteredProjects []models.Project

	for _, project := range Projects {
		if search == "" || strings.Contains(strings.ToLower(project.Title), search) || strings.Contains(strings.ToLower(project.Summary), search) {
			filteredProjects = append(filteredProjects, project)
		}
	}

	// For HTMX requests
	if c.Request().Header.Get("HX-Request") == "true" {
		if c.Request().Header.Get("HX-Trigger") == "search-form" {
			return Render(c, pages.ProjectsList(filteredProjects))
		}
		return Render(c, pages.ProjectsContent(filteredProjects))
	}

	return Render(c, pages.Projects(filteredProjects))
}

func (h *Handler) projectDetailHandler(c echo.Context) error {
	slug := c.Param("slug")

	for _, project := range Projects {
		if project.Slug == slug {
			if c.Request().Header.Get("HX-Request") == "true" {
				return Render(c, pages.ProjectDetailContent(project))
			}
			return Render(c, pages.ProjectDetail(project))
		}
	}

	return h.notFoundHandler(c)
}

func (h *Handler) postsHandler(c echo.Context) error {
	search := strings.ToLower(c.QueryParam("search"))
	var filteredPosts []models.Post

	for _, post := range Posts {
		if search == "" || strings.Contains(strings.ToLower(post.Title), search) || strings.Contains(strings.ToLower(post.Summary), search) {
			filteredPosts = append(filteredPosts, post)
		}
	}

	// For HTMX requests
	if c.Request().Header.Get("HX-Request") == "true" {
		if c.Request().Header.Get("HX-Trigger") == "search-form" {
			return Render(c, pages.PostsList(filteredPosts))
		}
		return Render(c, pages.PostsContent(filteredPosts))
	}

	return Render(c, pages.Posts(filteredPosts))
}

func (h *Handler) postDetailHandler(c echo.Context) error {
	slug := c.Param("slug")

	for _, post := range Posts {
		if post.Slug == slug {
			if c.Request().Header.Get("HX-Request") == "true" {
				return Render(c, pages.PostDetailContent(post))
			}
			return Render(c, pages.PostDetail(post))
		}
	}

	return h.notFoundHandler(c)
}

func (h *Handler) experienceHandler(c echo.Context) error {
	if c.Request().Header.Get("HX-Request") == "true" {
		return Render(c, pages.ExperienceContent(Experiences))
	}
	return Render(c, pages.Experience(Experiences))
}

func (h *Handler) notFoundHandler(c echo.Context) error {
	c.Response().Status = http.StatusNotFound
	return Render(c, pages.NotFound())
}
