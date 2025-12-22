package handler

import (
	"embed"
	"net/http"

	"github.com/elghazx/portfolio/internal/ui/components"
	"github.com/elghazx/portfolio/internal/ui/pages"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	StaticFiles embed.FS
	htmx        *HTMXHelper
}

func NewHandler(staticFiles embed.FS) *Handler {
	return &Handler{
		StaticFiles: staticFiles,
		htmx:        NewHTMXHelper(),
	}
}

func (h *Handler) homeHandler(c echo.Context) error {
	projects := Projects[:3]
	
	if h.htmx.IsHTMXRequest(c) {
		return Render(c, pages.HomeContent(projects))
	}
	return Render(c, pages.Home(projects))
}

func (h *Handler) projectsHandler(c echo.Context) error {
	searchTerm := h.htmx.GetSearchTerm(c)
	filteredProjects := h.htmx.FilterProjects(Projects, searchTerm)

	if h.htmx.IsHTMXRequest(c) {
		h.htmx.SetNoCache(c)
		
		if searchTerm != "" {
			return Render(c, components.SearchResults("projects", pages.ProjectsList(filteredProjects)))
		}
		
		return Render(c, pages.ProjectsContent(filteredProjects))
	}

	return Render(c, pages.Projects(filteredProjects))
}

func (h *Handler) projectDetailHandler(c echo.Context) error {
	slug := c.Param("slug")

	for _, project := range Projects {
		if project.Slug == slug {
			if h.htmx.IsHTMXRequest(c) {
				return Render(c, pages.ProjectDetailContent(project))
			}
			return Render(c, pages.ProjectDetail(project))
		}
	}

	return h.notFoundHandler(c)
}

func (h *Handler) postsHandler(c echo.Context) error {
	searchTerm := h.htmx.GetSearchTerm(c)
	filteredPosts := h.htmx.FilterPosts(Posts, searchTerm)

	if h.htmx.IsHTMXRequest(c) {
		h.htmx.SetNoCache(c)
		
		if searchTerm != "" {
			return Render(c, components.SearchResults("posts", pages.PostsList(filteredPosts)))
		}
		
		return Render(c, pages.PostsContent(filteredPosts))
	}

	return Render(c, pages.Posts(filteredPosts))
}

func (h *Handler) postDetailHandler(c echo.Context) error {
	slug := c.Param("slug")

	for _, post := range Posts {
		if post.Slug == slug {
			if h.htmx.IsHTMXRequest(c) {
				return Render(c, pages.PostDetailContent(post))
			}
			return Render(c, pages.PostDetail(post))
		}
	}

	return h.notFoundHandler(c)
}

func (h *Handler) experienceHandler(c echo.Context) error {
	if h.htmx.IsHTMXRequest(c) {
		return Render(c, pages.ExperienceContent(Experiences))
	}
	return Render(c, pages.Experience(Experiences))
}

func (h *Handler) notFoundHandler(c echo.Context) error {
	c.Response().Status = http.StatusNotFound
	return Render(c, pages.NotFound())
}
